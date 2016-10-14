# Gonfigurator
This package is used to load YAML configuration files and parse them into
structs.

# Install

## go

```sh
go get github.com/connectedventures/gonfigurator
```

## gb

```sh
gb vendor fetch github.com/connectedventures/gonfigurator
```

# Example

```go
import (
	"github.com/connectedventures/gonfigurator"
)

type cloudConfig struct {
	ProjectID string       `json:"project_id"`
}

type loggingConfig struct {
	Level string `json:"level"`
}

type myserviceConfig struct {
	Cloud   cloudConfig   `json:"cloud"`
	Logging loggingConfig `json:"logging"`
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
