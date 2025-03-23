package services

import (
	"versioner/application/constants"
	"versioner/application/selectors"
)

type LoadVersionService struct {
	fileAdapter selectors.IFileAdapter
}

func NewLoadVersionService(fileAdapter selectors.IFileAdapter) *LoadVersionService {
	return &LoadVersionService{
		fileAdapter: fileAdapter,
	}
}

func (l *LoadVersionService) Execute(versionerPath, version string) error {
	ignoreFiles := constants.IgnoreFiles

	currentPath := l.fileAdapter.GetCurrentDir()
	return l.fileAdapter.RemoveDir(currentPath, ignoreFiles)
}
