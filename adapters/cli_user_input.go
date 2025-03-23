package adapters

import (
	"bufio"
	"fmt"
	"os"
)

type CLIUserInputAdapter struct{}

func NewCLIUserInputAdapter() *CLIUserInputAdapter {
	return &CLIUserInputAdapter{}
}

func (c *CLIUserInputAdapter) AskUserInput(question string) string {
	fmt.Println(question + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
