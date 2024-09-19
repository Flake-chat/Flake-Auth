package main

import (
	"os"

	"github.com/Flake-chat/Flake-Auth/internal/api"
)

func main() {

	c := api.NewConfig()
	c.DB = os.Getenv("DATABASE_URL")
	c.Kafka_url = os.Getenv("KAFKA_URL")
	c.Token = os.Getenv("JWT_TOKEN")
	c.LogLevel = os.Getenv("LOGLEVEL")
	s := api.New(c)

	s.Start()
}
