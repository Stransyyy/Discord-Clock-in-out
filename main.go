package main

import (
	"fmt"
	"log"
	"os"

	data "github.com/Stransyyy/Sheet-Linker/db"
	"github.com/Stransyyy/Sheet-Linker/dc"
)

func main() {
	cred, err := data.JsonFileReader("credentials.json")
	if err != nil {
		return
	}

	fmt.Println("Welcome to MySQL")

	con, err := data.Connection(cred)
	if err != nil {
		panic(err)
	}

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
	dc.Run()

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	//Closes connection to the database
	defer data.CloseDB(con)
}
