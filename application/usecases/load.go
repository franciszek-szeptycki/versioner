package usecases

import (
	"fmt"
	"versioner/adapters/fileadapter"
	"versioner/application/services"
)

type LoadUseCase struct {
	getVersionerPathService services.GetVersionerPathService
	listVersionsService     services.ListVersionsService
	loadVersionService      services.LoadVersionService
}

func NewLoadUseCase() *LoadUseCase {

	fileAdapter := fileadapter.NewFileAdapter()

	return &LoadUseCase{
		getVersionerPathService: *services.NewGetVersionerPathService(fileAdapter),
		listVersionsService:     *services.NewListVersionsService(fileAdapter),
		loadVersionService:      *services.NewLoadVersionService(fileAdapter),
	}
}

func (l *LoadUseCase) Execute(version string) {
	versionerPath, err := l.getVersionerPathService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	availableVersions, err := l.listVersionsService.Execute(versionerPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, availableVersion := range availableVersions {
		if availableVersion == version {
			err = l.loadVersionService.Execute(versionerPath, version)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("version loaded")
			return
		}
	}

	fmt.Println("version not found")
}
