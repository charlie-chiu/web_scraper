package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Rental represent a rental house
// commented field mean we don't need it now
type Rental struct {
	//Preview    string `json:"preview"` // preview image
	Title      string `json:"title"`
	URL        string `json:"url"`
	Address    string `json:"address"`
	RentType   string `json:"rentType"`   // 以代號儲存的格局，不轉換的話沒有用
	OptionType string `json:"optionType"` // 獨立套房、整層住家… etc
	Ping       string `json:"ping"`       // 坪數
	Floor      string `json:"floor"`      //樓層
	Price      string `json:"price"`      // 租金
	//IsNew      bool   `json:"isNew"`
	ID      string `json:"id"` //出現於 url最後的識別，應該也是591內部的編號 R{id} for Rent{id}?
	PostBy  string `json:"-"`
	Phone   string `json:"-"`       //聯絡電話
	Section string `json:"section"` //行政區
}

func NewRental() *Rental {
	return &Rental{}
}

type Rentals []Rental

func (r *Rentals) Print() {
	for i, rental := range *r {
		log.Printf("%4d.|%s|%s|%s|%s|%s\n", i, rental.Section, rental.OptionType, rental.Price, rental.Title, rental.URL)
	}
}

// ReplaceSection replace all section code with section name
func (r *Rentals) ReplaceSection() {
	for i, rental := range *r {
		sectionCode := rental.Section
		(*r)[i].Section = sectionDict[sectionCode]
	}
}

func (r Rentals) SaveAsJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("create %s error %v", filename, err)
	}

	err = json.NewEncoder(file).Encode(r)
	if err != nil {
		return fmt.Errorf("json encode error %v", err)
	}

	return nil
}

func (r Rentals) SaveAsXLSX(filename string) error {
	x := newXlsx()
	err := x.WriteNextRow("區", "標題", "類型", "租金", "聯絡人", "電話", "連結")
	if err != nil {
		return fmt.Errorf("xlsx.WriteNextRow error %v", err)
	}
	for _, rental := range r {
		err := x.WriteNextRow(rental.Section, rental.Title, rental.OptionType, rental.Price, rental.PostBy, rental.Phone, rental.URL)
		if err != nil {
			return fmt.Errorf("xlsx.WriteNextRow error %v", err)
		}
	}

	err = x.Save(filename)
	if err != nil {
		return fmt.Errorf("xlsx save file error %v", err)
	}

	return nil
}
