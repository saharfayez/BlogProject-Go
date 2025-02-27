package impl

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"goproject/internal/app/context"
	testdatabase "goproject/test/database"
	"log"
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

	mainConfig = testdatabase.Config(mainConfig)

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
