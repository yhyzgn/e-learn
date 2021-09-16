package config

import "e-learn/logger"

// Config ...
type (
	Config struct {
		MySQL struct {
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
			Params   string `yaml:"params"`
		}

		Logging logger.Logging `yaml:"logging"`
	}
)
