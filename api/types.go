package main

import (
  "errors"
  "fmt"
)

// Response represents the standard API response structure
type Response struct {
  Message string      `json:"message"`
  Status  string      `json:"status"`
  Data    interface{} `json:"data,omitempty"`
  Error   string      `json:"error,omitempty"`
}

// Player represents a football player
type Player struct {
  ID           string `json:"id"`
  Name         string `json:"name"`
  JerseyNumber int8   `json:"jersey_number"`
  Rating       int8   `json:"rating"`
}

// PlayerRequest represents the request structure for creating/updating players
type PlayerRequest struct {
  Name         string `json:"name"`
  JerseyNumber int8   `json:"jersey_number"`
  Rating       int8   `json:"rating"`
}

// Custom errors
var (
  ErrPlayerNotFound    = errors.New("player not found")
  ErrInvalidInput      = errors.New("invalid input")
  ErrPlayerExists      = errors.New("player already exists")
  ErrInvalidJSONFormat = errors.New("invalid JSON format")
)

// Validate validates the player request data
func (pr *PlayerRequest) Validate() error {
  if pr.Name == "" {
    return fmt.Errorf("%w: name is required", ErrInvalidInput)
  }
  if pr.JerseyNumber < 1 || pr.JerseyNumber > 99 {
    return fmt.Errorf("%w: jersey number must be between 1 and 99", ErrInvalidInput)
  }
  if pr.Rating < 1 || pr.Rating > 99 {
    return fmt.Errorf("%w: rating must be between 1 and 99", ErrInvalidInput)
  }
  return nil
}

// ToPlayer converts PlayerRequest to Player with given ID
func (pr *PlayerRequest) ToPlayer(id string) Player {
  return Player{
    ID:           id,
    Name:         pr.Name,
    JerseyNumber: pr.JerseyNumber,
    Rating:       pr.Rating,
  }
}

// Update updates the player with new data
func (p *Player) Update(req PlayerRequest) {
  if req.Name != "" {
    p.Name = req.Name
  }
  if req.JerseyNumber > 0 {
    p.JerseyNumber = req.JerseyNumber
  }
  if req.Rating > 0 {
    p.Rating = req.Rating
  }
}
