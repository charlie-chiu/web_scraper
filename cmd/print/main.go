package main

import scraper "web_scraper"

func main() {
	//scraper.PrintAreaList()
	//scraper.PrintSectionDict()
	q := scraper.QueryMini
	//q := scraper.QueryTaiChung

	s := scraper.NewFiveN1()
	rentals := s.ScrapeList(q)

	rentals.ReplaceSection()
	rentals.Print()
}
