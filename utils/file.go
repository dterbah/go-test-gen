package file

import (
	"fmt"
	"os"
	"strings"
)

const GOLANG_FILE_EXTENSION = ".go"
const TEST_EXTENSION = "_test"

/*
Check if a path is corresponding to a file or dir
*/
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*
Return true if the path corresponds to a directory, else false
*/
func IsDir(path string) bool {
	file, err := os.Stat(path)

	if err != nil {
		return false
	}

	return file.IsDir()
}

/*
Return true if the path has the golang extension, else false
*/
func IsGoFile(path string) bool {
	return strings.HasSuffix(path, GOLANG_FILE_EXTENSION) && !strings.Contains(path, TEST_EXTENSION)
}

/*
Read the content of file
*/
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CreateTestFilePath(path string) string {
	index := strings.Index(path, GOLANG_FILE_EXTENSION)
	filename := fmt.Sprintf("%s_test%s", path[:index], GOLANG_FILE_EXTENSION)
	return filename
}
