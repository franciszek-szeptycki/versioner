/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"versioner/application/usecases"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("save called")
		if len(args) == 0 {
			fmt.Println("no version provided")
			return
		}
		saveUseCase := usecases.NewSaveUseCase()
		saveUseCase.Execute(args[0])
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
