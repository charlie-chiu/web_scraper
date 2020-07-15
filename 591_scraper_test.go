package scraper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func item120Handler(w http.ResponseWriter, r *http.Request) {
	testFixture := "test_fixture/591with120items.html"
	html, _ := ioutil.ReadFile(testFixture)
	w.Header().Set("Content-Type", "application/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(html)
}

func TestFiveN1_ScrapeList(t *testing.T) {
	t.Run("scrape with default cookie", func(t *testing.T) {
		wantValue := "1"
		cookieName := "urlJumpIp"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotCookie, err := r.Cookie(cookieName)
			assert.Nil(t, err)
			assert.Equal(t, wantValue, gotCookie.Value)
		}))
		defer server.Close()
		query := Query{
			RootURL: server.URL + "/?",
		}
		scraper := NewFiveN1()

		_ = scraper.ScrapeList(query)
	})

	// todo: shouldn't do this by package user
	t.Run("scrape with custom cookie", func(t *testing.T) {
		wantValue := "20"
		cookieName := "urlJumpIp"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotCookie, err := r.Cookie(cookieName)
			assert.Nil(t, err)
			assert.Equal(t, wantValue, gotCookie.Value)
		}))

		defer server.Close()
		query := Query{
			RootURL: server.URL + "/?",
		}

		scraper := NewFiveN1()
		scraper.SetReqCookie("20")

		_ = scraper.ScrapeList(query)
	})

	t.Run("scrape url with 120 items", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(item120Handler))
		defer server.Close()
		query := Query{
			RootURL: server.URL + "/?",
		}

		scraper := NewFiveN1()
		rentals := scraper.ScrapeList(query)

		assert.Equal(t, 120, scraper.records)
		assert.Equal(t, 4, scraper.pages)
		assert.Equal(t, 120, len(rentals))
	})

	t.Run("scrape url with 333 items", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testFixture := "test_fixture/591with333items.html"
			html, _ := ioutil.ReadFile(testFixture)
			w.Header().Set("Content-Type", "application/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(html)
		}))
		defer server.Close()
		query := Query{
			RootURL: server.URL + "/?",
		}

		scraper := NewFiveN1()

		_ = scraper.ScrapeList(query)

		assert.Equal(t, 333, scraper.records)
		assert.Equal(t, 12, scraper.pages)

		// this test will fail because test fixture always have 30 item per page.
		//assert.Equal(t, 333, len(rentals))
	})

	t.Run("scrape group by section and store section in Rental", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//requestCount++
			html, _ := ioutil.ReadFile("test_fixture/591with2items.html")
			w.Header().Set("Content-Type", "application/html")
			_, _ = w.Write(html)
		}))
		defer server.Close()

		query := Query{
			RootURL: server.URL + "/?",
			Section: "98,99,100",
		}

		f := NewFiveN1()
		gotRentals := f.ScrapeList(query)

		wantRentals := Rentals{
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				RentType:   "6",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				Phone:      "",
				Section:    "98",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				RentType:   "1",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				Phone:      "",
				Section:    "98",
			},
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				RentType:   "6",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				Phone:      "",
				Section:    "99",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				RentType:   "1",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				Phone:      "",
				Section:    "99",
			},
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				RentType:   "6",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				Phone:      "",
				Section:    "100",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				RentType:   "1",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				Phone:      "",
				Section:    "100",
			},
		}

		for i, wantRental := range wantRentals {
			if i >= len(gotRentals) {
				t.Fatalf("%dth wantRental not exsits", i)
			}
			gotRental := gotRentals[i]
			if !reflect.DeepEqual(wantRental, gotRental) {
				t.Errorf("%dth wantRental wantSection not equal, \nwant %v,\n got %v", i, wantRental, gotRental)
			}
		}
	})
}
