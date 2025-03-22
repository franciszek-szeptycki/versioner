package services

import (
	"errors"
	"path/filepath"
	"versioner/application/constants"
	"versioner/application/selectors"
)

type GetProjectInfoService struct {
	fileAdapter selectors.IFileAdapter
}

func NewGetProjectInfoService(fileAdapter selectors.IFileAdapter) *GetProjectInfoService {
	return &GetProjectInfoService{
		fileAdapter: fileAdapter,
	}
}

func (g *GetProjectInfoService) Execute() (selectors.VersionerConfigDTO, error) {

	currentPath := g.fileAdapter.GetCurrentDir()

	for {
		versionerPath := filepath.Join(currentPath, constants.VersionerDirName)
		if g.fileAdapter.CheckIfDirExists(versionerPath) {
			return g.readConfigFile(currentPath)
		}

		parentDir := g.fileAdapter.GetParentDir(currentPath)
		if parentDir == currentPath {
			return selectors.VersionerConfigDTO{}, errors.New("no project config found")
		}

		currentPath = parentDir
	}

}

func (g *GetProjectInfoService) readConfigFile(versionerPath string) (selectors.VersionerConfigDTO, error) {
	configPath := filepath.Join(versionerPath, constants.VersionerConfigFileName)

	config, err := g.fileAdapter.ReadVersionerConfigFile(configPath)
	if err != nil {
		return selectors.VersionerConfigDTO{}, err
	}

	return config, nil
}
