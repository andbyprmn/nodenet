package tests

import (
	"log"
	"nodenet/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var store *model.Store

func TestMain(m *testing.M) {
	var err error
	store, err = model.NewStore("localhost", "5432", "user", "password", "dbname") // Ganti dengan detail database Anda
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	code := m.Run()
	os.Exit(code)
}

func TestSetGetKey(t *testing.T) {
	kv := model.KeyValue{Key: "testKey", Value: "testValue"}
	err := store.SetKey(kv)
	assert.NoError(t, err)

	value, err := store.GetKey("testKey")
	assert.NoError(t, err)
	assert.Equal(t, "testValue", value)
}

func TestDeleteKey(t *testing.T) {
	kv := model.KeyValue{Key: "deleteKey", Value: "deleteValue"}
	err := store.SetKey(kv)
	assert.NoError(t, err)

	err = store.DeleteKey("deleteKey")
	assert.NoError(t, err)

	_, err = store.GetKey("deleteKey")
	assert.Error(t, err)
}

func TestGetAllKeys(t *testing.T) {
	store.SetKey(model.KeyValue{Key: "key1", Value: "value1"})
	store.SetKey(model.KeyValue{Key: "key2", Value: "value2"})

	allData, err := store.GetAllKeys()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(allData))
}
