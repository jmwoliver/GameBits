package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetSearch",
		"GET",
		"/search/{query}",
		GetSearch,
	},
	Route{
		"GetAlbum",
		"GET",
		"/album/{albumTitle}",
		GetAlbum,
	},
	Route{
		"GetDownloadTrackLink",
		"GET",
		"/track/{href}",
		GetDownloadTrackLink,
	},
	Route{
		"GetLetterSearch",
		"GET",
		"/letter/{href}",
		GetLetterSearch,
	},
	Route{
		"GetConsoleSearch",
		"GET",
		"/console/{console}",
		GetConsoleSearch,
	},
}
