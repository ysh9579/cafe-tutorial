package api

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"hello-cafe/internal/db"
)

const defaultConfigPath = "/root/.hello-cafe/config.yml"

type Configure struct {
	DB db.Config `yaml:"db"`
}

func unmarshalConfig(path string, cfg *Configure) error {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "read config file")
	}
	if err := yaml.Unmarshal(configBytes, cfg); err != nil {
		return errors.Wrap(err, "unmarshal config")
	}

	return nil
}

func Config() (*Configure, error) {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = defaultConfigPath
	}

	var cfg Configure
	if err := unmarshalConfig(path, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal config [ path = %s ]", path)
	}

	return &cfg, nil
}
