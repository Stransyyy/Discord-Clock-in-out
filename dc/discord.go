package dc

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken string
)

func Run() {
	// Creates a new discord session
	discord, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		log.Fatal("Error creating discord session", err)
	}

	// Add event handler for general messages
	discord.AddHandler(newMessage)

	// Open's and closes with the defer the session

	err = discord.Open()
	if err != nil {
		log.Fatal("Could not open session: ", err)
	}

	defer discord.Close()

	fmt.Println("Bot running...")

	// Exit the session/program CTRL + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignores bot own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}
	// Respond to User messages using switch statementso we can answer a set of  predetermined messages
	switch {
	case strings.Contains(message.Content, "time"):
		discord.ChannelMessageSend(message.ChannelID, "I can provide that information")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!/ Hola!")
	case strings.Contains(message.Content, "narti"):
		discord.ChannelMessageSend(message.ChannelID, "NartiBot es una basura")
	case strings.Contains(message.Content, "ahuevo"):
		discord.ChannelMessageSend(message.ChannelID, "ahuevo que si como no? ")
	case strings.Contains(message.Content, "Stransyyy"):
		discord.ChannelMessageSend(message.ChannelID, "Es el mejor we")
	}

}

// func MandaDatos() *discordgo.ApplicationCom {

// }

func askQuestionApplicationCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "clock-in",
		Description: "It is time to clock-in at work!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "request",
				Description: "Send your clock-in information to the database",
				Required:    true,
			},
		},
	}
}
