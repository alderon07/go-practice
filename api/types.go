package main

type Response struct {
    Message string   `json:"message"`
    Status  string   `json:"status"`
    Data    []Player `json:"data"`
}

type Player struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    JerseyNumber int8   `json:"jersey_number"`
    Rating       int8   `json:"rating"`
}

func (p *Player) Update(player Player) {
    p.JerseyNumber = player.JerseyNumber
    p.Rating = player.Rating
}
