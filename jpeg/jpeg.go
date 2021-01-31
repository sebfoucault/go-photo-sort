package jpeg

import (
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure"
)

func GetExifData(file string) []byte {

	jmp := jpegstructure.NewJpegMediaParser()
	mediaContext, err := jmp.ParseFile(file)
	if err != nil {

	}
	_, data, _ := mediaContext.Exif()
	return data
}
