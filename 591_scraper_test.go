package scraper

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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
	t.Run("scrape with cookie urlJumpIp=region", func(t *testing.T) {
		wantRegion := 20
		regionStr := "20"
		cookieName := "urlJumpIp"

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotCookie, err := r.Cookie(cookieName)
			assert.Nil(t, err)
			assert.Equal(t, regionStr, gotCookie.Value)
		}))

		defer server.Close()
		query := &Query{
			RootURL: server.URL + "/?",
			Region:  wantRegion,
		}

		scraper := NewFiveN1()

		_ = scraper.ScrapeRentals(query)
	})

	t.Run("scrape url with 120 items", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(item120Handler))
		defer server.Close()
		query := &Query{
			RootURL: server.URL + "/?",
		}

		scraper := NewFiveN1()
		rentals := scraper.ScrapeRentals(query)

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
		query := &Query{
			RootURL: server.URL + "/?",
		}

		scraper := NewFiveN1()

		_ = scraper.ScrapeRentals(query)

		assert.Equal(t, 333, scraper.records)
		assert.Equal(t, 12, scraper.pages)
	})

	t.Run("scrape to Rentals", func(t *testing.T) {
		wantQuerySection := []string{
			//parse will request twice per section
			"98", "98",
			"99", "99",
			"100", "100",
		}
		var gotQuerySection []string
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotQuerySection = append(gotQuerySection, r.URL.Query().Get("section"))

			html, _ := ioutil.ReadFile("test_fixture/591with2items.html")
			w.Header().Set("Content-Type", "application/html")
			_, _ = w.Write(html)
		}))
		defer server.Close()

		query := &Query{
			RootURL: server.URL + "/?",
			Section: "98,99,100",
		}

		f := NewFiveN1()
		gotRentals := f.ScrapeRentals(query)

		wantRentals := Rentals{
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				PostBy:     "代理人 高先生",
				Phone:      "",
				Section:    "98",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				PostBy:     "仲介 李士豪",
				Phone:      "",
				Section:    "98",
			},
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				PostBy:     "代理人 高先生",
				Phone:      "",
				Section:    "99",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				PostBy:     "仲介 李士豪",
				Phone:      "",
				Section:    "99",
			},
			Rental{
				Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
				URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
				Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
				OptionType: "整層住家",
				Ping:       "128",
				Floor:      "樓層：整棟",
				Price:      "48,000 元 / 月",
				ID:         "R9538360",
				PostBy:     "代理人 高先生",
				Phone:      "",
				Section:    "100",
			},
			Rental{
				Title:      "中興大學賺錢店面",
				URL:        "https://rent.591.com.tw/rent-detail-9484376.html",
				Address:    "賺錢住店南區 - 建成路 1727 號",
				OptionType: "整層住家",
				Ping:       "50.8",
				Floor:      "樓層：1/12",
				Price:      "50,000 元 / 月",
				ID:         "R9484376",
				PostBy:     "仲介 李士豪",
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

		// quick and dirty test
		if !reflect.DeepEqual(wantQuerySection, gotQuerySection) {
			t.Errorf("query section not equal, want %v, got %v", wantQuerySection, gotQuerySection)
		}
	})
}

func TestFiveN1_ScrapeDetail(t *testing.T) {

	const rentalDetailPath = "/rent-detail-9538360.html"
	rental := &Rental{
		Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
		URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
		Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
		OptionType: "整層住家",
		Ping:       "128",
		Floor:      "樓層：整棟",
		Price:      "48,000 元 / 月",
		ID:         "R9538360",
		Phone:      "",
		Section:    "99",
	}

	t.Run("request rental.URL", func(t *testing.T) {

		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, rentalDetailPath, r.URL.Path)
		}))
		defer svr.Close()
		rental.URL = svr.URL + strings.TrimPrefix(rental.URL, "https://rent.591.com.tw")

		scraper := NewFiveN1()
		_ = scraper.ScrapeRentalDetail(rental)
	})

	t.Run("update phone, community and layout", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testFixture := "test_fixture/591_detail.html"
			html, _ := ioutil.ReadFile(testFixture)
			//w.Header().Set("Content-Type", "application/html")
			//w.WriteHeader(http.StatusOK)
			_, _ = w.Write(html)
		}))
		defer svr.Close()

		rental.URL = svr.URL + rentalDetailPath

		scraper := NewFiveN1()
		_ = scraper.ScrapeRentalDetail(rental)

		assert.Equal(t, "0980-240-200", rental.Phone, "rental.Phone not equal")
		assert.Equal(t, "6房3廳4衛4陽台", rental.Layout, "rental.Layout not equal")
		//need a better fixture
		assert.Equal(t, "近好事多南區西區向上路黎明路永春東路", rental.Community, "rental.Community not equal")
	})

	t.Run("scrape rental without layout", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testFixture := "test_fixture/591_detail_without_layout.html"
			html, _ := ioutil.ReadFile(testFixture)
			_, _ = w.Write(html)
		}))
		defer svr.Close()

		rental := &Rental{
			Title:      "近中山醫，近愛買捷運套房",
			URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
			Address:    "南區-福田五街20號",
			OptionType: "型態 :  電梯大樓",
			Ping:       "12",
			Floor:      "樓層 :  3F/15F",
			Price:      "15,000 元/月",
			ID:         "R9491919",
			Phone:      "",
			Section:    "99",
		}
		rental.URL = svr.URL + rentalDetailPath

		scraper := NewFiveN1()
		_ = scraper.ScrapeRentalDetail(rental)

		assert.Equal(t, "0986-851-077 轉 1397162", rental.Phone, "rental.Phone not equal")
		assert.Equal(t, "", rental.Layout, "rental.Layout not equal")
		assert.Equal(t, "豐邑閱文心", rental.Community, "rental.Community not equal")
	})
}

func TestFiveN1_ScrapeRentalsDetail(t *testing.T) {
	const rentalDetailPath = "/rent-detail-9538360.html"
	rental := Rental{
		Title:      "稀有花園別墅⭐別墅透天⭐雙平車⭐可寵",
		URL:        "https://rent.591.com.tw/rent-detail-9538360.html",
		Address:    "近好事多南區西區向上路黎明路永春東路南屯區 - 惠中路三段",
		OptionType: "整層住家",
		Ping:       "128",
		Floor:      "樓層：整棟",
		Price:      "48,000 元 / 月",
		ID:         "R9538360",
		Phone:      "",
		Section:    "99",
	}
	rentals := Rentals{
		rental,
		rental,
	}

	t.Run("update all rental.Phone", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			testFixture := "test_fixture/591_detail.html"
			html, _ := ioutil.ReadFile(testFixture)
			//w.Header().Set("Content-Type", "application/html")
			//w.WriteHeader(http.StatusOK)
			_, _ = w.Write(html)
		}))
		defer svr.Close()

		rental.URL = svr.URL + rentalDetailPath

		scraper := NewFiveN1()
		scraper.ScrapeRentalsDetail(rentals)

		assert.Equal(t, "0980-240-200", rentals[0].Phone)
		assert.Equal(t, "0980-240-200", rentals[1].Phone)
	})
}
