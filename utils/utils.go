package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type FileInfoPredicate func(fileInfo os.FileInfo) bool

//
// Returns true if the baseFileName has its extension part of the extensions array, false otherwise
func HasExtension(baseFileName string, extensions []string) bool {

	for _, extension := range extensions {
		var lowerBaseFileName = strings.ToLower(baseFileName)
		if strings.HasSuffix(lowerBaseFileName, extension) {
			return true
		}
	}

	return false
}

//
// Lists all the files under rootDir that matches predicate
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
