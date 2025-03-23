package usecases

import (
	"fmt"
	"path/filepath"
	"versioner/adapters"
	"versioner/application/constants"
	"versioner/application/services"
)

type SaveUseCase struct {
	askForVersionNameService          services.AskForVersionNameService
	createVersionService              services.CreateVersionService
	getVersionerPathService           services.GetVersionerPathService
	saveCurrentVersionToConfigService services.SaveCurrentVersionToConfigService
}

func NewSaveUseCase() *SaveUseCase {
	fileAdapter := adapters.NewFileAdapter()
	cliInputAdapter := adapters.NewCLIUserInputAdapter()

	return &SaveUseCase{
		askForVersionNameService:          *services.NewAskForVersionNameService(cliInputAdapter),
		createVersionService:              *services.NewCreateVersionService(fileAdapter),
		getVersionerPathService:           *services.NewGetVersionerPathService(fileAdapter),
		saveCurrentVersionToConfigService: *services.NewSaveCurrentVersionToConfigService(fileAdapter),
	}
}

func (s *SaveUseCase) Execute() {
	versionerPath, err := s.getVersionerPathService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	version := s.askForVersionNameService.Execute()
	s.createVersionService.Execute(filepath.Join(versionerPath, version))

	versionConfigPath := filepath.Join(versionerPath, constants.VersionerConfig)
	err = s.saveCurrentVersionToConfigService.Execute(versionConfigPath, version)
	if err != nil {
		fmt.Println(err)
	}
}
