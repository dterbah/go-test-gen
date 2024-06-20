package generator

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/dterbah/go-test-gen/core/config"
	testfile "github.com/dterbah/go-test-gen/core/test"
	file "github.com/dterbah/go-test-gen/utils"
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/sirupsen/logrus"
)

/*
Represents a project test generator
*/
type ProjectTestGenerator struct {
	BaseTestGenerator
	rootPath     string
	testFiles    list.List[string]
	filesCreated int
	filesUpdated int
	config       config.TestGeneratorConfig
}

/*
Create a new Test Generator
*/
func NewProjectTestGenerator(rootPath string) *ProjectTestGenerator {
	config := config.LoadTestGeneratorConfig(rootPath)
	if !config.Verbose {
		logrus.SetLevel(logrus.PanicLevel)
	}

	return &ProjectTestGenerator{
		rootPath:  rootPath,
		testFiles: arraylist.New(comparator.StringComparator),
		config:    config,
	}
}

/*
Generate tests for the project path
*/
func (generator *ProjectTestGenerator) GenerateTests() {
	generator.generateTestsForDir(generator.rootPath)
	// reset the log level
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Infof("✅ Tests generation finished ! Files created %d, files updated %d", generator.filesCreated, generator.filesUpdated)
}

// ---- Private methods ---- //

/*
Generate tests for a specific directory
*/
func (generator *ProjectTestGenerator) generateTestsForDir(dir string) {
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
func (generator *ProjectTestGenerator) generateTestsForFile(path string) {
	logrus.Infof("🚀 Generating tests for file %s ...", path)
	fileTest := testfile.NewFileTest(path)
	status := fileTest.GenerateTests(generator.config.GeneratePrivateFunctions)

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
func (generator ProjectTestGenerator) shouldIncludeFile(filename string) bool {
	for _, excludeFileRegex := range generator.config.ExcludeFiles {
		matched, _ := filepath.Match(excludeFileRegex, filename)
		if matched {
			return false
		}
	}

	return true
}
