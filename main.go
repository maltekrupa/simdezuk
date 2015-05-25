package main

import (
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"time"

	gj "github.com/kpawlik/geojson"
	"golang.org/x/net/websocket"
)

var FIVEMETERS = 0.0000449

type Boid struct {
	Id      int
	Lon     gj.CoordType
	Lat     gj.CoordType
	history gj.Coordinates
}

func (b *Boid) MoveNorth() {
	// Move 5 meter to north
	latitude := float64(b.Lat)
	tmpLat := latitude + FIVEMETERS
	newCoordinates := gj.Coordinate{gj.Coord(tmpLat), b.Lon}
	b.MoveToCoord(newCoordinates)
}

func (b *Boid) MoveSouth() {
	// Move 5 meter to south
	latitude := float64(b.Lat)
	tmpLat := latitude - FIVEMETERS
	newCoordinates := gj.Coordinate{gj.Coord(tmpLat), b.Lon}
	b.MoveToCoord(newCoordinates)
}

func (b *Boid) MoveWest() {
	// Move 5 meter to west
	longitude := float64(b.Lon)
	tmpLon := longitude + FIVEMETERS/math.Cos(longitude)
	newCoordinates := gj.Coordinate{b.Lat, gj.Coord(tmpLon)}
	b.MoveToCoord(newCoordinates)
}

func (b *Boid) MoveEast() {
	// Move 5 meter to east
	longitude := float64(b.Lon)
	tmpLon := longitude - FIVEMETERS/math.Cos(longitude)
	newCoordinates := gj.Coordinate{b.Lat, gj.Coord(tmpLon)}
	b.MoveToCoord(newCoordinates)
}

func (b *Boid) MoveToCoord(newPos gj.Coordinate) {
	// Add current coordinates to history and move to the new ones.
	b.history = append(b.history, newPos)
	b.Lat = newPos[0]
	b.Lon = newPos[1]
}

func (b Boid) PrepareForFrontend() *gj.Feature {
	// Create a GeoJSON feature for the frontend
	props := map[string]interface{}{"id": b.Id, "history": b.history}
	c := gj.Coordinate{b.Lon, b.Lat}
	p := gj.NewPoint(c)
	f := gj.NewFeature(p, props, nil)
	return f
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

func randomHandler(ws *websocket.Conn) {
	// 50.119, 8.68211
	// Left up: 50.175444, 8.525362
	// Right down: 50.104591, 8.753328
	b1 := Boid{Lat: 50.11021730200894, Lon: 8.750400943993286}
	for {
		randMove := rand.Intn(4)
		switch {
		case randMove == 0:
			b1.MoveNorth()
		case randMove == 1:
			b1.MoveSouth()
		case randMove == 2:
			b1.MoveWest()
		case randMove == 3:
			b1.MoveEast()
		}
		f := b1.PrepareForFrontend()

		websocket.JSON.Send(ws, f)

		time.Sleep(500 * time.Millisecond)
	}
}

func tmplHandler(w http.ResponseWriter, r *http.Request) {
	// parse templates
	templates := template.Must(template.ParseGlob("templates/*.tmpl"))
	templates.ExecuteTemplate(w, "indexPage", nil)
}

func main() {
	http.Handle("/ws", websocket.Handler(randomHandler))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", tmplHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
