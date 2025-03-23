package usecases

import (
	"fmt"
	"path/filepath"
	"versioner/adapters"
	"versioner/application/services"
)

type SaveUseCase struct {
	askForVersionNameService services.AskForVersionNameService
	createVersionService     services.CreateVersionService
	getVersionerPathService  services.GetVersionerPathService
}

func NewSaveUseCase() *SaveUseCase {
	fileAdapter := adapters.NewFileAdapter()
	cliInputAdapter := adapters.NewCLIUserInputAdapter()

	return &SaveUseCase{
		askForVersionNameService: *services.NewAskForVersionNameService(cliInputAdapter),
		createVersionService:     *services.NewCreateVersionService(fileAdapter),
		getVersionerPathService:  *services.NewGetVersionerPathService(fileAdapter),
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

}
