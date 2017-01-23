package gonfigurator

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// YamlLoader implements the loader interface to load YAML files
type YamlLoader struct {
	ConfigPath string
}

// ConfigLoader should load a file at the given path into the target interface
type ConfigLoader interface {
	Parse(path string, target interface{}) error
	Path() string
}

// Parse loads the .yml file at the given path and reads it into v
func Parse(defaultPath string, v interface{}) error {
	f := flag.String("c", defaultPath, "Path to the configuration file")
	flag.Parse()

	contents, err := ioutil.ReadFile(*f)
	if err != nil {
		return fmt.Errorf("Could not read config file: %s", err.Error())
	}
	err = yaml.Unmarshal([]byte(contents), v)
	if err != nil {
		return fmt.Errorf("Could not parse config file: %s", err.Error())
	}
	return nil
}

// ParseConfig uses the specified loader to extract configuration into the target
func ParseConfig(loader ConfigLoader, target interface{}) error {
	path := loader.Path()
	if path == "" {
		return errors.New("no path could be loaded")
	}

	return loader.Parse(path, target)
}

