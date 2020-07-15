package main

import (
	scraper "web_scraper"
)

func main() {
	q := scraper.QueryMini

	s := scraper.NewFiveN1()
	rentals := s.ScrapeList(q)

	rentals.Print()
}
