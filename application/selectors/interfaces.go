package selectors

type IFileAdapter interface {
	GetCurrentDir() string
	GetParentDir(path string) string
	CheckIfDirExists(path string) bool
	ReadVersionerConfigFile(path string) (VersionerConfigDTO, error)
	CreateDir(path string) error
	CreateFile(path string) error

	WriteJsonFile(path string, data any) error
}

type ICLIAskUserInputAdapter interface {
	AskUserInput(question string) string
}
