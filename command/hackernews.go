package command

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"astuart.co/goq"
	"github.com/jinzhu/now"
)

// Page
type Page struct {
	Content []string `goquery:"div.post-body p"`
}

// HackerNewsRec as a struct to capture data from hackernews.com
type HackerNewsRec struct {
	Title []string `goquery:"h2.home-title"`
	Date  []string `goquery:"div.item-label"`
	Link  []string `goquery:"a.story-link,[href]"`
}

// HackerNew as individual title
type HackerNew struct {
	Updated time.Time
	Title   string
	Content string
	Link    string
	Summary string
}

type HackerNews []HackerNew

// GetHackerNews gets the last two days news from hackernews
func GetHackerNews(all bool) ([]HackerNew, error) {
	var results HackerNews
	today := time.Now().Format("2006-01-02")
	url := fmt.Sprintf("https://thehackernews.com/search?updated-max=%sT23:55:00-08:00&max-results=%d", today, 10)
	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	var allNews HackerNewsRec
	err = goq.NewDecoder(res.Body).Decode(&allNews)
	if err != nil {

	}

	titles := allNews.Title
	dates := allNews.Date
	links := allNews.Link

	if len(titles) == len(dates) &&
		len(titles) == len(links) &&
		len(titles) > 0 {

		for i, _ := range titles {
			date, _ := convertDate(dates[i])
			abstract, _ := getPageContent(links[i])
			h := HackerNew{
				Updated: date,
				Title:   titles[i],
				Link:    links[i],
				Content: abstract,
			}
			results = append(results, h)
		}
	} else {
		return results, fmt.Errorf("invalid lengths of return lists")
	}

	if len(results) == 0 {
		return results, fmt.Errorf("no results")
	}

	return results, nil
}

// extract date from this "March 02, 2021Ravie Lakshmanan"
func convertDate(s string) (t time.Time, err error) {
	matches := strings.Split(s, "")
	if len(matches) > 0 {
		date := matches[0]
		date = strings.ReplaceAll(date, "", "")
		t, err = now.Parse(date)
	}
	return
}

// getPageContent gets individual page content
func getPageContent(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var page Page
	err = goq.NewDecoder(res.Body).Decode(&page)
	if err != nil {
		return "", err
	}
	return strings.Join(page.Content, "\n\n\t"), nil
}
