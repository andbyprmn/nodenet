package api

import (
	"encoding/json"
	"log"
	"net/http"
	"nodenet/internal/data"
	"nodenet/internal/logging"
)

type APIServer struct {
	store  *data.KeyValueStore
	logger *logging.Logger
}

func NewAPIServer(store *data.KeyValueStore, logger *logging.Logger) *APIServer {
	return &APIServer{
		store:  store,
		logger: logger,
	}
}

// getHandler menangani request GET key.
func (api *APIServer) getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}
	val, err := api.store.Get(key)
	if err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{key: val})
}

// getAllHandler menangani request untuk mendapatkan semua key-value.
func (api *APIServer) getAllHandler(w http.ResponseWriter, r *http.Request) {
	values := api.store.GetAll()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(values)
}

// setHandler menangani request SET key.
func (api *APIServer) setHandler(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	for key, val := range req {
		api.store.Set(key, val)
	}
	w.WriteHeader(http.StatusOK)
}

// deleteHandler menangani request DELETE key.
func (api *APIServer) deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key is required", http.StatusBadRequest)
		return
	}
	if err := api.store.Delete(key); err != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// logHandler mengembalikan baris terakhir dari log.
func (api *APIServer) logHandler(w http.ResponseWriter, r *http.Request) {
	lines := api.logger.GetLastLogs(10) // Mendapatkan 10 baris terakhir
	json.NewEncoder(w).Encode(lines)
}

// StartServer memulai server API di port yang ditentukan.
func (api *APIServer) StartServer(port string) {
	http.HandleFunc("/get", api.getHandler)
	http.HandleFunc("/getAll", api.getAllHandler)
	http.HandleFunc("/set", api.setHandler)
	http.HandleFunc("/delete", api.deleteHandler)
	http.HandleFunc("/logs", api.logHandler)

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
