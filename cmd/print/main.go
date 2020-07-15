package main

import (
	"log"
	"time"

	scraper "web_scraper"
)

func main() {
	startTime := time.Now()
	//scraper.PrintAreaList()
	//scraper.PrintSectionDict()
	q := scraper.QueryMini
	//q := scraper.QueryTaiChung

	s := scraper.NewFiveN1()
	rentals := s.ScrapeList(q)

	filename := time.Now().Format("2006-01-02")
	rentals.ReplaceSection()
	rentals.Print()
	_ = rentals.SaveAsJSON(filename + ".json")
	_ = rentals.SaveAsXLSX(filename + ".xlsx")

	log.Printf("execution time %s", time.Since(startTime))
}
