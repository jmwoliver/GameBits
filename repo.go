package main

import (
	"bytes"
	"strings"
	"strconv"
	"encoding/base64"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"
)

var currentAlbumId int
var currentTrackId int

var albums Albums
var tracks Tracks

func RepoFindAlbums(query string) (Albums, error) {
	albums = nil
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	u := "https://downloads.khinsider.com/search?search=" + query


	req.SetRequestURI(u)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return nil, err
	}

	doc.Find("#EchoTopic > p:nth-child(3) > a").Each(func(i int, s *goquery.Selection) {
		// ENCODE AND DECODE URL
		url, exist := s.Attr("href")
		title := s.Text()
		encodedUrl := base64.StdEncoding.EncodeToString([]byte(url))
		// go RepoAlbumArt(title, encodedUrl)

		if exist {
			RepoCreateAlbum(Album{Id: encodedUrl, Title: title, URL: url})
		}
	})

	return albums, nil
}

func RepoCreateAlbum(a Album) Album {
	albums = append(albums, a)
	return a
}

func RepoGetLetterSearch(query string) (Albums, error) {
	albums = nil
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	u := BASE_URL + "/game-soundtracks/browse/" + query


	req.SetRequestURI(u)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return nil, err
	}

	doc.Find("#EchoTopic > p:nth-child(5) > a").Each(func(i int, s *goquery.Selection) {
		// ENCODE AND DECODE URL
		href, exist := s.Attr("href")
		title := s.Text()
		url := BASE_URL + href
		encodedUrl := base64.StdEncoding.EncodeToString([]byte(url))
		// go RepoAlbumArt(title, encodedUrl)

		if exist {
			RepoCreateAlbum(Album{Id: encodedUrl, Title: title, URL: url})
		}
	})

	return albums, nil
}

func RepoGetConsoleSearch(query string) (Albums, error) {
	albums = nil
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	u := BASE_URL + "/game-soundtracks/" + query
	fmt.Println(u)


	req.SetRequestURI(u)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return nil, err
	}

	doc.Find("#EchoTopic > p:nth-child(5) > a").Each(func(i int, s *goquery.Selection) {
		// ENCODE AND DECODE URL
		href, exist := s.Attr("href")
		title := s.Text()
		url := href
		encodedUrl := base64.StdEncoding.EncodeToString([]byte(url))
		// go RepoAlbumArt(title, encodedUrl)

		if exist {
			RepoCreateAlbum(Album{Id: encodedUrl, Title: title, URL: url})
		}
	})

	return albums, nil
}

// TODO: add functionality to go retreive album image with Bing Search API
// func RepoAlbumArt(title string, id string) {
// 	req := fasthttp.AcquireRequest()
// 	res := fasthttp.AcquireResponse()

// 	// Base64 Url Decoding
// 	decode, _ := base64.URLEncoding.DecodeString(encodedUrl)
// 	url := string(decode)

// 	req.SetRequestURI(url)

// 	if err := fasthttp.Do(req, res); err != nil {
// 		return nil, err
// 	}

// 	// Load the HTML document
// 	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
// 	if err != nil {
// 		return nil, err
// 	}
// 	doc.Find("#songlist_header > th:nth-child(2) > b").Eq(0).Text()

// }

func NewTitle(title string) string {
	r := strings.NewReplacer(
		" The ", " the ",
        " An ", " an ",
		" A ", " a ",
		" On ", " on ",
        " Of ", " of ",
		" In ", " in ",
		" With ", " with ",
        " By ", " by ",
		" To ", " to ",
		" At ", " at ",)

    result := r.Replace(strings.Title(title))
	return result
}

func RepoFindTracks(encodedUrl string) (Tracks, error) {
	tracks = nil
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	// Base64 Url Decoding
	decode, _ := base64.URLEncoding.DecodeString(encodedUrl)
	url := string(decode)

	req.SetRequestURI(url)

	if err := fasthttp.Do(req, res); err != nil {
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return nil, err
	}
	var index int

	//checking if there is a column entitled "Track"
	//if there is, then 3 is the correct index for retrieving title and duration
	//if not, 2 is correct
	if doc.Find("#songlist_header > th:nth-child(2) > b").Eq(0).Text() == "Track" {
		index = 3
	} else {
		index = 2
	}

	var durations []string

	doc.Find("#songlist > tbody > tr > td:nth-child("+strconv.Itoa(index+1)+") > a").Each(func(i int, s *goquery.Selection) {
		_, exist := s.Attr("href")
		duration := s.Text()

		if exist {
			durations = append(durations, duration)
		}
	})
	currentTrackId = 0
	doc.Find("#songlist > tbody > tr > td:nth-child("+strconv.Itoa(index)+") > a").Each(func(i int, s *goquery.Selection) {
		href, exist := s.Attr("href")
		title := NewTitle(s.Text())
		duration := durations[i]
		url := BASE_URL + href
		encodedUrl := base64.StdEncoding.EncodeToString([]byte(url))

		if exist {
			RepoCreateTrack(Track{Id: encodedUrl, Title: title, URL: url, Duration: duration})
		}
	})

	return tracks, nil
}

func RepoCreateTrack(t Track) Track {
	currentTrackId += 1
	t.TrackNumber = currentTrackId
	tracks = append(tracks, t)
	return t
}

//broken, doesn't like the slashes in url
func RepoGetDownloadTrackLink(encodedUrl string) (Link, error) {

	// Base64 Url Decoding
	url, _ := base64.URLEncoding.DecodeString(encodedUrl)
	fmt.Println(string(url))
	

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	req.SetRequestURI(string(url))

	if err := fasthttp.Do(req, res); err != nil {
		return Link{}, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		return Link{}, err
	}

	download, exist := doc.Find("#EchoTopic > p > a[href*='/ost/']").Eq(0).Attr("href")
	if !exist {
		return Link{}, err
	}

	link := Link{Download: download}
	return link, nil
}