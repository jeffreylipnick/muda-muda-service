package config

import "os"

type conf struct {
	Port string 
}

func NewConfig() conf {
	config := conf {
		Port: os.Getenv("MUDA_PORT"),
	}

	return config 
}
