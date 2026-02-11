package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB(dataSourceName string) *sql.DB {
	db, err := sql.Open("sqlite", dataSourceName)

	//Failure log
	if err != nil {
		log.Fatalf("[DB] Open Failure: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("[DB] Ping Failure: %v", err)
	}

	//Run migrations and seed demo data
	runMigrations(db)

	//Success log
	log.Println("[DB] Initialised")
	return db
}

func runMigrations(db *sql.DB) {
	//Create accounts table, note to self this is similar to flyway in spring
	createAccountsTable := `
	CREATE TABLE IF NOT EXISTS accounts (
		user_id TEXT PRIMARY KEY,
		balance INTEGER NOT NULL DEFAULT 0
	);`

	//Failure log
	if _, err := db.Exec(createAccountsTable); err != nil {
		log.Fatalf("[DB] Create Accounts Table Failure: %v", err)
	}

	//Demo account for testing the API
	seed := `INSERT OR IGNORE INTO accounts (user_id, balance) VALUES ('evaristeGalois', 322322322);`
	if _, err := db.Exec(seed); err != nil {
		log.Fatalf("[DB] Account Seeding Failure: %v", err)
	}
}
