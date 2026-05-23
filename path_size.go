package code

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func GetPathSize(path string, all bool, recursive bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	isDir := fileInfo.IsDir()

	if err := checkHidden(fileInfo.Name(), all, isDir); err != nil {
		return 0, err
	}

	if !isDir {
		return fileInfo.Size(), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return getFilesSize(files, path, all, recursive)
}

func GetSize(path string, human bool, all bool, recursive bool) (string, error) {
	size, err := GetPathSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

func getFilesSize(files []os.DirEntry, path string, all bool, recursive bool) (int64, error) {
	sumSize := int64(0)

	for _, file := range files {
		if file.IsDir() {
			if !recursive {
				continue
			}

			if err := checkHidden(file.Name(), all, true); err != nil {
				continue
			}

			subPath := filepath.Join(path, file.Name())

			subFiles, err := os.ReadDir(subPath)
			if err != nil {
				return 0, errors.New(err.Error())
			}

			subSize, err := getFilesSize(subFiles, subPath, all, recursive)
			if err != nil {
				return 0, errors.New(err.Error())
			}

			sumSize += subSize
			continue
		}

		fileInfo, err := os.Lstat(filepath.Join(path, file.Name()))
		if err != nil {
			return 0, errors.New(err.Error())
		}

		if err := checkHidden(fileInfo.Name(), all, false); err != nil {
			continue
		}

		sumSize += fileInfo.Size()
	}

	return sumSize, nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	s := float64(size)

	switch {
	case size < KB:
		return fmt.Sprintf("%dB", size)
	case size < MB:
		return fmt.Sprintf("%.1fKB", s/KB)
	case size < GB:
		return fmt.Sprintf("%.1fMB", s/MB)
	case size < TB:
		return fmt.Sprintf("%.1fGB", s/GB)
	case size < PB:
		return fmt.Sprintf("%.1fTB", s/TB)
	case size < EB:
		return fmt.Sprintf("%.1fPB", s/PB)
	default:
		return fmt.Sprintf("%.1fEB", s/EB)
	}
}

func checkHidden(path string, all bool, isDir bool) error {
	if !all && strings.HasPrefix(path, ".") {
		errorText := "hidden file"
		if isDir {
			errorText = "hidden directory"
		}
		return errors.New(errorText)
	}

	return nil
}
