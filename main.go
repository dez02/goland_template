package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Gamer struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type myChannels struct {
	nameChannel  chan string
	stateChannel chan string
}

var n string
var s string
var Gamers []Gamer

func manageData(channels myChannels) {
	for {
		fmt.Println("Waiting for a string")
		time.Sleep(2 * time.Second)
		fmt.Println("Sending a string")
		select { // c’est comme un while
		case n = <-channels.nameChannel:
			fmt.Println("Received this string", n)
		case s = <-channels.stateChannel:
			fmt.Println("Received this string", s)

		}
	}
}

func getAllGamers(w http.ResponseWriter, r *http.Request) {
		fmt.Println("returnGamers")
		json.NewEncoder(w).Encode(Gamers)
}

func post(channels myChannels) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		channels.nameChannel <- r.FormValue("name")
		channels.stateChannel <- r.FormValue("state")

		gamer := Gamer{
			Name:  n,
			State: s,
		}
		Gamers = append(Gamers, gamer)
		tmpl.Execute(w, nil)
		return
	}
}




func main() {

	var channels = myChannels{make(chan string), make(chan string)}
	go manageData(channels)

	http.HandleFunc("/", post(channels))
	http.HandleFunc("/gamers", getAllGamers)


	log.Fatal(http.ListenAndServe(":8082", nil))

}
