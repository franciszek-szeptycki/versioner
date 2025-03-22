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
		saveUseCase := usecases.NewSaveUseCase()
		saveUseCase.Execute()
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
