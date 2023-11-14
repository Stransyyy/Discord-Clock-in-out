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

	// Discord server ID
	StransyyyBotChanneId string
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

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignores bot own messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to User messages using switch statementso we can answer a set of  predetermined messages
	switch {
	case strings.Contains(message.Content, "time"):
		discord.ChannelMessageSend(message.ChannelID, "I can provide that information")
	case strings.Contains(message.Content, "hola"):
		discord.ChannelMessageSend(message.ChannelID, "Hola Jersey")
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

	commandHandler := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"clock-in": ClockIn,
	}

	commands := []*discordgo.ApplicationCommand{
		clockinTimeCommand(),
	}

	for _, c := range commands {
		_, cmnderr := discord.ApplicationCommandCreate(os.Getenv(StransyyyBotChanneId), "", c)

		if cmnderr != nil {
			panic(cmnderr)
		}
	}

	discord.AddHandler(func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if h, ok := commandHandler[interaction.ApplicationCommandData().Name]; ok {
			h(session, interaction)
		}
	})

}

func clockinTimeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "Clock-In",
		Description: "Run this command to clock in to work!",
	}

}

func ClockIn(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	session.ChannelMessageSend(StransyyyBotChanneId, "Clock-In? we still don't have that option yet, come back and use it soon.")
}
