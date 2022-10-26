package main

import (
	"math/rand"
	"net/http"
)

type Event struct {
	Name, Description, Author string
	Picked                    bool
}

var events = []Event{}

func eventRandomizer(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "eventpicker.html", &events)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func submitEvent(w http.ResponseWriter, r *http.Request) {
	for i := range events {
		events[i].Picked = false
	}
	events = append(events, Event{Name: r.FormValue("Name"), Description: r.FormValue("Description"), Author: r.FormValue("Author")})
	events[rand.Intn(len(events))].Picked = true
	http.Redirect(w, r, "/events/", http.StatusFound)
}
