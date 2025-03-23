package usecases

import (
	"fmt"
	"path/filepath"
	"versioner/adapters"
	"versioner/application/constants"
	"versioner/application/services"
)

type InfoUseCase struct {
	getVersionerPath services.GetVersionerPathService
	getInfoService   services.GetVersionerInfoService
}

func NewInfoUseCase() *InfoUseCase {

	fileAdapter := adapters.NewFileAdapter()

	return &InfoUseCase{
		getVersionerPath: *services.NewGetVersionerPathService(fileAdapter),
		getInfoService:   *services.NewGetVersionerInfoService(fileAdapter),
	}
}

func (i InfoUseCase) Execute() {
	versionerPath, err := i.getVersionerPath.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}

	versionerConfigPath := filepath.Join(versionerPath, constants.VersionerConfig)
	config, err := i.getInfoService.Execute(versionerConfigPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config)
}
