package generator

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	testfile "github.com/dterbah/go-test-gen/core/test"
	file "github.com/dterbah/go-test-gen/utils"
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/sirupsen/logrus"
)

type TestGenerator struct {
	rootPath     string
	testFiles    list.List[string]
	filesCreated int
	filesUpdated int
}

/*
Create a new Test Generator
*/
func NewTestGenerator(rootPath string) TestGenerator {
	return TestGenerator{rootPath: rootPath, testFiles: arraylist.New(comparator.StringComparator)}
}

/*
Generate tests for the project path
*/
func (generator *TestGenerator) GenerateTests() {
	generator.generateTestsForDir(generator.rootPath)

	logrus.Infof("✅ Test generation finished ! Files created %d, files updated %d", generator.filesCreated, generator.filesUpdated)
}

// ---- Private methods ---- //

/*
Generate tests for a specific directory
*/
func (generator *TestGenerator) generateTestsForDir(dir string) {
	entries, err := os.ReadDir(dir)
	dirPath, _ := filepath.Abs(dir)

	logrus.Infof("🚀 Generating tests for directory %s ...", dirPath)

	if err != nil {
		logrus.Errorf("Error when reading the directory %s", dir)
		return
	}

	goFiles := arraylist.New(nil, entries...).Filter(func(element fs.DirEntry) bool {
		return file.IsGoFile(element.Name()) || element.IsDir()
	})

	goFiles.ForEach(func(element fs.DirEntry, index int) {
		currentPath := path.Join(dirPath, element.Name())
		if element.IsDir() {
			generator.generateTestsForDir(currentPath)
		} else {
			generator.generateTestsForFile(currentPath)
		}
	})
}

/*
Generate test for a file
*/
func (generator *TestGenerator) generateTestsForFile(path string) {
	logrus.Infof("🚀 Generating test for file %s ...", path)
	fileTest := testfile.NewFileTest(path)
	status := fileTest.GenerateTests()

	if status == testfile.Created {
		generator.filesCreated++
	} else if status == testfile.Updated {
		generator.filesUpdated++
	}
}
