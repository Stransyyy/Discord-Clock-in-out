package dc

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func lastInsertedCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "Last-Inserted",
		Description: "Use this command to get the last inserted input into the database and get in response an embed with the info!",
	}
}

// lastInsertedResponse is the message the bot sends and the actions it takes whenever is being used
func lastInsertedResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate, db *sql.DB) {
	if db == nil {
		log.Fatal("Error db is nil in clockinresponse")
		return
	}

	channelID := "1172648319940558970"
	_, err := session.ChannelMessageSend(channelID, "Stransyyy bot esta siendo usado...")
	if err != nil {
		log.Fatal("Is the error here?", err)
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
	err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Data pulled succesfully",
		},
	})
	if err != nil {
		log.Println("Error responding to interaction:", err)
		return
	}

	//FollowUp message after you use the command
	mensaje := &discordgo.WebhookParams{
		Content: "Data pulled successfully",
	}

	_, err = session.FollowupMessageCreate(interaction.Interaction, false, mensaje)
	if err != nil {
		return
	}

	//Stores the data into the database
	messageId := interaction.ID
	authorId := interaction.Member.User.ID
	msgContent := "/lastinserted"
	dateSent, err := discordgo.SnowflakeTimestamp(interaction.ID)
	serverId := interaction.GuildID
	channelId := interaction.ChannelID

	if err != nil {
		log.Println("dateSent snowflake error ", err)
		dateSent = time.Now()
	}

	err, res := messagesDataBaseClockInHandler(db, serverId, channelId, messageId, authorId, msgContent, dateSent)
	if err != nil {
		log.Println("Error storing data in the database", err)
	}

	_, err = getLastInsertID("SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Println("Error getting last inserted ID", err)
		return
	}

	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		log.Println("Couldn't get the lsat inserted id", err)
	}

	// This is the embed that will display the values of the databse
	embed := &discordgo.MessageEmbed{
		Title: "Database Info",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "Last Inserted ID",
				Value: fmt.Sprintf("%d", lastInsertedID),
			},
		},
	}

	// Create a DM channel with the user who used the command
	dmChannel, err := session.UserChannelCreate(userID)
	if err != nil {
		log.Println("Error creating DM channel:", err)
		return
	}

	// Send a message with an embed to the user in the DM channel
	if dmChannel != nil && dmChannel.ID != "" {
		_, dmErr := session.ChannelMessageSendEmbed(dmChannel.ID, embed)

		if dmErr != nil {
			log.Println("Error sending DM with embed:", dmErr)
		}
	} else {
		log.Println("Invalid DM channel")
	}
}

func getLastInsertID(query string) (int64, error) {
	var db *sql.DB

	// Replace this with your actual database query
	result, err := db.Exec(query)
	if err != nil {
		return 0, err
	}

	// Get the last insert ID from the Result interface
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}
