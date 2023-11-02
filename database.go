package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ...

type ConnectionCredentials struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Database string `json:"Database"`
}

type CustomerTable struct {
	ID           int    `json:"ID"`
	NAME         string `json:"NAME"`
	EMAIL        string `json:"EMAIL"`
	DATE_CREATED string `json:"DATE_CREATED"`
}

/*
	    results, err := db.Query("SELECT id FROM testtable2")
	    if err !=nil {
	        panic(err.Error())
	    }
	    for results.Next() {
	        var testtable2 Testtable2
	        err = results.Scan(&testtable2.id)
	        if err !=nil {
	            panic(err.Error())
	        }
	        fmt.Println(testtable2.id)
	    }
	}
*/
func scanTableInputs(db *sql.DB) ([]CustomerTable, []string, error) {

	rows, err := db.Query("SELECT * FROM Customer;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rowsData := make([]CustomerTable, len(columns))

	for rows.Next() {
		var rowData CustomerTable
		err = rows.Scan(&rowData)
		if err != nil {
			panic(err.Error())
		}

		rowsData = append(rowsData, rowData)
		// fmt.Println(testtable2.id)
	}
	return rowsData, columns, nil
}

/*
// first create struct representing row
	// create dynamic slice of row type to hold rows data, return
	// do for loop calling rows.Next()
	// in loop create temprows var
	// run rows.Scan(&temprows)
	// then append data to rows slice
	// print rows data

	// values := make([]sql.RawBytes, len(columns))

	// if err != nil {
	// 	db.Close()
	// 	return nil, err
	// }

	// // db.Ping checks the
	// err = db.Ping()
	// if err != nil {
	// 	db.Close()
	// 	return nil, err
	// }
*/

// Reads the json file
func jsonFileReader(credentials string) (ConnectionCredentials, error) {

	// f stores the data from the json file
	var ConnectionInfo ConnectionCredentials

	data, derr := os.ReadFile(credentials)

	if derr != nil {
		return ConnectionInfo, derr
	}

	err := json.Unmarshal(data, &ConnectionInfo)

	if err != nil {
		return ConnectionInfo, err
	}

	return ConnectionInfo, err

}

// Connection makes the connection to the database using the "github.com/go-sql-driver/mysql" package
func Connection(y ConnectionCredentials) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", y.Username, y.Password, y.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected")

	// fmt.Print(rows)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Print(columns)

	defer db.Close()

	return db, nil
}

func main() {

	cred, err := jsonFileReader("credentials.json")

	if err != nil {
		return
	}
	fmt.Println("Welcome to MySQL")
	scanTableInputs(&sql.DB{})
	Connection(cred)
	fmt.Printf("\n")
}
