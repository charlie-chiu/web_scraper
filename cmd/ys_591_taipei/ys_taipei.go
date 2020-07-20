package main

import (
	"fmt"
	"log"
	"time"

	scraper "web_scraper"
)

func main() {
	startTime := time.Now()
	q := scraper.QueryTaipei

	s := scraper.NewFiveN1()
	rentals := s.ScrapeRentals(q)
	s.ScrapeRentalsDetail(rentals)

	region := "台北"
	date := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s-%s", region, date)
	rentals.ReplaceSection()
	rentals.Print()
	_ = rentals.SaveAsXLSX(filename + ".xlsx")

	log.Printf("execution time %s", time.Since(startTime))
}
