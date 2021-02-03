package util_test

import (
	"testing"

	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/sebfoucault/go-photo-sort/util"
)

func TestHasExtension(t *testing.T) {

	then.AssertThat(t, util.HasExtension("test.jpg", []string{".jpg", ".jpeg"}), is.True())
	then.AssertThat(t, util.HasExtension("test.jpg", []string{".JPG", ".JPEG"}), is.True())
	then.AssertThat(t, util.HasExtension("test.png", []string{".jpg", ".tiff"}), is.False())
}
