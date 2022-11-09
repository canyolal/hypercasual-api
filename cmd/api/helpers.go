package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type PublisherList struct {
	Name      string
	StoreLink string
}

// List of publishers and store links
var PUBLISHERS = []PublisherList{
	{
		Name:      "Voodoo",
		StoreLink: "https://apps.apple.com/us/developer/voodoo/id714804730?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Good Job Games",
		StoreLink: "https://apps.apple.com/tr/developer/good-job-games/id1191495496?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Ketchapp",
		StoreLink: "https://apps.apple.com/us/developer/ketchapp/id528065807?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Alictus",
		StoreLink: "https://apps.apple.com/tr/developer/alictus/id892399717?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Lion Studios",
		StoreLink: "https://apps.apple.com/us/developer/lion-studios/id1362220666?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Rollic",
		StoreLink: "https://apps.apple.com/us/developer/rollic-games/id1452111779?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Kwalee",
		StoreLink: "https://apps.apple.com/tr/developer/kwalee-ltd/id497961736?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "BoomBit",
		StoreLink: "https://apps.apple.com/kh/developer/boombit-inc/id1045926022?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Amanotes",
		StoreLink: "https://apps.apple.com/us/developer/amanotes-pte-ltd/id1441389613?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Azur Games",
		StoreLink: "https://apps.apple.com/us/developer/azur-interactive-games-limited/id1296347323?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Crazy Labs",
		StoreLink: "https://apps.apple.com/us/developer/crazy-labs/id721307559?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Coda",
		StoreLink: "https://apps.apple.com/us/developer/coda-platform-limited/id1475474579?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Ducky",
		StoreLink: "https://apps.apple.com/tr/developer/ducky-ltd/id1541013213?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Gismart",
		StoreLink: "https://apps.apple.com/us/developer/gismart/id666830030?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Green Panda Games",
		StoreLink: "https://apps.apple.com/tr/developer/green-panda-games/id669978473?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Homa",
		StoreLink: "https://apps.apple.com/tr/developer/homa-games/id1508492426?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "JoyPac",
		StoreLink: "https://apps.apple.com/tr/developer/joypac/id1422558565?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Moonee",
		StoreLink: "https://apps.apple.com/us/developer/moonee-publishing-ltd/id1469957859?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Playgendary",
		StoreLink: "https://apps.apple.com/us/developer/playgendary-limited/id1487320337?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "SayGames",
		StoreLink: "https://apps.apple.com/tr/developer/saygames-ltd/id1551847165?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Supersonic",
		StoreLink: "https://apps.apple.com/us/developer/supersonic-studios-ltd/id1499845738?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "TapNation",
		StoreLink: "https://apps.apple.com/tr/developer/tapnation/id1483575279?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Tastypill",
		StoreLink: "https://apps.apple.com/us/developer/tastypill/id1022434729?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Yso Corp",
		StoreLink: "https://apps.apple.com/us/developer/yso-corp/id659815325?see-all=i-phonei-pad-apps",
	},
	{
		Name:      "Zplay",
		StoreLink: "https://apps.apple.com/tr/developer/zplay-beijing-info-tech-co-ltd/id531022725?see-all=i-phonei-pad-apps",
	},
}

var wg sync.WaitGroup

// envelope type to nest json under struct name
type envelope map[string]interface{}

// returns games-genre duo from publisher's app store page
func Scrape(p *PublisherList) (map[string]string, string, error) {
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
		go func(p *PublisherList) {
			defer wg.Done()
			Scrape(p)
		}(&v)
	}
	wg.Wait()
}

// writeJSON() is a helper for sending JSON responses. Receives http.ResponseWriter,
// HTTP status code, HTTP headers and data to be sent as JSON.
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')

	// Header() returns a Header map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

// readJSON() reads JSON obj and returns error specific human readable messages in case of errors.
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {

		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		// Use the errors.As() function to check whether the error has the type
		// *json.SyntaxError. If it does, then return a plain-english error message
		// which includes the location of the problem.
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains a badly-formed JSON (at character %d)", syntaxError.Offset)

		// In some circumstances Decode() may also return an io.ErrUnexpectedEOF error
		// for syntax errors in the JSON. So we check for this using errors.Is() and
		// return a generic error message. There is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		// An io.EOF error will be returned by Decode() if the request body is empty. We
		// check for this with errors.Is() and return a plain-english error message
		// instead.
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// An io.EOF error will be returned by Decode() if the request body is empty. We
		// check for this with errors.Is() and return a plain-english error message
		// instead.
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// A json.InvalidUnmarshalError error will be returned if we pass a non-nil
		// pointer to Decode(). We catch this and panic, rather than returning an error
		// to our handler. At the end of this chapter we'll talk about panicking
		// versus returning errors, and discuss why it's an appropriate thing to do in
		// this specific situation.
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		// If the JSON contains a field which cannot be mapped to the target destination
		// then Decode() will now return an error message in the format "json: unknown
		// field "<name>"". We check for this, extract the field name from the error,
		// and interpolate it into our custom error message. Note that there's an open
		// issue at https://github.com/golang/go/issues/29035 regarding turning this
		// into a distinct error type in the future.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)

		// If the request body exceeds 1MB in size the decode will now fail with the
		// error "http: request body too large". There is an open issue about turning
		// this into a distinct error type at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}
