package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Search away!\n")
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var query string = vars["query"]
	var err error
	albums, err := RepoFindAlbums(query)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(albums); err != nil {
		panic(err)
	}
}

func GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var query string = vars["albumTitle"]
	var err error
	tracks, err := RepoFindTracks(query)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(tracks); err != nil {
		panic(err)
	}
}

func GetLetterSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var query string = vars["letter"]
	var err error
	letterAlbums, err := RepoGetLetterSearch(query)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(letterAlbums); err != nil {
		panic(err)
	}
}

func GetConsoleSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var query string = vars["console"]
	var err error
	consoleAlbums, err := RepoGetConsoleSearch(query)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(consoleAlbums); err != nil {
		panic(err)
	}
}


func GetDownloadTrackLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var query string = vars["href"]
	var err error
	link, err := RepoGetDownloadTrackLink(query)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(link); err != nil {
		panic(err)
	}
}
