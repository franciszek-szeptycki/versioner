package usecases

import (
	"fmt"
	"versioner/application/services"
)

type InfoUseCase struct {
	getInfoService services.GetProjectInfoService
}

func NewInfoUseCase(getInfoService services.GetProjectInfoService) *InfoUseCase {
	return &InfoUseCase{
		getInfoService: getInfoService,
	}
}

func (i InfoUseCase) Execute() {
	config, err := i.getInfoService.Execute()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)
}
