package main

import (
	"fmt"

	"github.com/ShuHado/GoBot_discord.git/bot"
	"github.com/ShuHado/GoBot_discord.git/config"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<- make(chan struct{})
}