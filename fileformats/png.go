package fileformats

import pngstructure "github.com/dsoprea/go-png-image-structure/v2"

// GetPngExifData returns the Exif metadata as an array of bytes
func GetPngExifData(file string) ([]byte, error) {

	jmp := pngstructure.NewPngMediaParser()
	mediaContext, err := jmp.ParseFile(file)
	if err != nil {
		return nil, err
	}
	_, data, _ := mediaContext.Exif()
	return data, nil
}
