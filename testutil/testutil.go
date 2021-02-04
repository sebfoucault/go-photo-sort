package testutil

// GetTestImgPath returns the full path of a test image from its base filename
func GetTestImgPath(filename string) string {

	result := "../test-resources/images/" // TODO - to be fixed
	result += filename
	return result
}

// GetTestResourcePath returns the full path of a test resource from its base filename
func GetTestResourcePath(filename string) string {

	result := "../test-resources/"
	result += filename
	return result
}
