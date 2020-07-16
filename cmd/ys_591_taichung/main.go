package main

import (
	"fmt"
	"log"
	"time"

	scraper "web_scraper"
)

func main() {
	startTime := time.Now()
	q := scraper.QueryTaiChung

	s := scraper.NewFiveN1()
	rentals := s.ScrapeList(q)

	for i, rental := range rentals {
		_ = s.ScrapeDetail(&rental)
		log.Println("scraping", rental.URL, "...")
		rentals[i] = rental
		time.Sleep(100 * time.Millisecond)
	}

	region := "台中"
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s", region, date)
	rentals.ReplaceSection()
	rentals.Print()
	_ = rentals.SaveAsXLSX(filename + ".xlsx")

	log.Printf("execution time %s", time.Since(startTime))
}
