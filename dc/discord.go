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
	responses map[string]UserDM = map[string]UserDM{}

	// Discord server ID
	StransyyyBotChanneId string
)

// Will handle the DM and answers back from the user
type UserDM struct {
	TotalTime  int
	WantToKnow string
}

// JSON Quote data
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
		"clockin":  ClockInResponse,
		"clockout": clockOutResponse,
	}

	discord.Identify.Intents = discordgo.IntentGuildMessages

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	var commands []*discordgo.ApplicationCommand

	commands = append(commands, clockinTimeCommand(), clockoutTimeCommand())

	for _, c := range commands {
		_, cmderr := discord.ApplicationCommandCreate(os.Getenv("BOT_APP_ID"), "", c)

		if cmderr != nil {
			log.Fatal("This is the commands error at line 76", cmderr)
		}

		_, cmderr = discord.ApplicationCommandCreate(os.Getenv("BOT_APP_ID"), "", c)
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

// newMessage sends a new message. Does not reply to slash commands
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

	// let the user use !bot and the key word just for the bot to reply to that specific input
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

// clockInTimeCommand on discord is the part where you use the slash command and shows a preview with the name of the command and the description of it
func clockinTimeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "clockin",
		Description: "Run this command to clock in to work!",
	}

}

// clockoutTimeCommand is the slash command that displays on discord and shows the description of it
func clockoutTimeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "clockout",
		Description: "Run this command to clockout from work, and send the data to the database",
	}

}

// clockInEmbed will send an embed
func clockInEmbed() *discordgo.MessageEmbed {

	image := discordgo.MessageEmbedImage{
		URL: "https://img.craiyon.com/2023-11-16/884s_1eZTiepm3y9B6d7nA.webp",
	}

	embed := discordgo.MessageEmbed{
		Title:       "Clock-In",
		Description: "Use this command to let you clock-in and send your data to the database",
		Timestamp:   "",
		Image:       &image,
	}

	return &embed
}

// clockOutEmbed is the set up for the clock-out embed
func clockOutEmbed() []*discordgo.MessageEmbed {

	image := discordgo.MessageEmbedImage{
		URL: "https://img.craiyon.com/2023-11-20/lwkWz-yhSRKqMl38plwCqw.webp",
	}

	embed := []*discordgo.MessageEmbed{}

	embed = append(embed, &discordgo.MessageEmbed{
		URL:         "https://vitalitysouth.com/",
		Title:       "Clock-Out",
		Description: "Run this command to clockout from work, and send the data to the database",
		Color:       10038562,
		Image:       &image,
	})

	return embed
}

// ClockInResponse sends the response of the bot when you use the slash command
func ClockInResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channelID := "1172648319940558970"
	_, ok := session.ChannelMessageSend(channelID, "Stransyyy bot esta siendo usado...")
	if ok != nil {
		panic(ok)
	}

	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})
	if err != nil {
		return
	}

	// _, serr := session.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
	// 	Content: "hello",
	// })

	// if serr != nil {
	// 	panic(serr)
	// }

	aquaEmbed := 1752220
	msgEmbed := []*discordgo.MessageEmbed{}
	msgEmbed = append(msgEmbed, &discordgo.MessageEmbed{
		Title: "ClockIn",
		URL:   "https://vitalitysouth.com/",
		Image: clockInEmbed().Image,
		Color: aquaEmbed,
	})

	clockinmsgData := &discordgo.WebhookParams{
		Embeds:     msgEmbed,
		Components: []discordgo.MessageComponent{},
	}

	_, smerr := session.FollowupMessageCreate(interaction.Interaction, false, clockinmsgData)
	if smerr != nil {
		log.Fatal(fmt.Sprintf("followup message create with embeds: %v", smerr))
		session.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Clockin",
					Description: "Discord isn't letting me send my full response. I would much rather have gone with another product than stay here with Discord. I don't know what all this trouble is about, but I'm sure it must be Discord's fault.",
					Color:       aquaEmbed,
				},
			},
		})
		return
	}
}

// clockOutResponse will send an embed as a response to the slash command call
func clockOutResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channelID := "1172648319940558970"
	_, ok := session.ChannelMessageSend(channelID, "Stransyyy bot esta siendo usado...")
	if ok != nil {
		panic(ok)
	}

	clockOutMsgData := &discordgo.WebhookParams{
		Embeds:     clockOutEmbed(),
		Components: []discordgo.MessageComponent{},
	}

	_, smerr := session.FollowupMessageCreate(interaction.Interaction, false, clockOutMsgData)
	if smerr != nil {
		log.Fatal(fmt.Sprintf("Follow up message create with embeds: %v", smerr))
		session.FollowupMessageCreate(interaction.Interaction, false, &discordgo.WebhookParams{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Clockout",
					Description: "Discord isn't letting me send my full response. I would much rather have gone with another product than stay here with Discord. I don't know what all this trouble is about, but I'm sure it must be Discord's fault.",
					Color:       10038562,
				},
			},
		})
	}
}
