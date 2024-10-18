package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nodenet/handler"
	"nodenet/model"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestSetKeyHandler(t *testing.T) {
	router := httprouter.New()
	api := handler.NewAPI(store)
	api.RegisterRoutes(router)

	kv := model.KeyValue{Key: "testKey", Value: "testValue"}
	body, _ := json.Marshal(kv)
	req, _ := http.NewRequest("POST", "/api/v1/keys", bytes.NewBuffer(body))
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response handler.CommandResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Key set successfully", response.Message)
}

func TestGetKeyHandler(t *testing.T) {
	router := httprouter.New()
	api := handler.NewAPI(store)
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/api/v1/keys/testKey", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response handler.CommandResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Value retrieved successfully", response.Message)
	assert.Equal(t, "testValue", response.Data)
}

func TestGetAllKeysHandler(t *testing.T) {
	router := httprouter.New()
	api := handler.NewAPI(store)
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("GET", "/api/v1/keys", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response handler.CommandResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "All key-value pairs retrieved successfully", response.Message)
}

func TestDeleteKeyHandler(t *testing.T) {
	router := httprouter.New()
	api := handler.NewAPI(store)
	api.RegisterRoutes(router)

	req, _ := http.NewRequest("DELETE", "/api/v1/keys/testKey", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response handler.CommandResponse
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(t, "Key deleted successfully", response.Message)
}
