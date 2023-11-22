package data

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

type CustomerRow struct {
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
func ScanTableInputs(db *sql.DB) ([]CustomerRow, []string, error) {

	rows, err := db.Query("SELECT * FROM Customer;")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rowsData := make([]CustomerRow, len(columns))

	for rows.Next() {
		var rowData CustomerRow
		err = rows.Scan(&rowData.ID, &rowData.NAME, &rowData.EMAIL, &rowData.DATE_CREATED)
		if err != nil {
			panic(err.Error())
		}
		rowsData = append(rowsData, rowData)
	}
	fmt.Print(rowsData)
	return rowsData, columns, err
}

/*
// first create struct representing row
	// create dynamic slice of row type to hold rows data, return
	// do for loop calling rows.Next()
	// in loop create temprows var
	// run rows.Scan(&temprows)
	// then append data to rows slice
	// print rows data
*/

// Reads the json file
func JsonFileReader(credentials string) (ConnectionCredentials, error) {

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

	return db, err
}

func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
