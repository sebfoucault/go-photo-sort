package tagreader

import (
	"time"

	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	log "github.com/dsoprea/go-logging/v2"
)

// ExifTagReader abstracts a type for reading Exif metadata
type ExifTagReader struct {
	rootIfd *exif.Ifd
}

// Tag abstracts the tags returned by GetAllTags
type Tag struct {
	Name  string
	Value string
	Path  string
}

// CreateExifTagReader creates an ExifagReader instance used to read the Exif metadata passed in parameter
func CreateExifTagReader(data []byte) *ExifTagReader {

	im, err := exifcommon.NewIfdMappingWithStandard()
	log.PanicIf(err)

	tagIndex := exif.NewTagIndex()

	_, index, err := exif.Collect(im, tagIndex, data)
	log.PanicIf(err)

	rootIfd := index.RootIfd

	return &ExifTagReader{rootIfd}
}

// GetTagByName does a lookup for an Exif tag based on the tag name
func (exifTagReader *ExifTagReader) GetTagByName(tagName string) string {

	// We know the tag we want is on IFD0 (the first/root IFD).
	results, err := exifTagReader.rootIfd.FindTagWithName(tagName)
	log.PanicIf(err)

	// This should never happen.
	if len(results) != 1 {
		log.Panicf("there wasn't exactly one result")
	}

	return getValueFromEntry(results[0])
}

// GetTagByID does a lookup for an Exif tag based on the tag id
func (exifTagReader *ExifTagReader) GetTagByID(tagID uint16) string {

	// We know the tag we want is on IFD0 (the first/root IFD).
	results, err := exifTagReader.rootIfd.FindTagWithId(tagID)
	log.PanicIf(err)

	// This should never happen.
	if len(results) != 1 {
		log.Panicf("there wasn't exactly one result")
	}

	return getValueFromEntry(results[0])
}

// GetAllTags gets all the tags as a map associating tag name to tag value (as string)
func (exifTagReader *ExifTagReader) GetAllTags() map[string]Tag {

	// Creates the resulting map
	result := make(map[string]Tag)

	// Creates the iterator that will add an entry in the map for each tag
	visitor := func(ifd *exif.Ifd, tagEntry *exif.IfdTagEntry) error {

		tagName := tagEntry.TagName()
		tagValue := getValueFromEntry(tagEntry)

		// Adds the standard tag in the Exif map
		result[tagName] = Tag{
			Name:  tagName,
			Value: tagValue,
			Path:  tagEntry.IfdPath(),
		}

		return nil
	}

	// Iterates on the tag
	exifTagReader.rootIfd.EnumerateTagsRecursively(visitor)

	return result
}

func getValueFromEntry(entry *exif.IfdTagEntry) string {

	result, err := entry.FormatFirst()
	log.PanicIf(err)
	return result
}

// exifDateLayout specifies the format (aka. layout) to be used for parsing Exif dates with time.Parse.
// This format is specified in "Exchangeable image file format for digital still cameras: Exif Version 2.2"
const exifDateLayout = "2006:01:02 15:04:05"

var exifDateTags = map[string]bool{
	"DateTimeOriginal":  true,
	"DateTimeDigitized": true,
}

// GetStandardExifDateTags returns an array containing the name of standard exif tag representing date
func GetStandardExifDateTags() []string {

	s := make([]string, len(exifDateTags))
	for k := range exifDateTags {
		s = append(s, k)
	}
	return s
}

// IsDateExifTag returns true if the tagName passed in parameter is a date tag
func IsDateExifTag(tagName string) bool {

	if _, ok := exifDateTags[tagName]; ok {
		return true
	}
	return false
}

// ParseExifDate returns a time.Time from the value of an Exif tag representing a date
func ParseExifDate(s string) (time.Time, error) {
	result, error := time.Parse(exifDateLayout, s)
	return result, error
}
