package util_test

import (
	"os"
	"testing"

	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/sebfoucault/go-photo-sort/testutil"
	"github.com/sebfoucault/go-photo-sort/util"
)

func TestHasExtension(t *testing.T) {

	then.AssertThat(t, util.HasExtension("test.jpg", []string{".jpg", ".jpeg"}), is.True())
	then.AssertThat(t, util.HasExtension("test.jpg", []string{".JPG", ".JPEG"}), is.True())
	then.AssertThat(t, util.HasExtension("test.png", []string{".jpg", ".tiff"}), is.False())
}

func TestListFiles(t *testing.T) {

	root := testutil.GetTestResourcePath("dirs/hierarchy-01")

	allFiles := util.ListFiles(root, func(fileInfo os.FileInfo) bool {
		return true
	})
	then.AssertThat(t, allFiles, has.Length(9)) // TODO - add more assertions
}
