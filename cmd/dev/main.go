package main

import (
	"log"
	"time"

	scraper "web_scraper"
)

func main() {
	startTime := time.Now()
	q := scraper.QueryMini

	s := scraper.NewFiveN1()
	rentals := s.ScrapeRentals(q)
	s.ScrapeRentalsDetail(rentals)

	filename := time.Now().Format("2006-01-02")
	rentals.ReplaceSection()
	rentals.Print()
	_ = rentals.SaveAsJSON(filename + ".json")
	_ = rentals.SaveAsXLSX(filename + ".xlsx")

	log.Printf("execution time %s", time.Since(startTime))
}
