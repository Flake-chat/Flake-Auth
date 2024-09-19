package store

type Config struct {
	DB string
}

func NewConfig() *Config {
	return &Config{}
}
