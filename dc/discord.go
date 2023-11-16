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
	Quotes []struct {
		Quote  string `json:"quote"`
		Author string `json:"author"`
	} `json:"quotes"`
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

	discord.SyncEvents = false

	commandHandlers := map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"clockin": ClockInResponse,
	}

	discord.Identify.Intents = discordgo.IntentGuildMessages

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	var commands []*discordgo.ApplicationCommand

	commands = append(commands, clockinTimeCommand())

	for _, c := range commands {
		_, cmderr := discord.ApplicationCommandCreate(os.Getenv("BOT_APP_ID"), "", c)

		if cmderr != nil {
			log.Fatal("This is the commands error at line 76", err)
		}
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

		return []string{}
	}

	var quotes []string

	for _, q := range data.Quotes {
		quotes = append(quotes, fmt.Sprintf("%s - %s", q.Quote, q.Author))
	}

	return quotes
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

	discord.ChannelMessageSendEmbed("prueba", clockEmbed())

}

func clockinTimeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "clockin",
		Description: "Run this command to clock in to work!",
	}

}

func clockEmbed() *discordgo.MessageEmbed {

	image := discordgo.MessageEmbedImage{
		URL: "https://img.craiyon.com/2023-11-16/884s_1eZTiepm3y9B6d7nA.webp",
	}

	embed := discordgo.MessageEmbed{
		Title:       "Clock-In",
		Description: "Use this command to let you clock-inh and send your data to the database",
		Timestamp:   "",
		Image:       &image,
	}

	return &embed
}

func ClockInResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	_, ok := session.ChannelMessageSend("1172648319940558970", "Stransyyy bot esta siendo usado...")

	if ok != nil {
		panic(ok)
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	if err != nil {
		return
	}

	_, serr := session.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
		Content: "hello",
	})

	if serr != nil {
		panic(serr)
	}

}
