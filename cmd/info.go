/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"versioner/application/usecases"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")

		infoUseCase := usecases.NewInfoUseCase()

		infoUseCase.Execute()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
