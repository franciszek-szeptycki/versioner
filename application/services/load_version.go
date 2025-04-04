package services

import (
	"fmt"
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

func (l *LoadVersionService) Execute(versionPath string) error {
	ignoreFiles := constants.IgnoreFiles
	var err error

	currentPath := l.fileAdapter.GetCurrentDir()
	err = l.fileAdapter.RemoveDir(currentPath, ignoreFiles)
	if err != nil {
		return err
	}

	fmt.Println(versionPath)
	err = l.fileAdapter.CopyDir(versionPath, currentPath, ignoreFiles)
	if err != nil {
		return err
	}

	return nil
}
