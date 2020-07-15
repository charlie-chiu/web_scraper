package scraper

import (
	"encoding/json"
	"log"
	"os"
)

type Area struct {
	City     string    `json:"city"`
	Sections []Section `json:"section"`
}

type Section struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var areas = readAreaList()
var sectionDict = generateSectionName()

func PrintAreaList() {
	for _, area := range areas {
		log.Printf("%+v\n", area)
	}
}

func PrintSectionDict() {
	log.Printf("%+v", sectionDict)
}

func readAreaList() (areas []Area) {
	const areaListFilename = "areas.json"
	file, err := os.Open(areaListFilename)
	if err != nil {
		log.Fatalf("open file %s error: %v", areaListFilename, err)
	}

	err = json.NewDecoder(file).Decode(&areas)
	if err != nil {
		log.Fatalf("decode %s error: %v", areaListFilename, err)
	}

	return areas
}

func generateSectionName() map[string]string {
	sectionMap := make(map[string]string)
	for _, area := range areas {
		for _, section := range area.Sections {
			sectionMap[section.Code] = section.Name
		}
	}

	return sectionMap
}
