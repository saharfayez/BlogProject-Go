package impl

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"goproject/internal/app/context"
	"log"
)

var k = koanf.New(".")

type propertiesConfig struct {
	profile     string
	databaseUrl string
}

func newPropertiesConfig() context.PropertiesConfig {

	err := k.Load(file.Provider("config.yaml"), yaml.Parser())
	if err != nil {
		log.Fatal(err)
	}

	profile := k.String("profile")
	databaseUrl := k.String("database_url")

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
