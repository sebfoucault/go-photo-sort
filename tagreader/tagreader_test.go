package tagreader_test

import (
	"testing"

	"github.com/sebfoucault/go-photo-sort/fileformats"
	"github.com/sebfoucault/go-photo-sort/tagreader"
	"github.com/sebfoucault/go-photo-sort/testutil"

	log "github.com/dsoprea/go-logging/v2"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestCreateReader(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("mandala-small.jpg"))
	then.AssertThat(t, reader, is.Not(is.Nil()))
}

func TestGetTagByName(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("mandala-small.jpg"))
	width := reader.GetTagByName("Model")
	then.AssertThat(t, width, is.EqualTo("DSC-RX100M2"))
}

func TestGetTagByID(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("mandala-small.jpg"))
	width := reader.GetTagByID(0x0110)
	then.AssertThat(t, width, is.EqualTo("DSC-RX100M2"))
}

func TestGetAllTags(t *testing.T) {

	reader := createReaderForFile(testutil.GetTestImgPath("mandala-small.jpg"))
	tags := reader.GetAllTags()
	then.AssertThat(t, tags, has.Length(46))
}

func createReaderForFile(file string) *tagreader.ExifTagReader {
	// Create reader for a test file
	return tagreader.CreateExifTagReader(getExifData(file))
}

func getExifData(file string) []byte {
	data, err := fileformats.GetJpegExifData(file)
	log.PanicIf(err)
	return data
}
