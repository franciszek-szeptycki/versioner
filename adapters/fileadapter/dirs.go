package fileadapter

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (fa *FileAdapter) GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func (fa *FileAdapter) CreateDir(path string) error {
	err := os.Mkdir(path, 0755)
	return err
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

func (fa *FileAdapter) RemoveDir(srcDir string, ignoredPaths []string) error {
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

		return os.Remove(path)
	})
}
