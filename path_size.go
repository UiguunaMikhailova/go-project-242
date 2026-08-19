package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetPathSize returns the size of the file or directory at path.
//
// For a directory the sizes of its entries are summed; nested directories are
// traversed only when recursive is set. When human is set the size is rendered
// in readable form ("7.8KB"), otherwise in bytes ("8000B"). The all flag affects
// only the traversal of contents: hidden entries (names starting with a dot) are
// counted only when all is set, but path itself is always measured, even if it
// is hidden.
//
// An error is returned only if path itself cannot be read.
func GetPathSize(path string, recursive bool, human bool, all bool) (string, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("error processing %s: %w", path, err)
	}

	if !fileInfo.IsDir() {
		return formatSize(fileInfo.Size(), human), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("error processing %s: %w", path, err)
	}

	return formatSize(getFilesSize(files, path, all, recursive), human), nil
}

func getFilesSize(files []os.DirEntry, path string, all bool, recursive bool) int64 {
	sumSize := int64(0)

	for _, file := range files {
		if !all && isHidden(file.Name()) {
			continue
		}

		subPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			if !recursive {
				continue
			}

			subFiles, err := os.ReadDir(subPath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: can't read %s: %v\n", subPath, err)
				continue
			}

			sumSize += getFilesSize(subFiles, subPath, all, recursive)
			continue
		}

		fileInfo, err := os.Lstat(subPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: can't stat %s: %v\n", subPath, err)
			continue
		}

		sumSize += fileInfo.Size()
	}

	return sumSize
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	const (
		kb = 1024
		mb = 1024 * kb
		gb = 1024 * mb
		tb = 1024 * gb
		pb = 1024 * tb
		eb = 1024 * pb
	)

	s := float64(size)

	switch {
	case size < kb:
		return fmt.Sprintf("%dB", size)
	case size < mb:
		return fmt.Sprintf("%.1fKB", s/kb)
	case size < gb:
		return fmt.Sprintf("%.1fMB", s/mb)
	case size < tb:
		return fmt.Sprintf("%.1fGB", s/gb)
	case size < pb:
		return fmt.Sprintf("%.1fTB", s/tb)
	case size < eb:
		return fmt.Sprintf("%.1fPB", s/pb)
	default:
		return fmt.Sprintf("%.1fEB", s/eb)
	}
}

func isHidden(name string) bool {
	return strings.HasPrefix(name, ".")
}
