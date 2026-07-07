package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-dockerizer",
	Short: "An interactive CLI to generate Docker environments",
	Long:  `A fast and flexible CLI tool built with Go to generate Dockerfile and docker-compose.yml for various projects.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}