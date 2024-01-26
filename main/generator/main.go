package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <struct_name>\n", os.Args[0])
		os.Exit(1)
	}

	structName := os.Args[1]

	// Parse the source code of the main project
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, ".", nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing source code: %s\n", err)
		os.Exit(1)
	}

	// Iterate through packages and files
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				// Check if the node is a struct declaration
				if typeSpec, ok := n.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						// Check if the struct has the specified name
						if typeSpec.Name.Name == structName {
							// Generate code based on the struct
							generatedCode := generateCode(structName, structType)
							fmt.Println(generatedCode)
							os.Exit(0)
						}
					}
				}
				return true
			})
		}
	}

	fmt.Fprintf(os.Stderr, "Struct with name %s not found\n", structName)
	os.Exit(1)
}

func generateCode(structName string, structType *ast.StructType) string {
	// Example: Generate code that prints the struct's name
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package main\n\n")
	fmt.Fprintf(&buf, "import \"fmt\"\n\n")
	fmt.Fprintf(&buf, "func main() {\n")
	fmt.Fprintf(&buf, "\t%s := %s{}\n", strings.ToLower(structName), structName)
	fmt.Fprintf(&buf, "\tfmt.Printf(\"%s: %%+v\\n\", %s)\n", structName, strings.ToLower(structName))
	fmt.Fprintf(&buf, "}\n")

	// Format the generated code
	formattedCode, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting generated code: %s\n", err)
		os.Exit(1)
	}

	return string(formattedCode)
}

