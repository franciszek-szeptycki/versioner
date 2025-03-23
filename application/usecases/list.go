package usecases

import (
	"fmt"
	"versioner/application/selectors"
	"versioner/application/services"
)

type ListUseCase struct {
	getVersionerPathService services.GetVersionerPathService
	listVersionsService     services.ListVersionsService
}

func NewListUseCase(fileAdapter selectors.IFileAdapter) *ListUseCase {
	return &ListUseCase{
		getVersionerPathService: *services.NewGetVersionerPathService(fileAdapter),
		listVersionsService:     *services.NewListVersionsService(fileAdapter),
	}
}

func (l *ListUseCase) Execute() {
	versionerPath, err := l.getVersionerPathService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	dirs, err := l.listVersionsService.Execute(versionerPath)
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}
