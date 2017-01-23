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

func TestMockLoader(t *testing.T) {
	mockedLoader := &ConfigLoaderMock{
		ParseFunc: func(path string, target interface{}) error {
			return errors.New("something went wrong")
		},
		PathFunc: func() string {
			return ""
		},
	}

	var someTarget interface{}
	err := ParseConfig(mockedLoader, someTarget)
	if err == nil {
		t.Error("expected no path error, found nil")
	}

	mockedLoader.PathFunc = func() string {
		return "a path"
	}

	err = mockedLoader.Parse("some path", someTarget)
	if err == nil {
		t.Error("expected parse error, found nil")
	}

	mockedLoader.ParseFunc = func(path string, target interface{}) error {
		return nil
	}

	err = mockedLoader.Parse("some path", someTarget)
	if err != nil {
		t.Error(err)
	}
}
