package services

import (
	"path/filepath"
	"versioner/application/constants"
	"versioner/application/selectors"
)

type CreateVersionerService struct {
	fileAdapter selectors.IFileAdapter
}

func NewCreateVersionerService(fileAdapter selectors.IFileAdapter) *CreateVersionerService {
	return &CreateVersionerService{
		fileAdapter: fileAdapter,
	}
}

func (i *CreateVersionerService) Execute() (string, error) {
	currentPath := i.fileAdapter.GetCurrentDir()

	versionerPath := filepath.Join(currentPath, constants.VersionerDir)
	err := i.fileAdapter.CreateDir(versionerPath)
	if err != nil {
		return "", err
	}

	versionerIgnorePath := filepath.Join(currentPath, constants.VersionerIgnore)
	err = i.fileAdapter.CreateFile(versionerIgnorePath)
	if err != nil {
		return "", nil
	}

	return versionerPath, err
}
