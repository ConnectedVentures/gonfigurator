package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

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
