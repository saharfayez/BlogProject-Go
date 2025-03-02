package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

func LoadProperties(mainConfig *koanf.Koanf) *koanf.Koanf {
	testConfig := koanf.New(".")
	testConfigPath := "config.yaml"

	err := testConfig.Load(file.Provider(testConfigPath), yaml.Parser())
	if err != nil {
		log.Fatalf("no testConfig config found at %s: %v", testConfigPath, err)
	}

	fmt.Println("profile after test", testConfig.String("app.profile"))

	if err = mainConfig.Merge(testConfig); err != nil {
		log.Fatalf("error merging testConfig: %v", err)
	}
	return mainConfig
}
