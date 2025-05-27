package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlayerService_CreatePlayer(t *testing.T) {
	service := NewPlayerService()
	
	tests := []struct {
		name        string
		request     PlayerRequest
		wantError   bool
		errorType   error
	}{
		{
			name: "valid player",
			request: PlayerRequest{
				Name:         "Test Player",
				JerseyNumber: 15,
				Rating:       85,
			},
			wantError: false,
		},
		{
			name: "invalid name",
			request: PlayerRequest{
				Name:         "",
				JerseyNumber: 15,
				Rating:       85,
			},
			wantError: true,
			errorType: ErrInvalidInput,
		},
		{
			name: "invalid jersey number",
			request: PlayerRequest{
				Name:         "Test Player",
				JerseyNumber: 0,
				Rating:       85,
			},
			wantError: true,
			errorType: ErrInvalidInput,
		},
		{
			name: "invalid rating",
			request: PlayerRequest{
				Name:         "Test Player",
				JerseyNumber: 15,
				Rating:       0,
			},
			wantError: true,
			errorType: ErrInvalidInput,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.CreatePlayer(tt.request)
			
			if tt.wantError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
				if tt.errorType != nil && !errors.Is(err, tt.errorType) {
					t.Errorf("Expected error type %v, got %v", tt.errorType, err)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got: %v", err)
				}
			}
		})
	}
}

func TestPlayerHandler_GetPlayers(t *testing.T) {
	service := NewPlayerService()
	handler := NewPlayerHandler(service)
	
	req := httptest.NewRequest("GET", "/players", nil)
	w := httptest.NewRecorder()
	
	handler.GetPlayers(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	
	var response Response
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	
	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}
	
	// Should have initial sample data (3 players)
	players, ok := response.Data.([]interface{})
	if !ok {
		t.Fatalf("Expected data to be array of players")
	}
	
	if len(players) != 3 {
		t.Errorf("Expected 3 players, got %d", len(players))
	}
}

func TestPlayerHandler_CreatePlayer(t *testing.T) {
	service := NewPlayerService()
	handler := NewPlayerHandler(service)
	
	tests := []struct {
		name           string
		request        PlayerRequest
		expectedStatus int
	}{
		{
			name: "valid player",
			request: PlayerRequest{
				Name:         "Test Player",
				JerseyNumber: 15,
				Rating:       85,
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "invalid player - empty name",
			request: PlayerRequest{
				Name:         "",
				JerseyNumber: 15,
				Rating:       85,
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "duplicate player",
			request: PlayerRequest{
				Name:         "Messi",
				JerseyNumber: 10,
				Rating:       99,
			},
			expectedStatus: http.StatusConflict,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.request)
			req := httptest.NewRequest("POST", "/players", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			
			handler.CreatePlayer(w, req)
			
			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

func TestPlayerHandler_GetPlayer(t *testing.T) {
	service := NewPlayerService()
	handler := NewPlayerHandler(service)
	
	// Test getting existing player
	req := httptest.NewRequest("GET", "/players/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	
	handler.GetPlayer(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	
	// Test getting non-existent player
	req = httptest.NewRequest("GET", "/players/999", nil)
	req.SetPathValue("id", "999")
	w = httptest.NewRecorder()
	
	handler.GetPlayer(w, req)
	
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestPlayerHandler_UpdatePlayer(t *testing.T) {
	service := NewPlayerService()
	handler := NewPlayerHandler(service)
	
	updateRequest := PlayerRequest{
		Name:         "Updated Messi",
		JerseyNumber: 10,
		Rating:       97,
	}
	
	body, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest("PUT", "/players/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	
	handler.UpdatePlayer(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	
	var response Response
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	
	if response.Status != "success" {
		t.Errorf("Expected status 'success', got '%s'", response.Status)
	}
}

func TestPlayerHandler_DeletePlayer(t *testing.T) {
	service := NewPlayerService()
	handler := NewPlayerHandler(service)
	
	// Test deleting existing player
	req := httptest.NewRequest("DELETE", "/players/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	
	handler.DeletePlayer(w, req)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
	
	// Test deleting non-existent player
	req = httptest.NewRequest("DELETE", "/players/999", nil)
	req.SetPathValue("id", "999")
	w = httptest.NewRecorder()
	
	handler.DeletePlayer(w, req)
	
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

// Helper function to check error types (simple implementation)
func ErrorIs(err, target error) bool {
	return err != nil && target != nil && err.Error() == target.Error()
} 