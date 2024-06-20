package cmd

import (
	"github.com/dterbah/go-test-gen/core/generator"
	file "github.com/dterbah/go-test-gen/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectPath string
var filePath string

// generateCmd représente la commande generate
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate tests for the specified project",
	Long: `Generate unit tests for the specified Go project.
You need to provide the path to the project using the --project flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		var testGenerator generator.BaseTestGenerator
		if projectPath != "" && filePath != "" {
			logrus.Error("You cannot use --project and --file options at the same time. Please specify only one.")
			return
		}

		if projectPath != "" {
			logrus.Infof("Generating tests for project at: %s", projectPath)

			if !file.Exists(projectPath) {
				logrus.Errorf("The project path %s doesn't exists", projectPath)
				return
			}

			if !file.IsDir(projectPath) {
				logrus.Errorf("The project path %s is not a directory. Please provide a directory", projectPath)
				return
			}

			testGenerator = generator.NewProjectTestGenerator(projectPath)
		} else {
			// generate tests for specific file
			if !file.Exists(filePath) {
				logrus.Errorf("The file path %s doesn't exists", filePath)
				return
			}

			if file.IsDir(filePath) {
				logrus.Errorf("The file path %s must not be a directory. Please provide a file", filePath)
				return
			}

			testGenerator = generator.NewFileTestGenerator(filePath)
		}

		testGenerator.GenerateTests()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Définir le flag --project pour la commande generate
	generateCmd.Flags().StringVarP(&projectPath, "project", "p", "", "Path to the project")
	generateCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file")
}
