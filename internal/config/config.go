package config

import (
	"log"
	"os"
	"sync"
)

var Config AppConfig
var readCfgOnce sync.Once

type AppConfig struct {
	Port           string `json:"-"`
	OpenAI_API_Key string `json:"-"`
}

func Init() {
	readCfgOnce.Do(func() {
		log.Println("Setting up config...")
		Config = AppConfig{
			Port:           os.Getenv("PORT"),
			OpenAI_API_Key: os.Getenv("OPENAI_API_KEY"),
		}
	})
}
