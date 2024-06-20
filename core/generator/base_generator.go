package generator

/*
Represents the methods used by all TestGenerator
*/
type BaseTestGenerator interface {
	/*
		Generate tests for a project or a file, depending of the implementation
	*/
	GenerateTests()
}
