package selectors

type IFileAdapter interface {
	GetCurrentDir() string
	GetParentDir(path string) string
	CheckIfDirExists(path string) bool
	ReadVersionerConfigFile(path string) (VersionerConfigDTO, error)
	CreateDir(path string) error
	CreateFile(path string) error

	WriteJsonFile(path string, data any) error

	ReadFile(path string) ([]string, error)
	CopyDir(src, dst string, ignoredPaths []string) error
	CopyFile(src, dst string) error
}

type ICLIUserInputAdapter interface {
	AskUserInput(question string) string
}
