package main

import (
	"log"
	"os"

	"github.com/Stransyyy/Sheet-Linker/dc"
)

func main() {
	// cred, err := db.JsonFileReader("credentials.json")
	// if err != nil {
	// 	return
	// }

	// fmt.Println("Welcome to MySQL")

	// con, err := db.Connection(cred)
	// if err != nil {
	// 	panic(err)
	// }

	// db.ScanTableInputs(con)

	// fmt.Println("inputs added ")

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	botToken, ok := os.LookupEnv("DISCORD_APIKEY")
	if !ok {
		log.Fatal("Must set Discord token as env variable: DISCORD_APIKEY")
	}

	dc.BotToken = botToken
	dc.Run()

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Closes connection to the databse
	// err = db.CloseDB(con)
	// if err != nil {
	// 	panic(err)
	// }
}
