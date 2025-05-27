package main

import (
  "encoding/json"
  "errors"
  "log"
  "net/http"
)

// PlayerHandler contains the player service and HTTP handlers
type PlayerHandler struct {
  service *PlayerService
}

// NewPlayerHandler creates a new PlayerHandler
func NewPlayerHandler(service *PlayerService) *PlayerHandler {
  return &PlayerHandler{service: service}
}

// sendJSONResponse is a helper function to send JSON responses
func (h *PlayerHandler) sendJSONResponse(w http.ResponseWriter, status int, response Response) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  
  if err := json.NewEncoder(w).Encode(response); err != nil {
    log.Printf("Error encoding JSON response: %v", err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
  }
}

// sendErrorResponse is a helper function to send error responses
func (h *PlayerHandler) sendErrorResponse(w http.ResponseWriter, status int, message string, err error) {
  response := Response{
    Status:  "error",
    Message: message,
    Error:   err.Error(),
  }
  
  log.Printf("Error: %s - %v", message, err)
  h.sendJSONResponse(w, status, response)
}

// GetPlayers handles GET /players - fetch all players
func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
  players := h.service.GetAllPlayers()
  
  response := Response{
    Status:  "success",
    Message: "Players fetched successfully",
    Data:    players,
  }
  
  log.Printf("GET /players - returned %d players", len(players))
  h.sendJSONResponse(w, http.StatusOK, response)
}

// GetPlayer handles GET /players/{id} - fetch a single player
func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  if id == "" {
    h.sendErrorResponse(w, http.StatusBadRequest, "Player ID is required", errors.New("missing player ID"))
    return
  }
  
  player, err := h.service.GetPlayerByID(id)
  if err != nil {
    if errors.Is(err, ErrPlayerNotFound) {
      h.sendErrorResponse(w, http.StatusNotFound, "Player not found", err)
      return
    }
    h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to get player", err)
    return
  }
  
  response := Response{
    Status:  "success",
    Message: "Player fetched successfully",
    Data:    player,
  }
  
  log.Printf("GET /players/%s - returned player: %s", id, player.Name)
  h.sendJSONResponse(w, http.StatusOK, response)
}

// CreatePlayer handles POST /players - create a new player
func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
  var req PlayerRequest
  
  // Parse JSON request body
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", err)
    return
  }
  
  // Create the player
  player, err := h.service.CreatePlayer(req)
  if err != nil {
    if errors.Is(err, ErrInvalidInput) {
      h.sendErrorResponse(w, http.StatusBadRequest, "Invalid input", err)
      return
    }
    if errors.Is(err, ErrPlayerExists) {
      h.sendErrorResponse(w, http.StatusConflict, "Player already exists", err)
      return
    }
    h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to create player", err)
    return
  }
  
  response := Response{
    Status:  "success",
    Message: "Player created successfully",
    Data:    player,
  }
  
  log.Printf("POST /players - created player: %s (ID: %s)", player.Name, player.ID)
  h.sendJSONResponse(w, http.StatusCreated, response)
}

// UpdatePlayer handles PUT /players/{id} - update an existing player
func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  if id == "" {
    h.sendErrorResponse(w, http.StatusBadRequest, "Player ID is required", errors.New("missing player ID"))
    return
  }
  
  var req PlayerRequest
  
  // Parse JSON request body
  if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
    h.sendErrorResponse(w, http.StatusBadRequest, "Invalid JSON format", err)
    return
  }
  
  // Update the player
  player, err := h.service.UpdatePlayer(id, req)
  if err != nil {
    if errors.Is(err, ErrPlayerNotFound) {
      h.sendErrorResponse(w, http.StatusNotFound, "Player not found", err)
      return
    }
    if errors.Is(err, ErrInvalidInput) {
      h.sendErrorResponse(w, http.StatusBadRequest, "Invalid input", err)
      return
    }
    if errors.Is(err, ErrPlayerExists) {
      h.sendErrorResponse(w, http.StatusConflict, "Player conflict", err)
      return
    }
    h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to update player", err)
    return
  }
  
  response := Response{
    Status:  "success",
    Message: "Player updated successfully",
    Data:    player,
  }
  
  log.Printf("PUT /players/%s - updated player: %s", id, player.Name)
  h.sendJSONResponse(w, http.StatusOK, response)
}

// DeletePlayer handles DELETE /players/{id} - delete a player
func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
  id := r.PathValue("id")
  if id == "" {
    h.sendErrorResponse(w, http.StatusBadRequest, "Player ID is required", errors.New("missing player ID"))
    return
  }
  
  player, err := h.service.DeletePlayer(id)
  if err != nil {
    if errors.Is(err, ErrPlayerNotFound) {
      h.sendErrorResponse(w, http.StatusNotFound, "Player not found", err)
      return
    }
    h.sendErrorResponse(w, http.StatusInternalServerError, "Failed to delete player", err)
    return
  }
  
  response := Response{
    Status:  "success",
    Message: "Player deleted successfully",
    Data:    player,
  }
  
  log.Printf("DELETE /players/%s - deleted player: %s", id, player.Name)
  h.sendJSONResponse(w, http.StatusOK, response)
}

// Legacy handlers for backward compatibility (keeping the original function signatures)
// These use the global service instance

var globalPlayerService *PlayerService

func init() {
  globalPlayerService = NewPlayerService()
}

// GetPlayers is the legacy handler for backward compatibility
func GetPlayers(w http.ResponseWriter, r *http.Request) {
  handler := NewPlayerHandler(globalPlayerService)
  handler.GetPlayers(w, r)
}

// CreatePlayer is the legacy handler for backward compatibility
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
  handler := NewPlayerHandler(globalPlayerService)
  handler.CreatePlayer(w, r)
}

// UpdatePlayer is the legacy handler for backward compatibility
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
  handler := NewPlayerHandler(globalPlayerService)
  handler.UpdatePlayer(w, r)
}

// DeletePlayer is the legacy handler for backward compatibility
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
  handler := NewPlayerHandler(globalPlayerService)
  handler.DeletePlayer(w, r)
}