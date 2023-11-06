package main

import (
	"flag"
	"fmt"

	"github.com/Stransyyy/Sheet-Linker/db"
)

func main() {
	cred, err := db.JsonFileReader("credentials.json")
	if err != nil {
		return
	}

	fmt.Println("Welcome to MySQL")

	con, err := db.Connection(cred)
	if err != nil {
		panic(err)
	}

	db.ScanTableInputs(con)

	fmt.Println("inputs added ")

	// Discord Stuff
	var Token string

	flag.StringVar(&Token, "t", "", "Bot token")
	flag.Parse()

	if Token == "" {
		fmt.Println("Please provide a bot using the -t flag.")
		return
	}

	hola := Runbot("fdf442234")

	fmt.Print(hola)

	// Closes connection to the databse
	err = db.CloseDB(con)
	if err != nil {
		panic(err)
	}
}
