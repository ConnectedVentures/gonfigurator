package gonfigurator

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Loader implements the library blueprint
type Loader struct {
	ConfigPath string
}

// LoaderBlueprint provides a mockable interface for the gonfigurator library
type LoaderBlueprint interface {
	Parse(string, interface{}) error
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

// Parse loads the .yml file from Loader.ConfigPath if set, defaultPath as a fallback.
// Adheres to the LoaderBlueprint interface.
func (l *Loader) Parse(defaultPath string, v interface{}) error {
	cfgPath := l.ConfigPath
	if l.ConfigPath == "" {
		f := flag.String("c", defaultPath, "Path to the configuration file")
		flag.Parse()
		cfgPath = *f
	}

	contents, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("Could not read config file: %s", err.Error())
	}
	err = yaml.Unmarshal([]byte(contents), v)
	if err != nil {
		return fmt.Errorf("Could not parse config file: %s", err.Error())
	}
	return nil
}
