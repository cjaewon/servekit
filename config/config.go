package config

import (
	"log"

	"github.com/spf13/viper"
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

	v.SetConfigType("toml")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME")
	v.SetConfigName(".servekit")

	v.SetDefault("server", map[string]interface{}{
		"port":     ":3000",
		"path":     "./static",
		"mode":     "none",
		"overview": false,
	})

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Using default configuration")
	} else {
		log.Printf("Using provided configuration")
	}

	var c Config
	v.Unmarshal(&c)

	return &c
}
