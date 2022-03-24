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
	Use:   "read",
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
		files, err := utils.WalkDir()

		for _, file := range files {
			functionNames, error := utils.ReadFunctions(file.Content, true, file.Path)
			utils.GetPackageName(file.Path, file.Content)
			if error != nil {
				fmt.Println(error)
			}
			for i, name := range functionNames {
				fmt.Println(i, name)
			}
		}

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Succes")
	},
}
