// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/bwmarrin/discordgo"
// 	_ "github.com/go-sql-driver/mysql"
// )

// // Database connection
// var db *sql.DB

// func init() {
// 	var err error

// 	// Initialize the database connection
// 	db, err = sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Ensure the database connection is valid
// 	if err = db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func main() {
// 	// Create a new Discord session
// 	dg, err := discordgo.New("Bot " + "YOUR_BOT_TOKEN")
// 	if err != nil {
// 		fmt.Println("error creating Discord session,", err)
// 		return
// 	}

// 	// Register the interactionCreate func to handle incoming slash commands
// 	dg.AddInteractionCreateHandler(interactionCreate)

// 	// Open a websocket connection to Discord and begin listening
// 	err = dg.Open()
// 	if err != nil {
// 		fmt.Println("error opening connection,", err)
// 		return
// 	}

// 	// Wait here until CTRL-C or other term signal is received
// 	fmt.Println("Bot is now running. Press CTRL-C to exit.")
// 	select {}
// }

// func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	// Check if the interaction is a slash command
// 	if i.Type == discordgo.InteractionApplicationCommand {
// 		// Handle the specific command
// 		switch i.ApplicationCommandData().Name {
// 		case "lastinsertid":
// 			// Define the parameters for your specific query
// 			serverID := "your_server_id"
// 			channelID := "your_channel_id"
// 			messageID := "your_message_id"
// 			authorID := "your_author_id"
// 			content := "your_message_content"
// 			dateSent := time.Now()

// 			// Perform the database operation using your existing function
// 			err := messagesDataBaseClockInHandler(db, serverID, channelID, messageID, authorID, content, dateSent)
// 			if err != nil {
// 				log.Println("Error storing data in the database:", err)
// 				return
// 			}

// 			// Use getLastInsertID to retrieve the last insert ID
// 			lastInsertID, err := getLastInsertID("SELECT LAST_INSERT_ID()")
// 			if err != nil {
// 				log.Println("Error getting last insert ID:", err)
// 				return
// 			}

// 			// Create an embed with the database info
// 			embed := &discordgo.MessageEmbed{
// 				Title: "Database Info",
// 				Fields: []*discordgo.MessageEmbedField{
// 					{
// 						Name:   "Last Insert ID",
// 						Value:  fmt.Sprintf("%d", lastInsertID),
// 						Inline: true,
// 					},
// 					// Add more fields as needed for other database information
// 				},
// 			}

// 			// Respond to the slash command with the embed
// 			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 				Type: discordgo.InteractionResponseChannelMessageWithSource,
// 				Data: &discordgo.InteractionApplicationCommandResponseData{
// 					Embeds: []*discordgo.MessageEmbed{embed},
// 				},
// 			})
// 		}
// 	}
// }

// func getLastInsertID(query string) (int64, error) {
// 	// Replace this with your actual database query
// 	result, err := db.Exec(query)
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Get the last insert ID from the Result interface
// 	lastInsertID, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

//		return lastInsertID, nil
//	}
package dc
