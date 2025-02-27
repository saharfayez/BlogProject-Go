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
	profile     string
	databaseUrl string
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
	} else {
		fmt.Println("we are in normal mode")
	}

	profile := mainConfig.String("app.profile")
	databaseUrl := mainConfig.String("database.url")

	return &propertiesConfig{
		profile,
		databaseUrl,
	}
}

func (config *propertiesConfig) GetProfile() string {
	return config.profile
}

func (config *propertiesConfig) GetDatabaseUrl() string {
	return config.databaseUrl
}
