package scraper

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/vinta/pangu"
)

type RentScraper interface {

	// ScrapeList scrape rent.591.com.tw
	ScrapeList(query Query) Rentals

	// ScrapePhone scrape rent.591.com.tw/rent-detail-{id}.html
	// modify this to Scrape detail if needed
	ScrapePhone(query Query) string
}

type FiveN1 struct {
	rentals  Rentals
	queryURL string

	records int
	pages   int

	wg     sync.WaitGroup
	rw     sync.RWMutex
	client *http.Client
	cookie *http.Cookie
}

func NewFiveN1() *FiveN1 {
	// default with Taipei
	defaultCookie := &http.Cookie{
		Name:  "urlJumpIp",
		Value: "1",
	}
	return &FiveN1{
		cookie: defaultCookie,
		client: &http.Client{},
	}
}

func (f *FiveN1) ScrapeList(query Query) Rentals {
	f.queryURL, _ = query.URL()

	//parse
	f.parseFirstPage()
	f.showQueryInfo()

	for page := 0; page < f.pages; page++ {
		f.wg.Add(1)
		go f.scrapeWorker(page)
	}

	f.wg.Wait()

	return f.rentals
}

func (f *FiveN1) parseFirstPage() {
	d := NewDocument()
	res := f.request(f.queryURL)
	d.clone(res)

	f.parseRecordsNum(d.doc) // Record pages number at first
}

func (f *FiveN1) request(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.AddCookie(f.cookie)

	res, err := f.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (f *FiveN1) parseRecordsNum(doc *goquery.Document) {
	doc.Find(".pull-left.hasData > i").Each(func(_ int, selector *goquery.Selection) {
		recordString := stringReplacer(selector.Text())
		replaceComma := strings.Replace(recordString, ",", "", -1)
		totalRecord, _ := strconv.Atoi(replaceComma)
		pages := totalRecord / itemsPerPage

		if totalRecord%itemsPerPage > 0 {
			pages += 1
		}

		f.records = totalRecord
		f.pages = pages
	})
}

func (f *FiveN1) scrapeWorker(page int) {
	defer f.wg.Done()

	d := NewDocument()
	firstRow := strconv.Itoa(page * itemsPerPage)
	res := f.request(f.queryURL + "&firstRow=" + firstRow)
	d.clone(res)

	f.parseRentHouse(page, d.doc)
}

func (f *FiveN1) parseRentHouse(page int, doc *goquery.Document) {
	doc.Find("#content").Each(func(_ int, selector *goquery.Selection) {
		selector.Find(".listInfo.clearfix").Each(func(item int, listInfo *goquery.Selection) {
			rental := NewRental()

			// Content Title
			title := listInfo.Find(".pull-left.infoContent > h3 > a[href]").Text()
			rental.Title = stringReplacer(title)

			// Content URL
			var url string
			if href, ok := listInfo.Find(".pull-left.infoContent > h3 > a").Attr("href"); ok {
				url = stringReplacer(href)
			}
			rental.URL = "https:" + url

			if ID, ok := listInfo.Find(".pull-left.infoContent > span > a").Attr("data-text"); ok {
				rental.ID = "R" + ID
			}

			//if crop, ok := listInfo.Find(".pull-left.imageBox > img").Attr("data-original"); ok {
			//	preview := strings.Replace(crop, "210x158.crop.jpg", "765x517.water3.jpg", 1)
			//	rental.Preview = preview
			//}

			listInfo.Find(".pull-left.infoContent").Each(func(_ int, infoContent *goquery.Selection) {
				// Rent House Description.
				description := stringReplacer(infoContent.Find(".lightBox").First().Text())

				splitDescription := strings.Split(description, "|")

				// Exchange
				if len(splitDescription) == 4 {
					tmp := splitDescription[2] // 坪數
					splitDescription[2] = splitDescription[1]
					splitDescription[1] = tmp
				}

				if len(splitDescription) < 4 {
					splitDescription = fillDescription(splitDescription)
				}

				rental.OptionType = trimTextSpace(splitDescription[0])
				rental.Ping = trimTextSpace(splitDescription[1])
				rental.RentType = trimTextSpace(splitDescription[2])
				rental.Floor = trimTextSpace(splitDescription[3])

				// Rent House Address
				address := stringReplacer(infoContent.Find(".lightBox").Eq(1).Text())
				rental.Address = address
			})

			// Rent Price
			listInfo.Find(".price").Each(func(_ int, price *goquery.Selection) {
				rental.Price = stringReplacer(price.Text())
			})

			// New Rent House
			//listInfo.Find(".newArticle").Each(func(_ int, n *goquery.Selection) {
			//	rental.IsNew = true
			//})

			// Add rent house into list
			f.rw.Lock()
			f.rentals = append(f.rentals, *rental)
			f.rw.Unlock()
		})
	})
}

func (f *FiveN1) ScrapePhone(url string) string {
	panic("implement me")
}

func (f *FiveN1) SetReqCookie(region string) {
	f.cookie.Value = region
}

func (f *FiveN1) showQueryInfo() {
	log.Printf("# Total Page: %3d | Total Record: %d\n", f.pages, f.records)
	log.Printf("# Query URL: %s\n", f.queryURL)
}

func stringReplacer(text string) string {
	replacer := strings.NewReplacer("\n", "", " ", "")

	return pangu.SpacingText(replacer.Replace(text))
}

func trimTextSpace(s string) string {
	return strings.Fields(s)[0]
}

func fillDescription(s []string) []string {
	s = append(s, s[2])
	s[2] = "沒有格局說明"

	return s
}

//todo: understand why clone document here, remove it if unnecessary
type Document struct {
	doc *goquery.Document
}

func NewDocument() *Document {
	return &Document{
		doc: &goquery.Document{},
	}
}

func (d *Document) clone(res *http.Response) {
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	d.doc = doc
}
