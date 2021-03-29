package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sebfoucault/go-photo-sort/fileformats"
	"github.com/sebfoucault/go-photo-sort/tagreader"
	"github.com/sebfoucault/go-photo-sort/util"
	"github.com/urfave/cli/v2"
)

func main() {

	var inputDirectory string
	var outputDirectory string
	var fileNameTemplate string
	var directoryNameTemplate string

	app := &cli.App{
		Name:  "go-photo-sort",
		Usage: "Sort your photo",
		Action: func(c *cli.Context) error {

			cfg := config{
				fileNameTemplate:      fileNameTemplate,
				directoryNameTemplate: directoryNameTemplate,
			}

			doIt(inputDirectory, outputDirectory, cfg)
			return nil
		},
	}

	app.Flags = []cli.Flag{
		&cli.PathFlag{
			Destination: &inputDirectory,
			Name:        "input",
			Aliases:     []string{"i"},
			Usage:       "Input directory - The directory where to read the photos to be sorted",
			Required:    true,
		},
		&cli.PathFlag{
			Destination: &outputDirectory,
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "Output directory - The directory where to write the sorted photos",
			Required:    true,
		},
		&cli.StringFlag{
			Destination: &fileNameTemplate,
			Name:        "file-template",
			Aliases:     []string{"ft"},
			Usage:       "Template for output filename",
			Value:       "{{.yyyy}}-{{.MM}}-{{.dd}}-{{.HH}}-{{.mm}}-{{.ss}}.{{.ext}}",
		},
		&cli.StringFlag{
			Destination: &directoryNameTemplate,
			Name:        "dir-template",
			Aliases:     []string{"dt"},
			Usage:       "Template for output directory name",
			Value:       "{{.yyyy}}/{{.MM}}",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}

func doIt(inputDirectory string, outputDirectory string, cfg config) {

	result := util.ListFiles(inputDirectory, util.IsFileWithExtension([]string{".jpg"}))

	for _, file := range result {

		exifData, _ := fileformats.GetJpegExifData(file)
		exifTagReader := tagreader.CreateExifTagReader(exifData)

		mapper := createMapper(file, exifTagReader, cfg)
		destDirectoryName, destFileName := util.MapFile(file, mapper)

		destination := filepath.Join(outputDirectory, destDirectoryName, destFileName)
		_, err := util.CopyFile(file, destination)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createMapper(file string, exifTagReader *tagreader.ExifTagReader, cfg config) util.FileMapper {

	exifTags := exifTagReader.GetAllTags()

	//
	// Creates the context for pattern application

	context := make(map[string]string)

	//
	// Adds the date fields (like year, month)

	time, _ := extractTimeFromTags(exifTags, tagreader.GetStandardExifDateTags())
	timeComponents := util.ExtractTimeComponents(time)
	for k, v := range timeComponents {
		context[k] = v
	}

	//
	// Adds the file extension

	ext := filepath.Ext(file)
	if ext != "" {
		ext = ext[1:] // Remove the leading dot
	}
	context["ext"] = ext

	//
	// Creates the mappers

	return &exifMapper{
		context: context,
		cfg:     cfg,
	}
}

type exifMapper struct {
	context map[string]string
	cfg     config
}

func (mapper *exifMapper) Map(directoryName string, fileName string) (string, string) {

	mappedDirectoryName := util.ApplyTemplate("directoryNameTemplate", mapper.cfg.directoryNameTemplate, mapper.context)
	mappedFileName := util.ApplyTemplate("fileNameTemplate", mapper.cfg.fileNameTemplate, mapper.context)

	return mappedDirectoryName, mappedFileName
}

type config struct {
	fileNameTemplate      string
	directoryNameTemplate string
}

func extractTimeFromTags(exifTags map[string]tagreader.Tag, tagNamesByPrecedence []string) (time.Time, error) {

	for _, tagName := range tagNamesByPrecedence {

		if tag, hasTag := exifTags[tagName]; hasTag {
			time, err := tagreader.ParseExifDate(tag.Value)
			if err != nil {
				log.Fatal(err) // TODO : maybe should log and go try next instead
			}
			return time, nil
		}
	}
	return time.Time{}, nil
}
