package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version     string
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "go-test-gen",
	Short: "A CLI tool to generate Go tests",
	Long: `go-test-gen is a CLI tool to generate Go tests for your project.
You can specify the path to your project and it will generate tests for all Go files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf("Version : v%s\n", version)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// init logger
	// Configuration de Logrus
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		ForceColors:            true,
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version")
	version = "1.0.3"
}
