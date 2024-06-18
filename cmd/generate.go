package cmd

import (
	gcore "github.com/dterbah/go-test-gen/core"
	file "github.com/dterbah/go-test-gen/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var projectPath string

// generateCmd représente la commande generate
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate tests for the specified project",
	Long: `Generate unit tests for the specified Go project.
You need to provide the path to the project using the --project flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		var testGenerator gcore.TestGenerator
		if projectPath == "" {
			logrus.Error("Please provide the path to the project using the --project flag.")
		} else {
			logrus.Infof("Generating tests for project at: %s", projectPath)

			if exists, _ := file.Exists(projectPath); !exists {
				logrus.Errorf("The project path %s doesn't exists", projectPath)
				return
			}

			if !file.IsDir(projectPath) {
				logrus.Errorf("The project path %s is not a directory. Please provide a directory", projectPath)
				return
			}

			testGenerator = gcore.NewTestGenerator(projectPath)
			testGenerator.GenerateTests()
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Définir le flag --project pour la commande generate
	generateCmd.Flags().StringVarP(&projectPath, "project", "p", "", "Path to the project")
}
