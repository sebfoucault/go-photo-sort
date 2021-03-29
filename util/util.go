package util

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// FileInfoPredicate represents a predicate on the os.FileInfo type used to filter in/out files from list of files.
type FileInfoPredicate func(fileInfo os.FileInfo) bool

// IsFileWithExtension returns a FileInfoPredicate that matches the FileInfo matching extensions
func IsFileWithExtension(extensions []string) FileInfoPredicate {

	return func(fileInfo os.FileInfo) bool {
		return !fileInfo.IsDir() && HasExtension(fileInfo.Name(), extensions)
	}
}

// FileMapper represents a mapper transforming components of a file path
type FileMapper interface {
	Map(directoryName string, fileName string) (string, string)
}

// HasExtension returns true if the baseFileName has its extension part of the extensions array, false otherwise
func HasExtension(baseFileName string, extensions []string) bool {

	var lowerExtension = filepath.Ext(baseFileName)
	for _, extension := range extensions {
		if lowerExtension == strings.ToLower(extension) {
			return true
		}
	}

	return false
}

// ListFiles returns all the files under rootDir that matches predicate
func ListFiles(rootDir string, predicate FileInfoPredicate) []string {

	var files []string

	err := filepath.Walk(rootDir, func(filePath string, fileInfo os.FileInfo, err error) error {

		if predicate == nil || predicate(fileInfo) {
			files = append(files, filePath)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

// MapFile maps path to a target directory and a target filename.
func MapFile(path string, mapper FileMapper) (string, string) {

	dirName := filepath.Dir(path)
	fileName := filepath.Base(path)

	mappedDirName, mappedFileName := mapper.Map(dirName, fileName)

	destDirName := strings.ReplaceAll(mappedDirName, "/", string(filepath.Separator))
	destFileName := mappedFileName

	return destDirName, destFileName
}

func CopyFile(sourcePath string, destinationPath string) (int64, error) {

	sourceFileStat, err := os.Stat(sourcePath)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", sourcePath)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destinationDirectory := filepath.Dir(destinationPath)
	err = os.MkdirAll(destinationDirectory, os.ModePerm)
	if err != nil {
		return 0, err
	}

	destination, err := os.Create(destinationPath)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	size, err := io.Copy(destination, source)
	return size, err
}

const ExifDateFormat = "2006:01:02 15:04:05"

func ApplyTemplate(templateName string, templateDefinition string, context interface{}) string {

	template := safeParseTemplate(templateName, templateDefinition)

	var buf bytes.Buffer
	template.Execute(&buf, context)
	return buf.String()
}

func safeParseTemplate(templateName string, templateDefinition string) *template.Template {

	tmpl, e := template.New(templateName).Parse(templateDefinition)
	if e != nil {
		log.Fatal(e)
	}
	return tmpl
}

func ExtractTimeComponents(t time.Time) map[string]string {

	// "2006:01:02 15:04:05"

	m := make(map[string]string)

	m["yy"] = t.Format("06")
	m["yyyy"] = t.Format("2006")
	m["dd"] = t.Format("02")
	m["MM"] = t.Format("01")
	m["MMM"] = t.Format("Jan")
	m["MMMM"] = t.Format("January")
	m["HH"] = t.Format("15")
	m["hh"] = t.Format("03")
	m["mm"] = t.Format("04")
	m["ss"] = t.Format("05")
	m["a"] = t.Format("pm")
	m["EE"] = t.Format("Mon")
	m["EEE"] = t.Format("Mon")
	m["EEEE"] = t.Format("Monday")

	return m
}
