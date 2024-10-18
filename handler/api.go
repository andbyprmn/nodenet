package handler

import (
	"encoding/json"
	"net/http"
	"nodenet/model"

	"github.com/julienschmidt/httprouter"
)

// CommandResponse structure for commands
type CommandResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// API struct untuk menyimpan instance Store
type API struct {
	Store *model.Store
}

// NewAPI untuk membuat instance API
func NewAPI(store *model.Store) *API {
	return &API{Store: store}
}

// SetKeyHandler untuk menyimpan kunci-nilai (set command)
func (api *API) SetKeyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var kv model.KeyValue
	if err := json.NewDecoder(r.Body).Decode(&kv); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := api.Store.SetKey(kv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CommandResponse{
		Message: "Key set successfully",
	})
}

// GetKeyHandler untuk mengambil nilai berdasarkan kunci (get command)
func (api *API) GetKeyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	value, err := api.Store.GetKey(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CommandResponse{
		Message: "Value retrieved successfully",
		Data:    value,
	})
}

// GetAllKeysHandler untuk mengambil semua kunci-nilai (get-all command)
func (api *API) GetAllKeysHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allData, err := api.Store.GetAllKeys()
	if err != nil {
		http.Error(w, "Error retrieving data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CommandResponse{
		Message: "All key-value pairs retrieved successfully",
		Data:    allData,
	})
}

// DeleteKeyHandler untuk menghapus kunci (delete command)
func (api *API) DeleteKeyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	if err := api.Store.DeleteKey(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CommandResponse{
		Message: "Key deleted successfully",
	})
}

// RegisterRoutes untuk mendaftarkan rute-rute API
func (api *API) RegisterRoutes(router *httprouter.Router) {
	router.POST("/api/v1/keys", api.SetKeyHandler)
	router.GET("/api/v1/keys/:key", api.GetKeyHandler)
	router.GET("/api/v1/keys", api.GetAllKeysHandler)
	router.DELETE("/api/v1/keys/:key", api.DeleteKeyHandler)
}
