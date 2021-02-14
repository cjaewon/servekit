package config

import (
	"log"
	"strings"

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
			log.Printf("Using default configuration")
		default:
			log.Fatalf("Failed to load a config file %s", err)
		}
	} else {
		log.Printf("Using provided configuration")
	}

	var c Config
	v.Unmarshal(&c)

	return &c
}
