package fileformats

import (
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
)

// GetJpegExifData returns the Exif metadata as an array of bytes
func GetJpegExifData(file string) ([]byte, error) {

	jmp := jpegstructure.NewJpegMediaParser()
	mediaContext, err := jmp.ParseFile(file)
	if err != nil {
		return nil, err
	}
	_, data, _ := mediaContext.Exif()
	return data, nil
}
