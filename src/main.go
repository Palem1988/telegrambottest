package main

import (
	"database/sql"
	"fmt"
	"os"
	stct "telegrambottest/src/bipdev/structs"
	"telegrambottest/src/bot"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := stct.Config{
		URL:        os.Getenv("URL"),
		Token:      os.Getenv("TOKEN"),
		Driver:     os.Getenv("DRIVER"),
		DataSource: os.Getenv("DATASOURCE"),
	}

	dbsql, err := sql.Open(conf.Driver, conf.DataSource)
	if err != nil {
		log.Fatal(err)
	}

	defer dbsql.Close()

	// Inizializaton users from DB, token for bot.
	bot := bot.InitBot(conf, dbsql)
	// Run bot
	fmt.Println("Bot started!")
	bot.Run()

}
