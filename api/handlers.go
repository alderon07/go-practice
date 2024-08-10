package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
)

// Sample data representing players
var data = map[string]Player{
    "1": {
        ID: "1",
        Name: "Messi",
        JerseyNumber: 10,
        Rating: 99,
    },
    "2": {
        ID: "2",
        Name: "Ronaldo",
        JerseyNumber: 7,
        Rating: 98,
    },
    "3": {
        ID: "3",
        Name: "Neymar",
        JerseyNumber: 10,
        Rating: 95,
    },
}

// GetPlayers handles the GET request to fetch all players
func GetPlayers(w http.ResponseWriter, r *http.Request) {
    response := Response{}
    players := []Player{}

    // Fetch all players from the data map
    for _, v := range data {
        players = append(players, v)
    }

    // Prepare the response
    response.Message = "players fetched successfully"
    response.Status = "success"
    response.Data = players
    
    // Convert response to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)

    // Log the response
    log.Printf("GET /players response: %v", string(jsonResponse))
}

// DeletePlayer handles the DELETE request to remove a player by ID
func DeletePlayer(w http.ResponseWriter, r *http.Request) {
    response := Response{}

    // Get the player ID from the request path
    id := r.PathValue("id")
    player := data[id]
    
    // Remove the player from the data map
    for key := range data {
        if key == id {
            delete(data, key)
            break
        }
    }

    // Prepare the response
    response.Message = "player deleted successfully"
    response.Status = "success"
    response.Data = []Player{player}
    
    // Convert response to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)

    // Log the response
    log.Printf("DELETE /players response: %#v", player)
}

// UpdatePlayer handles the PUT request to update a player's details
func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
    var response Response
    var player Player

    // Get the player ID from the request path
    id := r.PathValue("id")
    
    // Parse the jersey number and rating from the form values
    jerseyNumber, err := strconv.ParseInt(r.FormValue("jersey_number"), 10, 8)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    rating, err := strconv.ParseInt(r.FormValue("rating"), 10, 8)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create a Player struct with the updated details
    playerUpdateTo := Player{
        JerseyNumber: int8(jerseyNumber),
        Rating: int8(rating),
    }

    // Update the player in the data map
    player = data[id]
    player.Update(playerUpdateTo)
    data[id] = player

    // Prepare the response
    response.Message = "player updated successfully"
    response.Status = "success"
    response.Data = []Player{player}

    // Convert response to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)

    // Log the response
    log.Printf("PUT /players response: %#v", player)
}

// CreatePlayer handles the POST request to create a new player
func CreatePlayer(w http.ResponseWriter, r *http.Request) {
    var response Response
    var player Player
    
    // Parse the jersey number and rating from the form values
    jerseyNumber, err := strconv.ParseInt(r.FormValue("jersey_number"), 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    rating, err := strconv.ParseInt(r.FormValue("rating"), 10, 64)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create a new Player struct with the provided details
    player = Player{
        ID: strconv.Itoa(len(data) + 1),
        Name: r.FormValue("name"),
        JerseyNumber: int8(jerseyNumber),
        Rating: int8(rating),
    }

    // Add the new player to the data map
    data[player.ID] = player

    // Prepare the response
    response.Message = "player created successfully"
    response.Status = "success"
    response.Data = []Player{player}

    // Convert response to JSON
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set response headers and write the JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)

    // Log the response
    log.Printf("POST /players response: %#v", player)
}