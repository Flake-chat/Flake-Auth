package api

type Config struct {
	DB        string
	Token     string
	Kafka_url string
	LogLevel  string
}

func NewConfig() *Config {
	return &Config{}
}
