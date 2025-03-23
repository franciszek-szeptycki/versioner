package usecases

import (
	"versioner/application/selectors"
	"versioner/application/services"
)

type ListUseCase struct {
	listVersionsService services.ListVersionsService
}

func NewListUseCase(fileAdapter selectors.FileAdapter) *ListUseCase {
	return &ListUseCase{
		listVersionsService: NewListVersionsService(fileAdapter),
	}
}
