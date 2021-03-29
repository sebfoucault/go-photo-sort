package util_test

import (
	"os"
	"testing"
	"time"

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

func TestExtractTimeComponents(t *testing.T) {

	time, _ := time.Parse("2006:01:02 15:04:05", "2006:01:02 15:04:05")

	componentsMaps := util.ExtractTimeComponents(time)
	then.AssertThat(t, componentsMaps["yy"], is.EqualTo("06"))
	then.AssertThat(t, componentsMaps["yyyy"], is.EqualTo("2006"))
	then.AssertThat(t, componentsMaps["MM"], is.EqualTo("01"))
	then.AssertThat(t, componentsMaps["MMM"], is.EqualTo("Jan"))
	then.AssertThat(t, componentsMaps["MMMM"], is.EqualTo("January"))
	then.AssertThat(t, componentsMaps["HH"], is.EqualTo("15"))
	then.AssertThat(t, componentsMaps["hh"], is.EqualTo("03"))
	then.AssertThat(t, componentsMaps["mm"], is.EqualTo("04"))
	then.AssertThat(t, componentsMaps["ss"], is.EqualTo("05"))
	then.AssertThat(t, componentsMaps["a"], is.EqualTo("pm"))
	then.AssertThat(t, componentsMaps["EE"], is.EqualTo("Mon"))
	then.AssertThat(t, componentsMaps["EEE"], is.EqualTo("Mon"))
	then.AssertThat(t, componentsMaps["EEEE"], is.EqualTo("Monday"))
}
