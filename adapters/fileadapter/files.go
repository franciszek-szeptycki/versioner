package fileadapter

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

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
