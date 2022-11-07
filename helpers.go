package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var wg sync.WaitGroup

// returns games and genres for given publisher URLs
func Scrape(p *Publisher) (map[string]string, string, error) {
	res, err := http.Get(p.StoreLink)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	games := make(map[string]string)

	doc.Find(".l-row a").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".we-lockup__title").Text()
		title = strings.TrimSpace(title)
		genre := s.Find(".we-lockup__subtitle").Text()
		genre = strings.TrimSpace(genre)
		fmt.Printf("Game: %d: %s - %s\n", i, title, genre)
		games[title] = genre
	})
	return games, p.Name, nil
}

// FetchFromStore fetches games from all publishers' stores
func FetchFromStore() {
	for _, v := range PUBLISHERS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Scrape(&v)
		}()
	}
	wg.Wait()
}
