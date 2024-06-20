package generator

import (
	testfile "github.com/dterbah/go-test-gen/core/test"
	"github.com/sirupsen/logrus"
)

/*
Test generator for a specific file
*/
type FileTestGenerator struct {
	BaseTestGenerator
	filePath string
}

/*
Create a new file test generator
*/
func NewFileTestGenerator(path string) *FileTestGenerator {
	return &FileTestGenerator{
		filePath: path,
	}
}

func (generator FileTestGenerator) GenerateTests() {
	logrus.Infof("🚀 Generating tests for file %s ...", generator.filePath)
	fileTest := testfile.NewFileTest(generator.filePath)
	status := fileTest.GenerateTests(false)

	if status == testfile.Created {
		logrus.Infof("Tests for file %s created successfully !", generator.filePath)
	} else if status == testfile.Updated {
		logrus.Infof("Tests for file %s updated successfully !", generator.filePath)
	}
}
