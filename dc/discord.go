package dc

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	BotToken  string
	responses map[string]UserData = map[string]UserData{}
)

type UserData struct {
	OriginChannelId string
	Clock_In        int
	Clock_Out       int
}

type QuoteData struct {
	Quotes  []string `json:"quotes"`
	Authors []string `json:"author"`
}

const prefix string = "!bot"

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

func QuotesSend() []string {
	// Read the JSON file
	fileContent, err := os.ReadFile("/home/alan/src/golang-api-db/Sheet-Linker/dc/quotes.json")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return nil
	}

	// Create an instance of QuoteData
	var data QuoteData

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return data.Quotes
	}
	return data.Quotes
}

func EmbedMessage() {
	author := discordgo.MessageEmbedAuthor{
		Name: "",
	}
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
	}

	args := strings.Split(message.Content, " ")

	if args[0] != prefix {
		return
	}
	// Access the quotes as a slice of strings
	quotes := QuotesSend()

	// Selects a random quote from the slice of strings of quotes
	selection := rand.Intn(len(quotes))

	// The bot prints a random quote using the !bot prefix
	if args[1] == "quotes" {
		discord.ChannelMessageSend(message.ChannelID, quotes[selection])
	}

}

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
