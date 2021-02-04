package tagreader_test

import (
	"testing"

	"github.com/sebfoucault/go-photo-sort/jpeg"
	"github.com/sebfoucault/go-photo-sort/tagreader"
	"github.com/sebfoucault/go-photo-sort/testutil"

	log "github.com/dsoprea/go-logging/v2"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestCreateReader(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("scotland-nicolas-boulesteix.jpg"))
	then.AssertThat(t, reader, is.Not(is.Nil()))
}

func TestGetTagByName(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("scotland-nicolas-boulesteix.jpg"))
	width := reader.GetTagByName("ImageWidth")
	then.AssertThat(t, width, is.EqualTo("1920"))
}

func TestGetTagByID(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("scotland-nicolas-boulesteix.jpg"))
	width := reader.GetTagByID(256)
	then.AssertThat(t, width, is.EqualTo("1920"))
}

func TestGetAllTags(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("scotland-nicolas-boulesteix.jpg"))
	tags := reader.GetAllTags()
	then.AssertThat(t, tags, has.Length(42))
}

func createReaderForFile(file string) *tagreader.ExifTagReader {
	// Create reader for a test file
	return tagreader.CreateExifTagReader(getExifData(file))
}

func getExifData(file string) []byte {
	data, err := jpeg.GetExifData(file)
	log.PanicIf(err)
	return data
}
