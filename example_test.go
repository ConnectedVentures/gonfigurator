package main

import (
	"log"
	"testing"
)

type cloudConfig struct {
	ProjectID string `json:"project_id"`
}

type loggingConfig struct {
	Level string `json:"level"`
}

type myserviceConfig struct {
	Cloud   cloudConfig   `json:"cloud"`
	Logging loggingConfig `json:"logging"`
}

var config myserviceConfig

func TestTest(t *testing.T) {
	err := Parse("./example.yaml", &config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config.Cloud.ProjectID)
}
