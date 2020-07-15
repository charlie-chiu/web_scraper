package scraper

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
	ID      string `json:"id"`      //出現於 url最後的識別，應該也是591內部的編號 R{id} for Rent{id}?
	Phone   string `json:"phone"`   //聯絡電話
	Section string `json:"section"` //行政區
}

func NewRental() *Rental {
	return &Rental{}
}

type Rentals []Rental
