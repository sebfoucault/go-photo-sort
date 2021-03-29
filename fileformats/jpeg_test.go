package fileformats_test

import (
	"testing"

	"github.com/sebfoucault/go-photo-sort/fileformats"
	"github.com/sebfoucault/go-photo-sort/testutil"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestGetExifData(t *testing.T) {

	data, _ := fileformats.GetJpegExifData(testutil.GetTestImgPath("mandala-small.jpg"))

	then.AssertThat(t, data, is.Not(is.Nil()))
}

func TestGetExifDataWithNonExistingFile(t *testing.T) {

	data, err := fileformats.GetJpegExifData(testutil.GetTestImgPath("does-not-exist.jpg"))

	then.AssertThat(t, data, is.Nil())
	then.AssertThat(t, err, is.Not(is.Nil()))
}
