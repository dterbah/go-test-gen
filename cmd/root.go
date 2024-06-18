package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-test-gen",
	Short: "A CLI tool to generate Go tests",
	Long: `go-test-gen is a CLI tool to generate Go tests for your project.
You can specify the path to your project and it will generate tests for all Go files.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
