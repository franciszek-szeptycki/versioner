package services

import "versioner/application/selectors"

type AdapterService struct {
	fileAdapter selectors.IFileAdapter
}

func NewService(fileAdapter selectors.IFileAdapter) *AdapterService {
	return &AdapterService{
		fileAdapter: fileAdapter,
	}
}
