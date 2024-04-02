package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
)


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


func GetPlayers(w http.ResponseWriter, r *http.Request){
    response := Response{}
    players := []Player{}

    // database queries
    for _, v := range data {
        players = append(players, v)
    }

    response.Message = "players fetched successfully"
    response.Status = "success"
    response.Data = players
    
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)

    log.Printf("GET /players response: %v", string(jsonResponse))
}


func DeletePlayer(w http.ResponseWriter, r *http.Request){
    response := Response{}

    id := r.PathValue("id")
    player := data[id]
    
    // database manipulation
    for key := range data {
        if key == id {
            delete(data, key)
            break
        }
    }

    response.Message = "player deleted successfully"
    response.Status = "success"
    response.Data = []Player{player}
    
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
    log.Printf("DELETE /players response: %#v", player)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request){
    var response Response

    var player Player

    id := r.PathValue("id")
    
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

    playerUpdateTo := Player{
        JerseyNumber: int8(jerseyNumber),
        Rating: int8(rating),
    }


    player = data[id]
    player.Update(playerUpdateTo)
    data[id] = player

    response.Message = "player updated successfully"
    response.Status = "success"
    response.Data = []Player{player}

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
    log.Printf("PUT /players response: %#v", player)
}

func CreatePlayer(w http.ResponseWriter, r *http.Request){
    var response Response

    var player Player
    
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

    player = Player{
        ID: strconv.Itoa(len(data) + 1),
        Name: r.FormValue("name"),
        JerseyNumber: int8(jerseyNumber),
        Rating: int8(rating),
    }

    data[player.ID] = player

    response.Message = "player created successfully"
    response.Status = "success"
    response.Data = []Player{player}

    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
    log.Printf("POST /players response: %#v", player)
}