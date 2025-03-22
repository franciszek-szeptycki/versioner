/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"versioner/adapters"
	"versioner/application/selectors"
	"versioner/application/services"
	"versioner/application/usecases"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")

		var fileAdapter selectors.IFileAdapter
		fileAdapter = adapters.NewFileAdapter()

		getProjectInfoService := *services.NewGetProjectInfoService(fileAdapter)

		infoUseCase := usecases.NewInfoUseCase(getProjectInfoService)

		infoUseCase.Execute()
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
