package usecases

import (
	"fmt"
	"path/filepath"
	"versioner/adapters"
	"versioner/application/constants"
	"versioner/application/services"
)

type InitUseCase struct {
	getVersionerPathService services.GetVersionerPathService
	createVersionerService  services.CreateVersionerService
	writeConfigService      services.WriteConfigService
}

func NewInitUseCase() *InitUseCase {

	fileAdapter := adapters.NewFileAdapter()

	createVersionerService := *services.NewCreateVersionerService(fileAdapter)
	getVersionerPathService := *services.NewGetVersionerPathService(fileAdapter)
	writeConfigService := *services.NewWriteConfigService(fileAdapter)

	return &InitUseCase{
		getVersionerPathService: getVersionerPathService,
		createVersionerService:  createVersionerService,
		writeConfigService:      writeConfigService,
	}
}

func (i *InitUseCase) Execute() {
	path, err := i.getVersionerPathService.Execute()
	if err == nil {
		fmt.Printf("Versioner already initialized with config: %v\n", path)
		return
	}
	versionerPath, err := i.createVersionerService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	configPath := filepath.Join(versionerPath, constants.VersionerConfig)
	i.writeConfigService.Execute(configPath, struct{}{})
}
