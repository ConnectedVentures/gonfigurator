package gonfigurator

import (
	"errors"
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

func TestLoaderParse(t *testing.T) {
	var cfg myserviceConfig
	l := &Loader{
		ConfigPath: "./example.yaml",
	}

	err := l.Parse("./example.yaml", &cfg)
	if err != nil {
		t.Error(err)
	}

	if cfg.Logging.Level != "debug" {
		t.Error("expected logging.level \"debug\", instead found: " + cfg.Logging.Level)
	}
}

func TestMockLoader(t *testing.T) {
	mockedLoader := &LoaderBlueprintMock{
		ParseFunc: func(path string, target interface{}) error {
			return errors.New("something went wrong")
		},
	}
	var someTarget interface{}
	err := mockedLoader.Parse("some path", someTarget)
	if err == nil {
		t.Error("expected error, found nil")
	}

	mockedLoader.ParseFunc = func(path string, target interface{}) error {
		return nil
	}

	err = mockedLoader.Parse("some path", someTarget)
	if err != nil {
		t.Error(err)
	}
}
