package main

import (
	"fmt"
	"os"

	"./jpeg"
	"./tagreader"
	"./utils"
)

func main() {

	rootDir := "."

	result := utils.ListFiles(rootDir, isFileWithExtension([]string{".jpg"}))

	for _, file := range result {

		data := jpeg.GetExifData(file)
		tagReader := tagreader.CreateExifTagReader(data)
		fmt.Println(tagReader.GetTag())
	}
}

func isFileWithExtension(extensions []string) utils.FileInfoPredicate {

	return func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir() && utils.HasExtension(fileInfo.Name(), extensions)
	}
}
