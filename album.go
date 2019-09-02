package main

type Album struct {
	Id        	string    	`json:"id"`
	Title      	string    	`json:"title"`
	URL		 	string    	`json:"url"`
	Cover		string		`json:"cover"`
}

type Albums []Album
