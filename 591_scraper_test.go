package scraper

import (
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

		assert.Equal(t, 120, len(rentals))
	})
}
