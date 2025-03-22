package services

import (
	"versioner/application/selectors"
)

type GetVersionerInfoService struct {
	fileAdapter selectors.IFileAdapter
}

func NewGetVersionerInfoService(fileAdapter selectors.IFileAdapter) *GetVersionerInfoService {
	return &GetVersionerInfoService{
		fileAdapter: fileAdapter,
	}
}

func (g *GetVersionerInfoService) Execute(versionerConfigPath string) (selectors.VersionerConfigDTO, error) {
	config, err := g.fileAdapter.ReadVersionerConfigFile(versionerConfigPath)
	if err != nil {
		return selectors.VersionerConfigDTO{}, err
	}

	return config, nil
}
