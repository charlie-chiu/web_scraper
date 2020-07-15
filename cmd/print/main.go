package main

import (
	"strconv"

	scraper "web_scraper"
)

func main() {
	q := scraper.QueryMini

	s := scraper.NewFiveN1()
	s.SetReqCookie(strconv.Itoa(q.Region))
	rentals := s.ScrapeList(q)

	rentals.Print()
}
