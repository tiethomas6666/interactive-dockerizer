package generator

import (
	"embed"
	"os"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

type Config struct {
	Language    string
	CustomImage string
	Port        string
	Database    string
	Cache       string
}

func GenerateFiles(cfg Config) error {
	if err := renderTemplate("Dockerfile", "templates/Dockerfile.tmpl", cfg); err != nil {
		return err
	}
	if err := renderTemplate("docker-compose.yml", "templates/docker-compose.yml.tmpl", cfg); err != nil {
		return err
	}
	return nil
}

func renderTemplate(outputFilename, tmplPath string, cfg Config) error {
	tmplData, err := templateFS.ReadFile(tmplPath)
	if err != nil {
		return err
	}

	t, err := template.New(outputFilename).Parse(string(tmplData))
	if err != nil {
		return err
	}

	f, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, cfg)
}