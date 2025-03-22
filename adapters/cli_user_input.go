package adapters

import (
	"bufio"
	"fmt"
	"os"
)

type CLIAskUserInputAdapter struct{}

func NewCLIAskUserInputAdapter() *CLIAskUserInputAdapter {
	return &CLIAskUserInputAdapter{}
}

func (c *CLIAskUserInputAdapter) AskUserInput(question string) string {
	fmt.Println(question + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
