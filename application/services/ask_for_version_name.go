package services

import (
	"regexp"
	"strings"
	"versioner/application/selectors"
)

type AskForVersionNameService struct {
	cliAskUserInputAdapter selectors.ICLIUserInputAdapter
}

func NewAskForVersionNameService(cliUserInputAdapter selectors.ICLIUserInputAdapter) *AskForVersionNameService {
	return &AskForVersionNameService{
		cliAskUserInputAdapter: cliUserInputAdapter,
	}
}

func (a *AskForVersionNameService) Execute() string {

	var versionName string
	for {
		versionName = a.cliAskUserInputAdapter.AskUserInput("Enter new version name")
		sanitizedString := a.sanitizeString(versionName)

		return sanitizedString
	}
}

func (a *AskForVersionNameService) sanitizeString(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
	s = re.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, " ", "-")
	return s
}
