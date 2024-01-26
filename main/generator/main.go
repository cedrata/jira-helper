// generator.go

package main

import (
	"bytes"
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
)

var (
	structName string
	unmarshal  bool
	marshal    bool
)

func init() {
	flag.StringVar(&structName, "struct", "", "Name of the struct to process")
	flag.BoolVar(&unmarshal, "marshal", false, "Set to true to generate the unmarshal stub")
	flag.BoolVar(&marshal, "unmarshal", false, "Set to true to generate the marshal stub")
	flag.Parse()

	if structName == "" || (!unmarshal && !marshal) {
		fmt.Fprintf(os.Stderr, "Usage: %s -struct <struct_name> [[-marshal] | [-unmarshal] | [-marshal -unmarshal]]\n", os.Args[0])
		os.Exit(1)
	}
}

func main() {
	// Parse the source code of the main project
	fset := token.NewFileSet()
	wd, _ := os.Getwd()
	pkgs, err := parser.ParseDir(fset, wd, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing source code: %s\n", err)
		os.Exit(1)
	}

	// Iterate through packages and files
	for _, pkg := range pkgs {
		for filename, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				// Check if the node is a struct declaration
				if typeSpec, ok := n.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						// Check if the struct has the specified name
						if typeSpec.Name.Name == structName {

							generatedCode := generateCode(structName, structType, filename)
							fmt.Println(generatedCode)
							os.Exit(0)
						}
					}
				}
				return true
			})
		}
	}

    // Store the generated function to a []string
    // Search now in the file having a filename like <source_filename>_gen.go
    // the new functions, if already in there replace them
    // if not append them to the file

	fmt.Fprintf(os.Stderr, "Struct with name %s not found\n", structName)
	os.Exit(1)
}

func generateCode(structName string, structType *ast.StructType, filename string) string {
	// Example: Generate code that prints the struct's name, full file path, attributes, and JSON tags
	absPath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting absolute path: %s\n", err)
		os.Exit(1)
	}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package main\n\n")
	fmt.Fprintf(&buf, "import (\n\t\"fmt\"\n\t\"reflect\"\n)\n\n")
	fmt.Fprintf(&buf, "func main() {\n")
	fmt.Fprintf(&buf, "\t%s := %s{}\n", strings.ToLower(structName), structName)
	fmt.Fprintf(&buf, "\tfmt.Printf(\"%s: %%+v\\n\", %s)\n", structName, strings.ToLower(structName))
	fmt.Fprintf(&buf, "\tfmt.Printf(\"File path: %s\\n\")\n", absPath)

	// Iterate through struct fields and print their attributes and JSON tags
	for _, field := range structType.Fields.List {
		fieldName := field.Names[0].Name
		// fieldType := field.Type

		// Extract JSON tags, if present
		jsonTag := ""
		if field.Tag != nil {
			tagValue := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
			jsonTag = tagValue.Get("json")
		}

		// Print field attributes and JSON tags
		fmt.Fprintf(&buf, "\tfmt.Printf(\"Field: %s, Type: %%T, JSON Tag: %s\\n\", %s.%s)\n", fieldName, jsonTag, strings.ToLower(structName), fieldName)
	}

	fmt.Fprintf(&buf, "}\n")

	// Format the generated code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting generated code: %s\n", err)
		os.Exit(1)
	}

	return string(formattedCode)
}
