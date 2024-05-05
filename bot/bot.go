package bot

import (
	"fmt"

	"github.com/ShuHado/GoBot_discord.git/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot successfully connected")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == config.BotPrefix+"ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "cc" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "sv ?")
	}

	if m.Author.ID == "281451555058745344" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ta gueule")
	}
}