package main

import (
	"log"
	"os"
	"sync"
)

var Config AppConfig
var readCfgOnce sync.Once

type AppConfig struct {
	Port      string `json:"-"`
	OpenAIKey string `json:"-"`
}

func init() {
	readCfgOnce.Do(func() {
		log.Println("Setting up config...")
		Config = AppConfig{
			Port: os.Getenv("PORT"),
		}
	})
}
