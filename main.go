package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type UserState struct {
	Name      string
	CurrentVC string
}

var (
	stopBot = make(chan bool)
	discord *discordgo.Session
	usermap = map[string]*UserState{}
)

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".discord-notify")

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func main() {
	var err error
	token := "Bot " + viper.GetString("token")

	discord, err = discordgo.New(token)
	if err != nil {
		fmt.Printf("Login Error: %s", err)
	}

	discord.AddHandler(onVoiceStateUpdate)

	err = discord.Open()
	if err != nil {
		fmt.Printf("Session Error: %s", err)
	}

	fmt.Println("Listening...")
	<-stopBot
	return
}

func onVoiceStateUpdate(s *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	notify_channel_id := viper.GetString("notify_channel_id")
	me := viper.GetString("me")

	_, ok := usermap[vs.UserID]
	if !ok {
		usermap[vs.UserID] = new(UserState)
		user, _ := discord.User(vs.UserID)
		usermap[vs.UserID].Name = user.Username
	}

	if len(vs.ChannelID) > 0 && usermap[vs.UserID].CurrentVC != vs.ChannelID && usermap[vs.UserID].Name != me {
		channel, _ := discord.Channel(vs.ChannelID)
		message := usermap[vs.UserID].Name + " Joined to " + channel.Name + " Channel"

		s.ChannelMessageSend(notify_channel_id, message)
	}

	usermap[vs.UserID].CurrentVC = vs.ChannelID
}
