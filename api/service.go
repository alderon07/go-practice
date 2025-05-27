package main

import (
	"fmt"
	"strconv"
	"sync"
)

// PlayerService handles player-related operations with thread safety
type PlayerService struct {
	mu   sync.RWMutex
	data map[string]Player
	idCounter int
}

// NewPlayerService creates a new PlayerService with sample data
func NewPlayerService() *PlayerService {
	service := &PlayerService{
		data: make(map[string]Player),
		idCounter: 0,
	}
	
	// Add sample data
	samplePlayers := []Player{
		{ID: "1", Name: "Messi", JerseyNumber: 10, Rating: 99},
		{ID: "2", Name: "Ronaldo", JerseyNumber: 7, Rating: 98},
		{ID: "3", Name: "Neymar", JerseyNumber: 10, Rating: 95},
	}
	
	for _, player := range samplePlayers {
		service.data[player.ID] = player
		if id, err := strconv.Atoi(player.ID); err == nil && id > service.idCounter {
			service.idCounter = id
		}
	}
	
	return service
}

// GetAllPlayers returns all players
func (s *PlayerService) GetAllPlayers() []Player {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	players := make([]Player, 0, len(s.data))
	for _, player := range s.data {
		players = append(players, player)
	}
	return players
}

// GetPlayerByID returns a player by ID
func (s *PlayerService) GetPlayerByID(id string) (Player, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	player, exists := s.data[id]
	if !exists {
		return Player{}, ErrPlayerNotFound
	}
	return player, nil
}

// CreatePlayer creates a new player
func (s *PlayerService) CreatePlayer(req PlayerRequest) (Player, error) {
	if err := req.Validate(); err != nil {
		return Player{}, err
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// Check if a player with same name and jersey number already exists
	for _, player := range s.data {
		if player.Name == req.Name && player.JerseyNumber == req.JerseyNumber {
			return Player{}, fmt.Errorf("%w: player with name %s and jersey number %d already exists", 
				ErrPlayerExists, req.Name, req.JerseyNumber)
		}
	}
	
	s.idCounter++
	id := strconv.Itoa(s.idCounter)
	player := req.ToPlayer(id)
	s.data[id] = player
	
	return player, nil
}

// UpdatePlayer updates an existing player
func (s *PlayerService) UpdatePlayer(id string, req PlayerRequest) (Player, error) {
	if err := req.Validate(); err != nil {
		return Player{}, err
	}
	
	s.mu.Lock()
	defer s.mu.Unlock()
	
	player, exists := s.data[id]
	if !exists {
		return Player{}, ErrPlayerNotFound
	}
	
	// Check if another player has the same name and jersey number
	for existingID, existingPlayer := range s.data {
		if existingID != id && existingPlayer.Name == req.Name && existingPlayer.JerseyNumber == req.JerseyNumber {
			return Player{}, fmt.Errorf("%w: another player with name %s and jersey number %d already exists", 
				ErrPlayerExists, req.Name, req.JerseyNumber)
		}
	}
	
	player.Update(req)
	s.data[id] = player
	
	return player, nil
}

// DeletePlayer deletes a player by ID
func (s *PlayerService) DeletePlayer(id string) (Player, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	player, exists := s.data[id]
	if !exists {
		return Player{}, ErrPlayerNotFound
	}
	
	delete(s.data, id)
	return player, nil
}

// PlayerExists checks if a player exists by ID
func (s *PlayerService) PlayerExists(id string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	_, exists := s.data[id]
	return exists
} 