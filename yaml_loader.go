package gonfigurator

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Parse loads the .yml file from Loader.ConfigPath if set, defaultPath as a fallback.
// Adheres to the LoaderBlueprint interface.
func (l *YamlLoader) Parse(path string, target interface{}) error {
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Could not read config file: %s", err.Error())
	}
	err = yaml.Unmarshal([]byte(contents), target)
	if err != nil {
		return fmt.Errorf("Could not parse config file: %s", err.Error())
	}
	return nil
}

// Path returns the config path for the yaml file
// returns ConfigPath if specified on the struct, else attempts to parse from -c flag
func (l *YamlLoader) Path() string {
	cfgPath := l.ConfigPath
	if l.ConfigPath == "" {
		f := flag.String("c", "", "Path to the configuration file")
		flag.Parse()
		cfgPath = *f
	}

	return cfgPath
}
