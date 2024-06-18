package cmd

import (
	"fmt"

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
		if projectPath == "" {
			fmt.Println("Please provide the path to the project using the --project flag.")
		} else {
			fmt.Printf("Generating tests for project at: %s\n", projectPath)
			// Ajoutez ici la logique pour générer les tests
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Définir le flag --project pour la commande generate
	generateCmd.Flags().StringVarP(&projectPath, "project", "p", "", "Path to the project")
}
