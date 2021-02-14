package config

import (
	"servekit/logger"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config is servekit's config struct type
type Config struct {
	Server struct {
		Port     string `mapstructure:"port"`
		Path     string `mapstructure:"path"`
		Mode     string `mapstructure:"mode"`
		Overview bool   `mapstructure:"overview"`
	} `mapstructure:"server"`
}

// LoadInConfig returns config or default config
func LoadInConfig() *Config {
	v := viper.New()

	v.SetConfigName(".servekit")
	v.SetConfigType("toml")
	v.AddConfigPath("$HOME")
	v.AddConfigPath(".")
	v.SetEnvPrefix("SERVEKIT")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.SetDefault("server", map[string]interface{}{
		"port":     ":3000",
		"path":     "./static",
		"mode":     "none",
		"overview": false,
	})

	for _, key := range v.AllKeys() {
		v.BindEnv(key)
	}

	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logger.Log().Info("Using default configuration")
		default:
			logger.Log().Fatal("Failed to load a config file", zap.Error(err))
		}
	} else {
		logger.Log().Info("Using a provided configuration", zap.String("config_file", v.ConfigFileUsed()))
	}

	var c Config
	v.Unmarshal(&c)

	return &c
}
