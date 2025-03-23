package services

import "versioner/application/selectors"

type ListVersionsService struct {
	fileAdapter selectors.IFileAdapter
}

func NewListVersionsService(fileAdapter selectors.IFileAdapter) *ListVersionsService {
	return &ListVersionsService{
		fileAdapter: fileAdapter,
	}
}
