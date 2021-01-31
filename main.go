package main

import (
	"fmt"
	"os"

	"./tagreader"
	"./utils"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

func main() {

	rootDir := "."

	result := utils.ListFiles(rootDir, isFileWithExtension([]string{".jpg"}))

	for _, file := range result {

		jmp := jpegstructure.NewJpegMediaParser()
		mediaContext, err := jmp.ParseFile(file)
		if err != nil {

		}
		_, data, _ := mediaContext.Exif()

		tagReader := tagreader.CreateExifTagReader(data)
		fmt.Println(tagReader.GetTag())
	}
}

func isFileWithExtension(extensions []string) utils.FileInfoPredicate {

	return func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir() && utils.HasExtension(fileInfo.Name(), extensions)
	}
}
