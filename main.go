package main

import (
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Location struct {
	Lat float64
	Lon float64
}

func rLat() float64 {
	max := 50.175444
	min := 50.104591
	return rand.Float64()*(max-min) + min
}

func rLon() float64 {
	max := 8.753328
	min := 8.525362
	return rand.Float64()*(max-min) + min
}

func echoHandler(ws *websocket.Conn) {
	//loc := Location{`{"lat":"50.07", "lon":"8.14"}`}
	// 50.119, 8.68211
	// Left up: 50.175444, 8.525362
	// Right down: 50.104591, 8.753328
	for {
		randLat := rLat()
		randLon := rLon()
		loc := Location{Lat: randLat, Lon: randLon}
		websocket.JSON.Send(ws, loc)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
