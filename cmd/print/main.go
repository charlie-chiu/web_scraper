package main

import (
	"log"
	"strconv"

	scraper "web_scraper"
)

func main() {
	//o := scraper.NewOptions()
	//o.Region = 8
	//o.Section = "98"
	o := scraper.TestOption

	url, err := o.URL()
	if err != nil {
		log.Fatalf("option.URL error %v", err)
	}

	s := scraper.NewFiveN1()
	s.SetReqCookie(strconv.Itoa(o.Region))
	rentals := s.ScrapeList(url)

	rentals.Print()
}
