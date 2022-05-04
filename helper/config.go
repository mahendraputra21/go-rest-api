package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
)

//create connection with posgres
func CreateConnection() *sql.DB {
	//load .env file
	err := godotenv.Load(os.ExpandEnv("go-rest-api/.env")) //please set path for the .env if you got the error message

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//open connection to DB
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	//check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to DB Success!")
	//return the connection
	return db
}
