package main

import (
	"fmt"
	"log"
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

	discord, err = discordgo.New()
	discord.Token = token
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
	channel_id := viper.GetString("channel_id")

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

		s.ChannelMessageSend(channel_id, message)
	}

	usermap[vs.UserID].CurrentVC = vs.ChannelID

	fmt.Printf("%+v", vs.VoiceState)
}
