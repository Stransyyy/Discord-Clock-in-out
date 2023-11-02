package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type Connection struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}

func jsonFileReader(configFile string) (Connection, error) {

	// f stores the data from the json file
	var f Connection

	data, derr := os.ReadFile(configFile)

	if derr != nil {
		return f, derr
	}

	err := json.Unmarshal(data, &f)

	if err != nil {
		return f, err
	}

	return f, err

}

func main() {
	fmt.Println("Mysql tutorial")
	db, err := sql.Open("mysql", "dbuser1:test@tcp(127.0.0.1:3306)/Test")
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected")
	// See "Important settings" section.
	defer db.Close()

	insert, err := db.Query("INSERT INTO Customer (NAME, EMAIL, DATE_CREATED) VALUES ('Roberto', 'robert@gmail.com', '2008-11-11');")

	if err != nil {
		panic(err)
	}

	defer insert.Close()
	fmt.Println("Successfully inserted")
}
