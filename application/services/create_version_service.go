package services

import (
	"fmt"
	"path/filepath"
	"versioner/application/constants"
	"versioner/application/selectors"
)

type CreateVersionService struct {
	fileAdapter selectors.IFileAdapter
}

func NewCreateVersionService(fileAdapter selectors.IFileAdapter) *CreateVersionService {
	return &CreateVersionService{
		fileAdapter: fileAdapter,
	}
}

func (c *CreateVersionService) Execute(versionPath string) {
	err := c.fileAdapter.CreateDir(versionPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentPath := c.fileAdapter.GetCurrentDir()
	versionerIgnorePath := filepath.Join(currentPath, constants.VersionerIgnore)
	ignoreFiles, err := c.fileAdapter.ReadFile(versionerIgnorePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	ignoreFiles = append(ignoreFiles, constants.IgnoreFiles...)

	c.fileAdapter.CopyDir(currentPath, versionPath, ignoreFiles)
}
