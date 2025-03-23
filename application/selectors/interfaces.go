package selectors

type IFileAdapter interface {
	GetCurrentDir() string
	GetParentDir(path string) string
	CheckIfDirExists(path string) bool
	CreateDir(path string) error
	CopyDir(src, dst string, ignoredPaths []string) error
	ListDirs(path string) ([]string, error)
	RemoveDir(path string, ignoredPaths []string) error

	CreateFile(path string) error
	WriteJsonFile(path string, data any) error
	ReadFile(path string) ([]string, error)
	CopyFile(src, dst string) error
	ReadJsonFile(path string, dto interface{}) error
}

type ICLIUserInputAdapter interface {
	AskUserInput(question string) string
}
