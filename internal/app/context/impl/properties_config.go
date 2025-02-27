package impl

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"goproject/internal/app/context"
	"log"
)

type propertiesConfig struct {
	profile     string
	databaseUrl string
}

func newPropertiesConfig() context.PropertiesConfig {
	k := koanf.New(".")

	err := k.Load(file.Provider("../../cmd/goproject/config.yaml"), yaml.Parser())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("profile after load main", k.String("app.profile"))
	fmt.Println("databaseUrl after load main", k.String("database.url"))

	override := koanf.New(".")
	overridePath := "config.yaml"

	err = override.Load(file.Provider(overridePath), yaml.Parser())
	if err != nil {
		log.Printf("no override config found at %s: %v", overridePath, err)
	}

	fmt.Println("profile after load override", override.String("app.profile"))
	fmt.Println("databaseUrl after load override", override.String("database.url"))

	if err = k.Merge(override); err != nil {
		log.Fatalf("error merging override config: %v", err)
	}

	profile := k.String("app.profile")
	databaseUrl := k.String("database.url")

	fmt.Println("profile after merge", profile)
	fmt.Println("databaseUrl after merge", databaseUrl)

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
