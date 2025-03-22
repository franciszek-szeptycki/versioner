package adapters

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"versioner/application/selectors"
)

type FileAdapter struct{}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{}
}

func (fa *FileAdapter) GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func (fa *FileAdapter) GetParentDir(path string) string {
	return filepath.Dir(path)
}

func (fa *FileAdapter) CheckIfDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (fa *FileAdapter) ReadVersionerConfigFile(path string) (selectors.VersionerConfigDTO, error) {
	file, err := os.Open(path)
	if err != nil {
		return selectors.VersionerConfigDTO{}, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return selectors.VersionerConfigDTO{}, err
	}

	var config selectors.VersionerConfigDTO
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return selectors.VersionerConfigDTO{}, err
	}

	return config, nil
}

func (fa *FileAdapter) CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	return err
}

func (fa *FileAdapter) CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}

func (fa *FileAdapter) WriteJsonFile(path string, data any) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
