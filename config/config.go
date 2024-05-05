package config

import (
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string
)

func LoadConfig() error {
	fmt.Println("Loading env values")
	
	Token = os.Getenv("TOKEN")
	BotPrefix = os.Getenv("BOT_PREFIX")

	fmt.Println("Successfully loaded env values")

	return nil
}