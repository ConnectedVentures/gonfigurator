package gonfigurator

import "testing"

func TestYamlLoaderParse(t *testing.T) {
	var cfg myserviceConfig
	l := &YamlLoader{
		ConfigPath: "./no.yaml",
	}

	err := ParseConfig(l, &cfg)
	if err == nil {
		t.Error("expected no file found error")
	}

	l.ConfigPath = "./example.yaml"
	err = ParseConfig(l, &cfg)
	if err != nil {
		t.Error(err)
	}

	if cfg.Logging.Level != "debug" {
		t.Error("expected logging.level \"debug\", instead found: " + cfg.Logging.Level)
	}
}
