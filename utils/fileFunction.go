package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func ReadFunctions(src string) ([]string, error) {
	fileSet := token.NewFileSet()


	file, err := parser.ParseFile(fileSet, "main.go", src, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	functionNames := []string{}

	for _, declaration := range file.Decls {

		fn, ok := declaration.(*ast.FuncDecl)
		if !ok {
			continue
		}
		

		functionNames = append(functionNames, fn.Name.Name)

	}

	for i, name := range(functionNames) {
		functionNames[i] = "Test" + name
	}

	return functionNames, nil
}