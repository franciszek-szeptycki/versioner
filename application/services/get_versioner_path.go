package services

import (
	"errors"
	"path/filepath"
	"versioner/application/constants"
	"versioner/application/selectors"
)

type GetVersionerPathService struct {
	fileAdapter selectors.IFileAdapter
}

func NewGetVersionerPathService(fileAdapter selectors.IFileAdapter) *GetVersionerPathService {
	return &GetVersionerPathService{
		fileAdapter: fileAdapter,
	}
}

func (g *GetVersionerPathService) Execute() (string, error) {

	currentPath := g.fileAdapter.GetCurrentDir()

	for {
		versionerPath := filepath.Join(currentPath, constants.VersionerDir)
		if g.fileAdapter.CheckIfDirExists(versionerPath) {
			return versionerPath, nil
		}

		parentDir := g.fileAdapter.GetParentDir(currentPath)
		if parentDir == currentPath {
			return "", errors.New("no project config found")
		}

		currentPath = parentDir
	}
}
