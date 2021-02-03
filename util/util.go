package util

import (
	"os"
	"path/filepath"
	"strings"
)

// FileInfoPredicate represents a predicate on the os.FileInfo type
type FileInfoPredicate func(fileInfo os.FileInfo) bool

// HasExtension returns true if the baseFileName has its extension part of the extensions array, false otherwise
func HasExtension(baseFileName string, extensions []string) bool {

	var lowerBaseFileName = strings.ToLower(baseFileName)
	for _, extension := range extensions {
		if strings.HasSuffix(lowerBaseFileName, strings.ToLower(extension)) {
			return true
		}
	}

	return false
}

// ListFiles returns all the files under rootDir that matches predicate
func ListFiles(rootDir string, predicate FileInfoPredicate) []string {

	var files []string

	err := filepath.Walk(rootDir, func(filePath string, fileInfo os.FileInfo, err error) error {

		if predicate(fileInfo) {
			files = append(files, filePath)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}
