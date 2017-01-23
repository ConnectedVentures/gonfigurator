package gonfigurator

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

// Loader provides a mockable interface for the gonfigurator library
type Loader interface {
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
