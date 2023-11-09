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
	// Respond to User messages using switch statementso we can answer a set of messages
	switch {
	case strings.Contains(message.Content, "time"):
		discord.ChannelMessageSend(message.ChannelID, "I can provide that information")
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!/ Hola!")
	case strings.Contains(message.Content, "narti"):
		discord.ChannelMessageSend(message.ChannelID, "NartiBot es una basura")
	}

}

// func RunBot(Token string) {
// 	dg, err := discordgo.New("Bot" + Token)
// 	if err != nil {
// 		fmt.Print("Error creating discord session: ", err)
// 		return
// 	}

// 	dg.AddHandler(MessageCreate)

// 	dg.Identify.Intents = discordgo.IntentGuildMessages

// 	err = dg.Open()
// 	if err != nil {
// 		fmt.Print("Error opening connection", err)
// 		return
// 	}

// }

// func TerminationSignal(dg *discordgo.Session) {
// 	fmt.Print("Bot is now running. Press CTRL + C to exit")
// 	sc := make(chan os.Signal, 1)
// 	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
// 	<-sc

// 	dg.Close()
// }

// func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	if m.Content == "ping" && m.Author.ID != s.State.User.ID {
// 		channel, err := s.UserChannelCreate(m.Author.ID)
// 		if err != nil {
// 			fmt.Println("Error creating channel:", err)
// 			s.ChannelMessageSend(m.ChannelID, "Something went wrong while sending the DM!")
// 			return
// 		}

// 		_, err = s.ChannelMessageSend(channel.ID, "Pong!")
// 		if err != nil {
// 			fmt.Println("Error sending DM message:", err)
// 			s.ChannelMessageSend(m.ChannelID, "Failed to send you a DM. Did you disable DM in your privacy settings?")
// 		}
// 	}
// }
