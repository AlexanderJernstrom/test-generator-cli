package cmd

import (
	"fmt"
	"test-generator/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCommand)
}

var readCommand = &cobra.Command{
	Use: "read",
	Short: "Gets all the functions in a .go file",
	Run: func(cmd *cobra.Command, args []string) {
		//pathToFile := args[0] 
		const src = `package main  
		func main() { 
			 fmt.Println("Hello, world") 
		} 
		
		func calc() { 
			fmt.Println("Hello, world") 
		}`
		functionNames, error := utils.ReadFunctions(src)

		if error != nil {
			fmt.Println(error)
		}

		for i, name := range functionNames {
			fmt.Println(i, name)
		}
		files, err := utils.WalkDir()

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(files)
	},
}