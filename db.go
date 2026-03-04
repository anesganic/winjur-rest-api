package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	db               *sqlx.DB
	validTokenHashes map[string]bool
)

func InitDB() {
	validTokenHashes = make(map[string]bool)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if strings.HasPrefix(parts[0], "API_TOKEN_HASH_") {
			validTokenHashes[parts[1]] = true
		}
	}

	server := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	if server == "" || user == "" || len(validTokenHashes) == 0 {
		log.Println("WARNUNG: Env Variablen oder Tokens fehlen.")
	}

	query := url.Values{}
	query.Add("database", database)
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(user, password),
		Host:     server,
		RawQuery: query.Encode(),
	}

	var err error
	db, err = sqlx.Connect("sqlserver", u.String())
	if err != nil {
		log.Fatalf("DB Fehler: %v", err)
	}

	fmt.Println("✓ DB verbunden (sqlx)")
}
