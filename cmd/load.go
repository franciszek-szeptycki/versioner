/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"versioner/application/usecases"

	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("load called")
		if len(args) == 0 {
			fmt.Println("no version provided")
			return
		}
		loadUseCase := usecases.NewLoadUseCase()
		loadUseCase.Execute(args[0])
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
