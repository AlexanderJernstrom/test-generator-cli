package utils

import (
	"fmt"
	"testing"
)

func TestReadFunctions(t *testing.T) {
	code := `package main  
	func main() { 
		 fmt.Println("Hello, world") 
	} 
	
	func calc() { 
		fmt.Println("Hello, world") 
	}`

	expectedNames := []string{"Testmain", "Testcalc"}

	functionNames, err := ReadFunctions(code, false, "main.go")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(functionNames)

	for i, name := range functionNames {
		if name != expectedNames[i] {
			t.Errorf("Expected %s, got %s", expectedNames[i], name)
		}
	}
}
