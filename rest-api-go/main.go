// See https://tutorialedge.net/software-eng/designing-a-rest-api/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/stianeikeland/go-rpio"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/trains", returnAllTrains)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllTrains(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Trains)
}

type Train struct {
	Id    int
	Title string
}

var Trains []Train

func main() {
	Trains = []Train{
		Train{Id: 1, Title: "Güterzug"},
	}

	handleRequests()
}
