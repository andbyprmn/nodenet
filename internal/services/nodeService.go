package services

import (
	"errors"
	"nodenet/internal/logging"
)

// NodeService handles the business logic for node operations
type NodeService struct {
	data   map[string]string
	logger *logging.Logger
}

// NewNodeService creates a new instance of NodeService
func NewNodeService(initialData map[string]string, logger *logging.Logger) *NodeService {
	return &NodeService{
		data:   initialData,
		logger: logger,
	}
}

// Get retrieves a value by key
func (s *NodeService) Get(key string) (string, error) {
	value, exists := s.data[key]
	if !exists {
		s.logger.LogEvent("GET failed: key=" + key + " not found")
		return "", errors.New("key not found")
	}
	s.logger.LogEvent("GET key=" + key)
	return value, nil
}

// Set stores a value by key
func (s *NodeService) Set(key, value string) error {
	s.data[key] = value
	s.logger.LogEvent("SET key=" + key + " value=" + value)
	return nil
}

// Delete removes a key-value pair
func (s *NodeService) Delete(key string) error {
	_, exists := s.data[key]
	if !exists {
		s.logger.LogEvent("DELETE failed: key=" + key + " not found")
		return errors.New("key not found")
	}
	delete(s.data, key)
	s.logger.LogEvent("DELETE key=" + key)
	return nil
}

// GetAll retrieves all key-value pairs
func (s *NodeService) GetAll() map[string]string {
	s.logger.LogEvent("GET all values")
	return s.data
}
