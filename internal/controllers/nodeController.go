package controllers

import (
	"encoding/json"
	"net/http"
	"nodenet/internal/models"
	"nodenet/internal/services"
)

// NodeController handles the HTTP requests related to node operations
type NodeController struct {
	service *services.NodeService
}

// NewNodeController creates a new instance of NodeController
func NewNodeController(service *services.NodeService) *NodeController {
	return &NodeController{
		service: service,
	}
}

// GetValue handles GET request to retrieve a value by key
func (c *NodeController) GetValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, err := c.service.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{key: value})
}

// SetValue handles POST request to set a value by key
func (c *NodeController) SetValue(w http.ResponseWriter, r *http.Request) {
	var node models.Node
	if err := json.NewDecoder(r.Body).Decode(&node); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	c.service.Set(node.Key, node.Value)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// DeleteValue handles DELETE request to remove a key-value pair
func (c *NodeController) DeleteValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if err := c.service.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

// GetAllValues handles GET request to retrieve all key-value pairs
func (c *NodeController) GetAllValues(w http.ResponseWriter, r *http.Request) {
	values := c.service.GetAll()
	json.NewEncoder(w).Encode(values)
}
