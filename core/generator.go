package generator

import (
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
)

type TestGenerator struct {
	rootPath  string
	testFiles list.List[string]
}

func NewTestGenerator(rootPath string) *TestGenerator {
	return &TestGenerator{rootPath: rootPath, testFiles: arraylist.New(comparator.StringComparator)}
}
