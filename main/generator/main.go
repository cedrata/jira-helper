package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/cedrata/jira-helper/generator"
)

var (
	structName string
	unmarshal  bool
)

func init() {
	flag.StringVar(&structName, "struct", "", "Name of the struct to process")
	flag.BoolVar(&unmarshal, "unmarshal", false, "Set to true to generate the unmarshal stub")
	flag.Parse()

	if structName == "" || !unmarshal {
		fmt.Fprintf(os.Stderr, "Usage: %s -struct <struct_name> -unmarshal\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
	fullPath, _ := filepath.Abs(os.Getenv("GOFILE"))
	fullPathDestinationDir := filepath.Dir(fullPath)
	_, err := os.Stat(fullPathDestinationDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	fset := token.NewFileSet()
	parsedFile, _ := parser.ParseFile(fset, fullPath, nil, parser.ParseComments)

	knownProperties := []string{}

	var errs error

	// This function should be used to check weather or not the required action
	// for the selected struct are allowed, if some validation error occurs
	// then an error should be returned.
	inspectFunction := func(n ast.Node) bool {
		var typeSpec *ast.TypeSpec
		var ok bool
		if typeSpec, ok = n.(*ast.TypeSpec); !ok {
			return true
		}

		if typeSpec.Name.Name != structName {
			return true
		}

		var structType *ast.StructType
		if structType, ok = typeSpec.Type.(*ast.StructType); !ok {
			return true
		}

		// Validating for unmarshal action
		for _, field := range structType.Fields.List {
			if field.Tag == nil {
				errs = errors.Join(
					errs,
					fmt.Errorf(
						"in %s struct, missing tag for the %v field",
						structName,
						field.Names,
					),
				)
				continue
			}

			tagValue := reflect.StructTag(
				field.Tag.Value[1 : len(field.Tag.Value)-1],
			)

			jsonKey := strings.Split(tagValue.Get("json"), ",")[0]

			if jsonKey == "" {
				errs = errors.Join(
					errs,
					fmt.Errorf(
						"in %s struct, JSON tag is not set for the %v field",
						structName, field.Names,
					),
				)
				continue
			}

			if jsonKey != "-" {
				knownProperties = append(knownProperties, jsonKey)
				continue
			}

			switch t := field.Type.(type) {
			case *ast.MapType:
				_, isString := t.Key.(*ast.Ident)
				_, isInterface := t.Value.(*ast.InterfaceType)
				if !isString || !isInterface {
					errs = errors.Join(
						errs,
						fmt.Errorf(
							"in %s struct, expected to have %v field having JSON tag \"-\" to be of type map[string]interface{}",
							structName, field.Names,
						),
					)
				}
			default:
				errs = errors.Join(
					errs,
					fmt.Errorf(
						"in %s struct, expected to have %v field having JSON tag \"-\" to be of type map[string]interface{}",
						structName, field.Names,
					),
				)
			}
		}

		return false
	}

	ast.Inspect(parsedFile, inspectFunction)

	if errs != nil {
		fmt.Fprintf(
			os.Stderr,
			"some errors occured validating the requested actions for the %s struct: %s\n",
			structName, errs,
		)
		os.Exit(1)
	}

	code := generator.GenerateCodeSteps(
		generator.GenerateHeader,
		generator.GeneratePackageGenerator(os.Getenv("GOPACKAGE")),
		generator.GenerateImportGenerator([]string{"encoding/json"}),
		generator.GenerateUnmarshalCodeGenerator(structName, knownProperties),
	)

	formattedCode, err := format.Source(code)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting generated code: %s\n", err)
		os.Exit(1)
	}

	fullPathDestination := filepath.Join(
		fullPathDestinationDir,
		fmt.Sprintf("%sgenerated.go",
			strings.ToLower(structName),
		),
	)
	destinatioFile, err := os.OpenFile(
        fullPathDestination, 
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0644,
	)
	defer func() {
		err := destinatioFile.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		}
	}()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		os.Exit(1)
	}

	fmt.Fprintf(destinatioFile, "%s", formattedCode)
}
