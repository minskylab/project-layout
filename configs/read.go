package configs

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/pkg/errors"
)

func ConfigFromFile(filepath string) (*Config, error) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)
	config.WithOptions(func(opt *config.Options) {
		opt.TagName = "config"
	})

	if err := config.LoadFiles(filepath); err != nil {
		return nil, errors.WithStack(err)
	}

	c := new(Config)

	return c, config.BindStruct("", c)
}

func MustConfigFromFile(filepath string) *Config {
	c, err := ConfigFromFile(filepath)
	if err != nil {
		panic(err)
	}

	return c
}
