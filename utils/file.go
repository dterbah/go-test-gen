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
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
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

func CreateTestFilePath(path string) string {
	index := strings.Index(path, GOLANG_FILE_EXTENSION)
	filename := fmt.Sprintf("%s_test%s", path[:index], GOLANG_FILE_EXTENSION)
	return filename
}
