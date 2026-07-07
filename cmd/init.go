package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/ttsx/interactive-dockerizer/internal/generator"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Docker files interactively",
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() {
	config := generator.Config{}

	// 1. Language
	promptLang := &survey.Select{
		Message: "What is the primary language of your project?",
		Options: []string{"Go", "Node.js", "Python", "Other (Custom)"},
	}
	survey.AskOne(promptLang, &config.Language)

	if config.Language == "Other (Custom)" {
		promptCustomImg := &survey.Input{
			Message: "Enter your desired base Docker image (e.g., ubuntu:22.04):",
		}
		survey.AskOne(promptCustomImg, &config.CustomImage)
	}

	// 2. Port
	promptPort := &survey.Input{
		Message: "Which port does your application listen on?",
		Default: "8080",
	}
	survey.AskOne(promptPort, &config.Port)

	// 3. Database
	promptDB := &survey.Select{
		Message: "Do you need a database?",
		Options: []string{"None", "MySQL", "PostgreSQL", "MongoDB"},
	}
	survey.AskOne(promptDB, &config.Database)

	// 4. Cache
	promptCache := &survey.Select{
		Message: "Do you need caching?",
		Options: []string{"None", "Redis"},
	}
	survey.AskOne(promptCache, &config.Cache)

	fmt.Println("\nGenerating Docker files...")
	err := generator.GenerateFiles(config)
	if err != nil {
		fmt.Printf("Error generating files: %v\n", err)
		return
	}
	fmt.Println("Done! Run `docker compose up -d` to start your environment.")
}