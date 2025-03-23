package adapters

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func (fa *FileAdapter) ReadJsonFile(path string, dto interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(byteValue, dto); err != nil {
		return err
	}

	return nil
}

func (fa *FileAdapter) ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (fa *FileAdapter) CopyDir(srcDir, dstDir string, ignoredPaths []string) error {
	ignoredSet := make(map[string]struct{})
	for _, ignore := range ignoredPaths {
		ignoredSet[ignore] = struct{}{}
	}

	return filepath.WalkDir(srcDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for ignore := range ignoredSet {
			if strings.Contains(path, ignore) {
				return nil
			}
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dstDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		return fa.CopyFile(path, destPath)
	})
}

func (fa *FileAdapter) CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func (fa *FileAdapter) ListDirs(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Błąd:", err)
		return nil, err
	}
	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}
