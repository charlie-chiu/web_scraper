package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		url := server.URL + "/?"
		scraper := NewFiveN1()

		_ = scraper.ScrapeList(url)
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
		url := server.URL + "/?"
		scraper := NewFiveN1()
		scraper.SetReqCookie("20")

		_ = scraper.ScrapeList(url)
	})

	t.Run("scrape url with 120 items", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(item120Handler))
		defer server.Close()
		url := server.URL + "/?"

		scraper := NewFiveN1()

		rentals := scraper.ScrapeList(url)

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
		url := server.URL + "/?"

		scraper := NewFiveN1()

		_ = scraper.ScrapeList(url)

		assert.Equal(t, 333, scraper.records)
		assert.Equal(t, 12, scraper.pages)

		// this test will fail because test fixture always have 30 item per page.
		//assert.Equal(t, 333, len(rentals))
	})

	t.Run("scrape group by section and store section in Rental", func(t *testing.T) {
		var requestCount int

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestCount++
			html, _ := ioutil.ReadFile("test_fixture/591with2items.html")
			w.Header().Set("Content-Type", "application/html")
			_, _ = w.Write(html)
		}))
		defer server.Close()

		url := fmt.Sprintf("%s/?%s", server.URL, "wantSection=98,99,100")

		f := NewFiveN1()
		rentals := f.ScrapeList(url)

		// 2 request per section
		wantRequest := 6
		assert.Equal(t, wantRequest, requestCount)

		wantSections := []string{
			"98",
			"98",
			"99",
			"99",
			"100",
			"100",
		}
		for i, wantSection := range wantSections {
			if i >= len(rentals) {
				t.Fatalf("%dth rental not exsits", i)
			}
			gotSection := rentals[i].Section
			if wantSection != gotSection {
				t.Errorf("%dth rental wantSection not equal, want %s, got %s", i, wantSection, gotSection)
			}
		}
	})
}
