package services

import "versioner/application/selectors"

type WriteConfigService struct {
	fileAdapter selectors.IFileAdapter
}

func NewWriteConfigService(fileAdapter selectors.IFileAdapter) *WriteConfigService {
	return &WriteConfigService{
		fileAdapter: fileAdapter,
	}
}

func (w *WriteConfigService) Execute(path string, data any) error {
	return w.fileAdapter.WriteJsonFile(path, data)
}
