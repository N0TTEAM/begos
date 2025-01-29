package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"text/template"
)

const modelTmp = `package model
import "gorm.io/gorm"
type {{.ModelName}} struct {
gorm.Model
}
`

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Model name is required")
	}
	modelName := os.Args[1]

	if len(modelName) == 0 || !isUppercase(modelName[0]) {
		log.Fatalf("Invalid model name: %s. Model name must start with capital")
	}

	filepath := "internal/http/model/" + strings.ToLower(modelName) + ".go"
	if _, err := os.Stat(filepath); err == nil {
		log.Fatalf("file %s already exists", filepath)
	}

	tmpl, err := template.New("model").Parse(modelTmp)
	if err != nil {
		log.Fatalf("Failed to parse tmp: %v", err)
	}

	var rendered bytes.Buffer
	err = tmpl.Execute(&rendered, map[string]string{
		"ModelName": modelName,
	})
	if err != nil {
		log.Fatalf("failed to render tmp: %v", err)
	}

	err = os.WriteFile(filepath, rendered.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write model file: %v", err)
	}

	log.Printf("Model %s has been created successfully at %s", modelName, filepath)
}

func isUppercase(c byte) bool {
	return c >= 'A' && c <= 'Z'
}
