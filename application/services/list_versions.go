package services

import (
	"versioner/application/selectors"
)

type ListVersionsService struct {
	fileAdapter selectors.IFileAdapter
}

func NewListVersionsService(fileAdapter selectors.IFileAdapter) *ListVersionsService {
	return &ListVersionsService{
		fileAdapter: fileAdapter,
	}
}

func (l *ListVersionsService) Execute(versionerPath string) ([]string, error) {
	dirs, err := l.fileAdapter.ListDirs(versionerPath)

	var versionNames []string
	for _, dir := range dirs {
		if dir[0] != '.' {
			versionNames = append(versionNames, dir)
		}
	}

	return versionNames, err
}
