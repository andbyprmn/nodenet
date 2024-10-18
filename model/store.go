package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

var keyValueStore = make(map[string]string)

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=nodenet sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetKeyValue(key string) (string, error) {
	if value, ok := keyValueStore[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("key not found")
}

func SetKeyValue(key, value string) error {
	keyValueStore[key] = value
	_, err := db.Exec("INSERT INTO kv_store (key, value) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET value = $2", key, value)
	return err
}
