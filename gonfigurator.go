package gonfigurator

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type config struct {
	flag   *string
	target interface{}
}

var (
	configs []config
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

// Parse loads the YAML file at the given path and reads it into v
func Parse(defaultPath string, v interface{}) error {
	ParseCustomFlag(defaultPath, "c", v)
	return Load()
}

// ParseCustomFlag will schedule to load the YAML file from the default path or using the
// path specified using custom command line flag (flagName) when Load() is called
func ParseCustomFlag(defaultPath string, flagName string, v interface{}) {
	f := flag.String(flagName, defaultPath, "Path to the configuration file")
	configs = append(configs, config{flag: f, target: v})
}

func Load() error {
	flag.Parse()

	for _, conf := range configs {
		f := conf.flag
		contents, err := ioutil.ReadFile(*f)
		if err != nil {
			return fmt.Errorf("Could not read config file: %s", err.Error())
		}
		err = yaml.Unmarshal([]byte(contents), conf.target)
		if err != nil {
			return fmt.Errorf("Could not parse config file: %s", err.Error())
		}
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
