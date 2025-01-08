package main

import "fmt"

type Config struct {
	token string
}

func LoadConfig() Config {
	return LoadConfigFromFile("./config")
}

func LoadConfigFromFile(filepath string) Config {
	fmt.Println("[CONFIG] Load config file from", filepath);
	return Config{}
}
