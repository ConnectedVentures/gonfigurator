# Gonfigurator
This package is used to load YAML configuration files and parse them into
structs.

# Install

## go

```sh
go get github.com/fresh8gaming/gonfigurator
```

## gb

```sh
gb vendor fetch github.com/fresh8gaming/gonfigurator
```

# Example

```go
import (
	"github.com/fresh8gaming/gonfigurator"
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
