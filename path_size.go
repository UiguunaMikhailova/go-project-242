package code

import (
	"errors"
	"fmt"
	"os"
)

func GetSize(path string) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", errors.New(err.Error())
	}

	isDir := fileInfo.IsDir()

	if isDir {
		files, err := os.ReadDir(path)
		if err != nil {
			return "", errors.New(err.Error())
		}

		size := int64(0)

		for _, file := range files {
			sizeOfFile, err := getSizeOfFile(path + "/" + file.Name())
			if err != nil {
				return "", errors.New(err.Error())
			}
			size += sizeOfFile
			fmt.Println("File in dir: ", file.Name())
		}

		return fmt.Sprintf("%dB", size), nil
	}

	fmt.Println("FileInfo: ", fileInfo.Name(), fileInfo.Size(), fileInfo.IsDir())

	size, err := getSizeOfFile(path)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return fmt.Sprintf("%dB", size), nil
}

func getSizeOfFile(path string) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	if fileInfo.IsDir() {
		// return 0, errors.New("is a directory")
		return getSizeOfFile(path)
	}

	return fileInfo.Size(), nil
}
