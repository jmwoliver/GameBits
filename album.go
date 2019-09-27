package main

type Album struct {
	Id        	string    	`json:"id"`
	Title      	string    	`json:"title"`
	Duration	string		`json:"duration"`
	URL		 	string    	`json:"url"`
	Cover		string		`json:"cover"`
	Tracks		Tracks		`json:"tracks"`
}

type Albums []Album
