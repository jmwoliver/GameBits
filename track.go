package main

type Track struct {
	Id        	string      `json:"id"`
	TrackNumber	int			`json:"track_number"`
	Title      	string    	`json:"title"`
	URL		 	string    	`json:"url"`
	Duration	string		`json:"duration"`
}

type Tracks []Track
