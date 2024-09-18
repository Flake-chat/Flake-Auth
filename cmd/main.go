package main

import (
	"fmt"
	"os"
)

type config struct {
	DB        string
	Token     string
	Kafka_url string
}

func main() {
	c := config{}
	c.DB = os.Getenv("ZSH")
	c.Kafka_url = os.Getenv("KAFKA_URL")
	c.Token = os.Getenv("JWT_TOKEN")
	fmt.Println(c)
}
