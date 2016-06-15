# Gonfigurator
This package is used to load YAML configuration files and parse them into
structs.

# Install

## go

```sh
go get github.com/ConnectedVentures/gonfigurator
```

## gb

```sh
gb vendor fetch github.com/ConnectedVentures/gonfigurator
```

# Example

```go
import (
	"github.com/ConnectedVentures/gonfigurator"
)

type cloudConfig struct {
	ProjectID string       `yaml:"project_id"`
}

type loggingConfig struct {
	Level string `yaml:"level"`
}

type myserviceConfig struct {
	Cloud   cloudConfig   `yaml:"cloud"`
	Logging loggingConfig `yaml:"logging"`
}

var config myserviceConfig


func main() {
	err := gonfigurator.Parse("/etc/f8-uploader/uploader.yaml", &config)
	if err != nil {
		log.Fatal(err)
	}
  log.Println(config.Cloud.ProjectID)
}
```
