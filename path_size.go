package code

import (
	"errors"
	"fmt"
	"os"
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

func GetSize(path string, human bool, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", errors.New(err.Error())
	}

	fmt.Println("FileInfo: ", fileInfo.Name(), fileInfo.Size(), fileInfo.IsDir())

	isDir := fileInfo.IsDir()

	errHidden := checkHidden(fileInfo.Name(), all, isDir)
	if errHidden != nil {
		return "", errors.New(errHidden.Error())
	}

	if isDir {
		return dirSize(path, human, all)
	}

	return formatSize(fileInfo.Size(), human), nil
}

func dirSize(path string, human, all bool) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", errors.New(err.Error())
	}

	sumSize := int64(0)

	for _, file := range files {
		fileInfo, err := os.Lstat(path + "/" + file.Name())
		if err != nil {
			return "", errors.New(err.Error())
		}

		fmt.Println("File in dir: ", file.Name())

		errHidden := checkHidden(fileInfo.Name(), all, true)
		if errHidden != nil {
			continue
		}

		sumSize += fileInfo.Size()
	}

	return formatSize(sumSize, human), nil
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
