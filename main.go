package main

import (
	"fmt"

	"go-discord-music-bot/bot"
	"go-discord-music-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
