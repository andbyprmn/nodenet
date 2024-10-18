package model

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

// KeyValue struct untuk menyimpan pasangan kunci-nilai
type KeyValue struct {
	Key   string
	Value string
}

// Store struct untuk koneksi database
type Store struct {
	DB *sql.DB
}

// NewStore untuk membuat instance Store
func NewStore(host, port, user, password, dbname string) (*Store, error) {
	psqlInfo := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}

// SetKey untuk menyimpan kunci-nilai ke database
func (s *Store) SetKey(kv KeyValue) error {
	_, err := s.DB.Exec("INSERT INTO kv_store (key, value) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET value = EXCLUDED.value", kv.Key, kv.Value)
	return err
}

// GetKey untuk mengambil nilai berdasarkan kunci dari database
func (s *Store) GetKey(key string) (string, error) {
	var value string
	err := s.DB.QueryRow("SELECT value FROM kv_store WHERE key = $1", key).Scan(&value)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("key not found")
		}
		return "", err
	}
	return value, nil
}

// GetAllKeys untuk mengambil semua kunci-nilai dari database
func (s *Store) GetAllKeys() (map[string]string, error) {
	rows, err := s.DB.Query("SELECT key, value FROM kv_store")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	kvMap := make(map[string]string)
	for rows.Next() {
		var kv KeyValue
		if err := rows.Scan(&kv.Key, &kv.Value); err != nil {
			return nil, err
		}
		kvMap[kv.Key] = kv.Value
	}
	return kvMap, nil
}

// DeleteKey untuk menghapus kunci dari database
func (s *Store) DeleteKey(key string) error {
	_, err := s.DB.Exec("DELETE FROM kv_store WHERE key = $1", key)
	return err
}
