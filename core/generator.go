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
	config       TestGeneratorConfig
}

/*
Create a new Test Generator
*/
func NewTestGenerator(rootPath string) TestGenerator {
	config := LoadTestGeneratorConfig(rootPath)
	return TestGenerator{
		rootPath:  rootPath,
		testFiles: arraylist.New(comparator.StringComparator),
		config:    config,
	}
}

/*
Generate tests for the project path
*/
func (generator *TestGenerator) GenerateTests() {
	generator.generateTestsForDir(generator.rootPath)

	logrus.Infof("✅ Tests generation finished ! Files created %d, files updated %d", generator.filesCreated, generator.filesUpdated)
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
		name := element.Name()
		currentPath := path.Join(dirPath, name)
		if element.IsDir() {
			generator.generateTestsForDir(currentPath)
		} else {
			// check if we should include this file
			shouldInclude := generator.shouldIncludeFile(name)
			if !shouldInclude {
				logrus.Infof("🟠 %s ignored !", name)
			} else {
				generator.generateTestsForFile(currentPath)
			}
		}
	})
}

/*
Generate test for a file
*/
func (generator *TestGenerator) generateTestsForFile(path string) {
	logrus.Infof("🚀 Generating tests for file %s ...", path)
	fileTest := testfile.NewFileTest(path)
	status := fileTest.GenerateTests()

	if status == testfile.Created {
		generator.filesCreated++
	} else if status == testfile.Updated {
		generator.filesUpdated++
	}
}

/*
Verify if a filename should be included during the tests generation.
Return true if at least one rule is verified, else false
*/
func (generator TestGenerator) shouldIncludeFile(filename string) bool {
	for _, excludeFileRegex := range generator.config.ExcludeFiles {
		matched, _ := filepath.Match(excludeFileRegex, filename)
		if matched {
			return false
		}
	}

	return true
}
