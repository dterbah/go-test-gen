package testfile

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/sirupsen/logrus"
)

const TEST_FUNCTION_EXTENSION = "Test"

type FileTest struct {
	path        string
	functions   list.List[string]
	packageName string
}

/*
Create new file test case
*/
func NewFileTest(path string) FileTest {
	return FileTest{path: path, functions: arraylist.New(comparator.StringComparator)}
}

/*
Generate the tests for a specific file
*/
func (fileTest FileTest) GenerateTests() {
	// find functions of the current file
	fileSet := token.NewFileSet()

	node, err := parser.ParseFile(fileSet, fileTest.path, nil, parser.AllErrors)

	if err != nil {
		logrus.Errorf("Error when parsing the file %s", fileTest.path)
		return
	}

	fileTest.packageName = node.Name.Name

	for _, decl := range node.Decls {
		if function, isFunction := decl.(*ast.FuncDecl); isFunction {
			functionName := function.Name.Name
			fileTest.functions.Add(functionName)
		}
	}

	fileTest.writeTests()
}

/*
Write the tests corresponding to the file
*/
func (fileTest FileTest) writeTests() {
	//testFileName := file.CreateTestFile(fileTest.path)
	testFileContent := fmt.Sprintf("%s\n\n%s\n\n%s\n", getPackageContent(fileTest.packageName), getTestFileHeader(), getImportsForTestFile())

	// todo: don't override file content

	fileTest.functions.ForEach(func(function string, index int) {
		testFunctionContent := getTestFunctionContent(function)
		testFileContent = fmt.Sprintf("%s\n %s", testFileContent, testFunctionContent)
	})

	logrus.Info(testFileContent)
}

func getTestFileHeader() string {
	return "/* Test autogenerated with the tool go-test-gen */"
}

func getImportsForTestFile() string {
	return `import("testing")`
}

func getTestFunctionContent(function string) string {
	testFunctionName := fmt.Sprintf("%s%s", function, TEST_FUNCTION_EXTENSION)
	return fmt.Sprintf("func %s(t *t.Testing) {\n }", testFunctionName)
}

func getPackageContent(packageName string) string {
	return fmt.Sprintf("package %s", packageName)
}
