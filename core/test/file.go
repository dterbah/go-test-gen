package testfile

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
	"time"

	file "github.com/dterbah/go-test-gen/utils"
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/sirupsen/logrus"
)

const TEST_FUNCTION_EXTENSION = "Test"
const IMPORT_CONTENT = `import("testing")`
const FILE_HEADER = "/* Test autogenerated with the tool go-test-gen */"

type FileTest struct {
	path        string
	functions   list.List[string]
	packageName string
}

type FileTestStatus int

const (
	Created FileTestStatus = iota
	Updated
	Error
)

/*
Create new file test case
*/
func NewFileTest(path string) FileTest {
	return FileTest{path: path, functions: arraylist.New(comparator.StringComparator)}
}

/*
Generate the tests for a specific file
*/
func (fileTest FileTest) GenerateTests(includePrivateFunction bool) FileTestStatus {
	// find functions of the current file
	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, fileTest.path, nil, parser.AllErrors)

	if err != nil {
		logrus.Errorf("Error when parsing the file %s", fileTest.path)
		return Error
	}

	fileTest.packageName = node.Name.Name

	for _, decl := range node.Decls {
		if function, isFunction := decl.(*ast.FuncDecl); isFunction {
			functionName := function.Name.Name
			isPrivate := isPrivateFunction(functionName)
			if !isPrivate || (includePrivateFunction && isPrivate) {
				fileTest.functions.Add(functionName)
			}
		}
	}

	return fileTest.writeTests()
}

/*
Write the tests corresponding to the file
*/
func (fileTest FileTest) writeTests() FileTestStatus {
	testFileContent := fmt.Sprintf("%s\n\n%s\n\n%s\n", getPackageContent(fileTest.packageName), getFileHeader(), IMPORT_CONTENT)
	fileTestPath := file.CreateTestFilePath(fileTest.path)
	status := Created

	existingTestFunctions := arraylist.New(comparator.StringComparator)
	if file.Exists(fileTestPath) {
		status = Updated
		testFileContent = ""
		// parse this file and retrieve test functions name
		fileSet := token.NewFileSet()
		node, err := parser.ParseFile(fileSet, fileTestPath, nil, parser.AllErrors)
		if err != nil {
			logrus.Errorf("Error when parsing the file %s", fileTestPath)
			return Error
		}

		for _, decl := range node.Decls {
			if function, isFunction := decl.(*ast.FuncDecl); isFunction {
				functionName := function.Name.Name
				existingTestFunctions.Add(functionName)
			}
		}
	}

	fileTest.functions.ForEach(func(function string, index int) {
		// not override existing test functions
		if !existingTestFunctions.Contains(getTestFunctionName(function)) {
			testFunctionContent := getTestFunctionContent(function)
			testFileContent = fmt.Sprintf("%s\n%s\n", testFileContent, testFunctionContent)
		}
	})

	// write new file content in the file test

	file, err := os.OpenFile(fileTestPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		logrus.Errorf("Error when opening file %s", fileTestPath)
	}

	if _, err := file.Write([]byte(testFileContent)); err != nil {
		logrus.Errorf("Error when writing into file %s", fileTestPath)
	}

	if err := file.Close(); err != nil {
		logrus.Errorf("Error when closing file %s", fileTestPath)
	}

	if err != nil {
		return Error
	}

	if status == Created {
		logrus.Infof("✅ File test %s created !", fileTestPath)
	} else {
		logrus.Infof("✅ File test %s updated !", fileTestPath)
	}

	return status
}

/*
Return a generated function content
*/
func getTestFunctionContent(function string) string {
	testFunctionName := getTestFunctionName(function)
	return fmt.Sprintf("func %s(t *testing.T) {\n	// todo implement \n}", testFunctionName)
}

/*
Return test function name according tp his original name
*/
func getTestFunctionName(function string) string {
	return fmt.Sprintf("%s%s", TEST_FUNCTION_EXTENSION, function)
}

/*
Return the package content for the test file generated
*/
func getPackageContent(packageName string) string {
	return fmt.Sprintf("package %s", packageName)
}

/*
Return the test file header
*/
func getFileHeader() string {
	currentTime := time.Now()
	today := currentTime.Format("2006-01-02 15:04:05 Monday")

	return fmt.Sprintf("/* Test autogenerated with the tool go-test-gen. Created %s */", today)
}

/*
Check if a function is private or not, based on her name
*/
func isPrivateFunction(functionName string) bool {
	return strings.ToLower(string(functionName[0])) == string(functionName[0])
}
