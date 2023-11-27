package main

import (
	"log"
	"os"

	data "github.com/Stransyyy/Sheet-Linker/db"
	"github.com/Stransyyy/Sheet-Linker/dc"
)

func main() {
	//Load MySQL credentials for the connection
	cred, err := data.JsonFileReader("credentials.json")
	if err != nil {
		return
	}
	//Stablis connection with the database
	con, err := data.Connection(cred)
	if err != nil {
		log.Fatal("Error establishing connection to the database:", err)
	}

	//Closes connection to the database
	defer data.CloseDB(con)

	//data.ScanTableInputs(con)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	botToken, ok := os.LookupEnv("DISCORD_APIKEY")
	if !ok {
		log.Fatal("Must set Discord token as env variable: DISCORD_APIKEY")
	}

	stransyyyBotChanneId, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Must set Discord token as env variable: DISCORD_APIKEY")
	}

	dc.BotToken = botToken
	dc.StransyyyBotChanneId = stransyyyBotChanneId
	dc.Run(con)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
}
