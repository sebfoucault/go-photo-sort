package testutil

// TestImgPath returns the full path of a test image from its base filename
func TestImgPath(filename string) string {

	result := "../test-resources/images/" // TODO - to be fixed
	result += filename
	return result
}
