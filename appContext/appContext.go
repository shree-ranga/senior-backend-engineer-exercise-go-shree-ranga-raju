package appcontext

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type appContext struct {
	dbClient *sql.DB
}

var appCtx *appContext

func Init() {
	appCtx = &appContext{}
	appCtx.dbClient = newDBConnection()
}

func newDBConnection() *sql.DB {
	// Connect to SQLite database file
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Successfully connected to SQLite database!")
	return db
}

func GetDB() *sql.DB {
	return appCtx.dbClient
}
