package services

import (
	"versioner/application/selectors"
)

type SaveCurrentVersionToConfigService struct {
	fileAdapter selectors.IFileAdapter
}

func NewSaveCurrentVersionToConfigService(fileAdapter selectors.IFileAdapter) *SaveCurrentVersionToConfigService {
	return &SaveCurrentVersionToConfigService{
		fileAdapter: fileAdapter,
	}
}

func (s *SaveCurrentVersionToConfigService) Execute(versionerPath string, version string) error {
	var config = selectors.VersionerConfigDTO{}
	err := s.fileAdapter.ReadJsonFile(versionerPath, &config)
	if err != nil {
		return err
	}

	config.CurrentVersion = &version
	err = s.fileAdapter.WriteJsonFile(versionerPath, config)
	return err
}
