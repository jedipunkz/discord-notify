package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type UserState struct {
	Name      string
	CurrentVC string
}

var (
	Token   = "Bot " + ""
	BotName = ""
	stopBot = make(chan bool)
	discord *discordgo.Session
	usermap = map[string]*UserState{}
)

func main() {
	var err error
	discord, err = discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discord.AddHandler(onVoiceStateUpdate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-stopBot
	return
}

func onVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {

	_, ok := usermap[vs.UserID]
	if !ok {
		usermap[vs.UserID] = new(UserState)
		user, _ := discord.User(vs.UserID)
		usermap[vs.UserID].Name = user.Username
		log.Print("new user added : " + user.Username)
	}

	if len(vs.ChannelID) > 0 && usermap[vs.UserID].CurrentVC != vs.ChannelID {
		channel, _ := discord.Channel(vs.ChannelID)
		message := usermap[vs.UserID].Name + "さんが" + channel.Name + "にジョインしました"
		log.Print(message)
		fmt.Println(message)

		s.ChannelMessageSend("", message)
	}

	usermap[vs.UserID].CurrentVC = vs.ChannelID

	fmt.Printf("%+v", vs.VoiceState)
}
