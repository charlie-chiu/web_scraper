package main

import (
	"log"
	"strconv"

	scraper "web_scraper"
)

func main() {
	q := scraper.QueryMini

	url, err := q.URL()
	if err != nil {
		log.Fatalf("option.URL error %v", err)
	}

	s := scraper.NewFiveN1()
	s.SetReqCookie(strconv.Itoa(q.Region))
	rentals := s.ScrapeList(url)

	rentals.Print()
}
