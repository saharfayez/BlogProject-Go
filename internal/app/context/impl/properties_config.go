package impl

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"goproject/internal/app/context"
	"log"
	"testing"
)

type propertiesConfig struct {
	Profile     string `koanf:"app.profile"`
	DatabaseUrl string `koanf:"database.url"`
}

func newPropertiesConfig() context.PropertiesConfig {
	mainConfig := koanf.New(".")
	mainConfigPath := "../../cmd/goproject/config.yaml"

	err := mainConfig.Load(file.Provider(mainConfigPath), yaml.Parser())
	if err != nil {
		log.Fatal(err)
	}

	if testing.Testing() {
		fmt.Println("we run tests")

		testConfig := koanf.New(".")
		testConfigPath := "config.yaml"

		err = testConfig.Load(file.Provider(testConfigPath), yaml.Parser())
		if err != nil {
			log.Printf("no testConfig config found at %s: %v", testConfigPath, err)
		}

		if err = mainConfig.Merge(testConfig); err != nil {
			log.Fatalf("error merging testConfig config: %v", err)
		}
	}

	var propConfig propertiesConfig

	err = mainConfig.UnmarshalWithConf("", &propConfig, koanf.UnmarshalConf{Tag: "koanf", FlatPaths: true})
	if err != nil {
		log.Fatal(err)
	}

	return &propConfig
}

func (config *propertiesConfig) GetProfile() string {
	return config.Profile
}

func (config *propertiesConfig) GetDatabaseUrl() string {
	return config.DatabaseUrl
}
