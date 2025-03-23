package cmd

import (
	"fmt"
	"versioner/adapters"
	"versioner/application/usecases"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")

		fileAdapter := adapters.NewFileAdapter()
		listUseCase := usecases.NewListUseCase(fileAdapter)
		listUseCase.Execute()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
