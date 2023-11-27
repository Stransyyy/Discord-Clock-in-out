package dc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
)

var (
	BotToken string
	//responses map[string]UserDM = map[string]UserDM{}

	// Discord server ID
	StransyyyBotChanneId string
)

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
		"clockout": ClockOutResponse,
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
		URL: "https://pics.craiyon.com/2023-11-07/4db06060d78340a29c18a0436d9eaa56.webp",
	}

	embed := discordgo.MessageEmbed{
		URL:         "https://vitalitysouth.com/",
		Title:       "Clock-In",
		Description: "",
		Color:       5763719,
		Image:       &image,
	}

	return &embed
}

// clockOutEmbed is the set up for the clock-out embed
func clockOutEmbed() *discordgo.MessageEmbed {

	image := discordgo.MessageEmbedImage{
		URL: "https://img.craiyon.com/2023-11-20/lwkWz-yhSRKqMl38plwCqw.webp",
	}

	embed := &discordgo.MessageEmbed{
		URL:         "https://vitalitysouth.com/",
		Title:       "Clock-Out",
		Description: "",
		Color:       15548997,
		Image:       &image,
	}

	return embed
}

// ClockInResponse is the message the bot sends and the actions it takes whenever is being used
func ClockInResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channelID := "1172648319940558970"
	_, ok := session.ChannelMessageSend(channelID, "Stransyyy bot esta siendo usado...")
	if ok != nil {
		log.Fatal("Is the error here?", ok)
	}

	// Check if interaction or interaction.Interaction is nil
	if interaction == nil || interaction.Interaction == nil {
		log.Println("Invalid interaction object")
		return
	}

	userID := func() string {
		if interaction.User != nil {
			return interaction.User.ID
		}

		if interaction.Member != nil {
			return interaction.Member.User.ID
		}

		return ""
	}()

	// Respond to the slash command interaction with a deferred response
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You have clocked-in succesfully!",
		},
	})
	if err != nil {
		log.Println("Error responding to interaction:", err)
		return
	}

	//FollowUp message after you use the command
	mensaje := &discordgo.WebhookParams{
		Content: "You have Clocked-in successfuly",
	}

	_, err = session.FollowupMessageCreate(interaction.Interaction, false, mensaje)
	if err != nil {
		return
	}

	// Create a DM channel with the user who used the command
	dmChannel, err := session.UserChannelCreate(userID)
	if err != nil {
		log.Println("Error creating DM channel:", err)
		return
	}

	// Send a message with an embed to the user in the DM channel
	if dmChannel != nil && dmChannel.ID != "" {
		_, dmErr := session.ChannelMessageSendEmbed(dmChannel.ID, clockInEmbed())

		if dmErr != nil {
			log.Println("Error sending DM with embed:", dmErr)
		}
	} else {
		log.Println("Invalid DM channel")
	}
}

// clockOutResponse will send an embed as a response to the slash command call
func ClockOutResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	channelID := "1172648319940558970"
	_, ok := session.ChannelMessageSend(channelID, "Stransyyy bot esta siendo usado...")
	if ok != nil {
		log.Fatal("Is the error here?", ok)
	}

	// Check if interaction or interaction.Interaction is nil
	if interaction == nil || interaction.Interaction == nil {
		log.Println("Invalid interaction object")
		return
	}

	userID := func() string {
		if interaction.User != nil {
			return interaction.User.ID
		}

		if interaction.Member != nil {
			return interaction.Member.User.ID
		}

		return ""
	}()

	// Respond to the slash command interaction with a deferred response
	err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You have clocked-Out succesfully!",
		},
	})
	if err != nil {
		log.Println("Error responding to interaction:", err)
		return
	}

	//FollowUp message after you use the command
	mensaje := &discordgo.WebhookParams{
		Content: "You have Clocked-Out successfuly, you can now rest!",
	}

	_, err = session.FollowupMessageCreate(interaction.Interaction, false, mensaje)
	if err != nil {
		return
	}

	// Create a DM channel with the user who used the command
	dmChannel, err := session.UserChannelCreate(userID)
	if err != nil {
		log.Println("Error creating DM channel:", err)
		return
	}

	// Insert inputs to the database whenever the slash command is being used
	result, updatedSession, dbErr := DiscordDataBase(nil, nil, session)
	if dbErr != nil {
		log.Println("Error storing data in the database:", dbErr)
	}

	// Check the result of the database operation
	if result != nil {
		log.Println("Database operation successful:", result)
	}

	// Use the updated session if it's not nil
	if updatedSession != nil {
		log.Println("Session updated successfully")
	}

	// Send a message with an embed to the user in the DM channel (this will be something else)
	if dmChannel != nil && dmChannel.ID != "" {
		_, dmErr := session.ChannelMessageSendEmbed(dmChannel.ID, clockOutEmbed())

		if dmErr != nil {
			log.Println("Error sending DM with embed:", dmErr)
		}
	} else {
		log.Println("Invalid DM channel")
	}
}

func DiscordDataBase(db *sql.DB, message *discordgo.MessageCreate, session *discordgo.Session) (sql.Result, *discordgo.Session, error) {
	// Access the channel information
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting channel information: %v", err)
	}

	channelName := channel.Name
	messageID := message.ID
	authorID := message.Author.ID
	messageContent := message.Content
	messageTimestamp := message.Timestamp
	formattedTimestamp := messageTimestamp.Format(time.RFC3339)
	channelID := message.ChannelID

	// Use a transaction for atomicity
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, fmt.Errorf("error beginning database transaction: %v", err)
	}
	defer tx.Rollback()

	// Execute each query individually within the transaction
	result, err := tx.Exec("INSERT INTO messages (message_id, author_id, message_content, date_sent, time_sent) VALUES (?, ?, ?, ?, ?)", messageID, authorID, messageContent, messageTimestamp, formattedTimestamp)
	if err != nil {
		return nil, nil, fmt.Errorf("error inserting into messages table: %v", err)
	}

	_, err = tx.Exec("INSERT INTO users (author_id) VALUES (?)", authorID)
	if err != nil {
		return nil, nil, fmt.Errorf("error inserting into users table: %v", err)
	}

	_, err = tx.Exec("INSERT INTO channels (channel_id, channel_name) VALUES (?, ?)", channelID, channelName)
	if err != nil {
		return nil, nil, fmt.Errorf("error inserting into channels table: %v", err)
	}

	// Commit the transaction if all queries succeed
	err = tx.Commit()
	if err != nil {
		return nil, nil, fmt.Errorf("error committing database transaction: %v", err)
	}

	return result, session, nil
}
