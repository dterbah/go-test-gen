package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExists(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"file doesn't exists", "toto", false},
		{"file exists", "file.go", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Exists(test.path)
			assert.Equal(test.expected, result)
		})
	}
}

func TestIsDir(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"not dir", "file.go", false},
		{"not exists", "toto", false},
		{"is dir", ".", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsDir(test.path)
			assert.Equal(test.expected, result)
		})
	}
}

func TestIsGoFile(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{"not go file", "file.gogo", false},
		{"not go file", "file._go", false},
		{"is go file", "toto.go", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsGoFile(test.path)
			assert.Equal(test.expected, result)
		})
	}
}

func TestCreateTestFilePath(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name     string
		path     string
		expected string
	}{
		{"test go file generated", "file.go", "file_test.go"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CreateTestFilePath(test.path)
			assert.Equal(test.expected, result)
		})
	}
}
