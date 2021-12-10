package apiserver

import "http-rest-api/iternal/app/store"

type Config struct {
	BindAdr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

//NewConfig...
func NewConfig() *Config {
	return &Config{
		BindAdr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
