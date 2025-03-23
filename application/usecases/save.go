package usecases

import (
	"fmt"
	"path/filepath"
	"versioner/adapters/fileadapter"
	"versioner/application/constants"
	"versioner/application/services"
)

type SaveUseCase struct {
	createVersionService              services.CreateVersionService
	getVersionerPathService           services.GetVersionerPathService
	saveCurrentVersionToConfigService services.SaveCurrentVersionToConfigService
}

func NewSaveUseCase() *SaveUseCase {
	fileAdapter := fileadapter.NewFileAdapter()

	return &SaveUseCase{
		createVersionService:              *services.NewCreateVersionService(fileAdapter),
		getVersionerPathService:           *services.NewGetVersionerPathService(fileAdapter),
		saveCurrentVersionToConfigService: *services.NewSaveCurrentVersionToConfigService(fileAdapter),
	}
}

func (s *SaveUseCase) Execute(version string) {
	versionerPath, err := s.getVersionerPathService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	s.createVersionService.Execute(filepath.Join(versionerPath, version))

	versionConfigPath := filepath.Join(versionerPath, constants.VersionerConfig)
	err = s.saveCurrentVersionToConfigService.Execute(versionConfigPath, version)
	if err != nil {
		fmt.Println(err)
	}
}
