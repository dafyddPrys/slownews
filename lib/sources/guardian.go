package sources

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	slownews "github.com/dafyddprys/slownews/lib"
)

type Guardian struct {
	apiKey string
}

func New(key string) Guardian {
	return Guardian{
		apiKey: key,
	}
}

// GetNews satisfies the Source interface. It gets news articles from the last 24 hours
func (g Guardian) GetNews() (error, []slownews.Article) {
	startTime := time.Now().AddDate(0, 0, -1) // take 1 day off time.
	guardianUrl := fmt.Sprintf("https://content.guardianapis.com/search?from-date=%s&api-key=%s",
		url.QueryEscape(startTime.Format(time.RFC3339)),
		g.apiKey)
	resp, err := http.Get(guardianUrl)
	if err != nil {
		return fmt.Errorf("could not get news from Guardian: %s", err), nil
	}
	fmt.Println("ok")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could read response from Guardian: %s", err), nil
	}

	fmt.Printf("%s", body)
	return nil, []slownews.Article{}
}

// func parseBody(in string) (error, []slownews.Article) {

// }
