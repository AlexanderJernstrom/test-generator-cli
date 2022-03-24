package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func GetPackageName(fileName string, fileContent string) {
	fileSet := token.NewFileSet()

	file, err := parser.ParseFile(fileSet, fileName, fileContent, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, declaration := range file.Decls {
		_, ok := declaration.(*ast.FuncDecl)
		if ok {
			continue
		}

	}
}
