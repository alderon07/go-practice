package main

import (
	"log"
	"net/http"
)

func main(){
    // make route handler
    router := http.NewServeMux()

    router.HandleFunc("GET /players", GetPlayers)
    router.HandleFunc("DELETE /players/{id}", DeletePlayer)
    router.HandleFunc("PUT /players/{id}", UpdatePlayer)
    router.HandleFunc("POST /players", CreatePlayer)

    // describe server config
    server := http.Server{
        Addr: ":8080",
        Handler: router,
    }

    server.ListenAndServe()
    log.Printf("Started server on %v", server.Addr)
}

