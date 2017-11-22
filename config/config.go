package config

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	Mode string `json:"mode"`
}

var conf *Config

const (
	DebugMode      = "debug"
	ProductionMode = "production"
)

func init() {
	// TODO infer mode from env var?
	conf = &Config{
		Mode: DebugMode,
	}
}

// Cfg returns the configuration - will panic if the config has not been loaded or is nil (which shouldn't happen as that's implicit in the package init)
func Cfg() *Config {
	return conf
}

func (cfg *Config) GetLogger() *logrus.Logger {
	var l = logrus.New()
	// l.Formatter = &logrus.JSONFormatter{}
	return l
}

func (cfg *Config) IsDebugMode() bool {
	return cfg.Mode == DebugMode
}

func (cfg *Config) IsProductionMode() bool {
	return cfg.Mode == ProductionMode
}
