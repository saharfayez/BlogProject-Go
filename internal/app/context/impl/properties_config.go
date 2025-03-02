package impl

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"goproject/internal/app/context"
	"goproject/test/config"
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
		log.Fatalf("no mainConfig found at %s: %v", mainConfigPath, err)
	}

	profile := mainConfig.String("app.profile")

	fmt.Println("profile after main", profile)

	if testing.Testing() {
		mainConfig = config.LoadProperties(mainConfig)
	}

	profile = mainConfig.String("app.profile")
	databaseURL := mainConfig.String("database.url")

	fmt.Println("profile after merge", profile)
	fmt.Println("databaseURL after merge", databaseURL)

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
