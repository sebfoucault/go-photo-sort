package main

import (
	"fmt"
	"os"

	"github.com/sebfoucault/go-photo-sort/jpeg"
	"github.com/sebfoucault/go-photo-sort/tagreader"
	"github.com/sebfoucault/go-photo-sort/util"
)

func main() {

	rootDir := "."

	result := util.ListFiles(rootDir, isFileWithExtension([]string{".jpg"}))

	for _, file := range result {

		data, _ := jpeg.GetExifData(file)

		tagReader := tagreader.CreateExifTagReader(data)
		tags := tagReader.GetAllTags()
		fmt.Println(tags)
	}
}

func isFileWithExtension(extensions []string) util.FileInfoPredicate {

	return func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir() && util.HasExtension(fileInfo.Name(), extensions)
	}
}
