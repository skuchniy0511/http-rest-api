package apiserver

type Config struct {
	BindAdr     string `toml:"bind_adr"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
}

//NewConfig...
func NewConfig() *Config {
	return &Config{
		BindAdr:  ":8080",
		LogLevel: "debug",
	}
}
