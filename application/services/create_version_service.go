package services

import (
	"fmt"
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
}
