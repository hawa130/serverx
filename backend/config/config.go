package config

import (
	"log"
	"net/http"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	config         GlobalConfig
	isLoaded       = false
	onConfigChange func()
)

func OnConfigChange(run func()) {
	onConfigChange = run
}

func Config() *GlobalConfig {
	if !isLoaded {
		initViper()
		load()
	}
	return &config
}

func load() {
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file: " + err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Error unmarshalling config: " + err.Error())
	}

	isLoaded = true
}

func initViper() {
	initDefault()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigType("toml")
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config changed:", e.Name)
		load()
		onConfigChange()
	})
	viper.WatchConfig()
}

func initDefault() {
	viper.SetDefault("server.address", ":8080")

	viper.SetDefault("graphql.playground", true)
	viper.SetDefault("graphql.playground_endpoint", "/playground")
	viper.SetDefault("graphql.introspection", true)
	viper.SetDefault("graphql.endpoint", "/graphql")

	viper.SetDefault("jwt.duration", 24)
	viper.SetDefault("jwt.renew_duration", 12)

	viper.SetDefault("argon2.memory", 64*1024)
	viper.SetDefault("argon2.iterations", 3)
	viper.SetDefault("argon2.parallelism", 2)
	viper.SetDefault("argon2.salt_length", 16)
	viper.SetDefault("argon2.key_length", 32)

	viper.SetDefault("server.cors.allowed_origins", []string{"*"})
	viper.SetDefault("server.cors.allowed_methods", []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete})
	viper.SetDefault("server.cors.allowed_headers", []string{})
	viper.SetDefault("server.cors.exposed_headers", []string{})
	viper.SetDefault("server.cors.allow_credentials", false)
	viper.SetDefault("server.cors.max_age", 0)
}
