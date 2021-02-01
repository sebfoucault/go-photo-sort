package tagreader

import (
	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	log "github.com/dsoprea/go-logging/v2"
)

type ExifTagReader struct {
	rootIfd *exif.Ifd
}

func CreateExifTagReader(data []byte) ExifTagReader {

	im, err := exifcommon.NewIfdMappingWithStandard()
	log.PanicIf(err)

	tagIndex := exif.NewTagIndex()

	_, index, err := exif.Collect(im, tagIndex, data)
	log.PanicIf(err)

	rootIfd := index.RootIfd

	return ExifTagReader{rootIfd}
}

func (exifTagReader *ExifTagReader) GetTag() string {

	// We know the tag we want is on IFD0 (the first/root IFD).
	results, err := exifTagReader.rootIfd.FindTagWithId(306)
	// results, err := rootIfd.FindTagWithName(tagName)

	log.PanicIf(err)

	// This should never happen.
	if len(results) != 1 {
		log.Panicf("there wasn't exactly one result")
	}

	ite := results[0]

	valueRaw, err := ite.Value()
	log.PanicIf(err)

	value := valueRaw.(string)
	return value
}
