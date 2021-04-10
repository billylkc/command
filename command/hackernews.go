package command

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"

	"astuart.co/goq"
	. "github.com/logrusorgru/aurora"
)

// HackerNews as a struct to capture data from hackernews.com
type HackerNews struct {
	Title []string `goquery:"h2.home-title"`
	Date  []string `goquery:"div.item-label"`
	Link  []string `goquery:"a.story-link,[href]"`
}

// HackerNew as individual title
type HackerNew struct {
	Date     string
	Title    string
	Link     string
	Abstract string
}

// Page
type Page struct {
	Content []string `goquery:"div.post-body p"`
}

// GetHackerNews gets the last two days news from hackernews
func GetHackerNews(all bool) ([]HackerNew, error) {
	var results []HackerNew
	today := time.Now().Format("2006-01-02")
	todayF := time.Now().Format("January 02, 2006")                 // another format for checking, e.g. March 05, 2021
	ytdF := time.Now().AddDate(0, 0, -1).Format("January 02, 2006") // allow yesterday as well
	url := fmt.Sprintf("https://thehackernews.com/search?updated-max=%sT23:55:00-08:00&max-results=%d", today, 10)
	res, err := http.Get(url)
	if err != nil {
		return results, err
	}
	defer res.Body.Close()

	var allNews HackerNews

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
			date := convertDate(dates[i])
			abstract, _ := getPageContent(links[i])
			h := HackerNew{
				Date:     date,
				Title:    titles[i],
				Link:     links[i],
				Abstract: abstract,
			}
			if all {
				results = append(results, h)
			} else {
				// just today and ytd
				if h.Date == todayF || h.Date == ytdF {
					results = append(results, h)
				}
			}
		}
	} else {
		return results, fmt.Errorf("invalid lengths of return lists")
	}

	if len(results) == 0 {
		return results, fmt.Errorf("no results")
	}

	return results, nil
}

// PrintResult is just some group by date formatting, decending order
func PrintResult(news []HackerNew) {
	var keys []string
	m := make(map[string][]HackerNew) // Group by Date

	for _, n := range news {
		if val, ok := m[n.Date]; ok {
			m[n.Date] = append(val, n)
		} else {
			m[n.Date] = []HackerNew{n}
			keys = append(keys, n.Date)
		}
	}

	// sort date
	sort.Slice(keys, func(i, j int) bool {
		return keys[j] < keys[i]
	})

	for _, k := range keys {
		fmt.Println("")
		fmt.Println(Bold(Yellow(k))) // Date
		fmt.Println("")

		for _, rec := range m[k] {
			fmt.Printf("\tTitle - %s\n\n\t%s\n\n\t%s\n\t-------------------------------------------------\n\n",
				Yellow(rec.Title),
				Cyan(rec.Link),
				White(rec.Abstract))
		}
	}
}

// extract date from this "March 02, 2021Ravie Lakshmanan"
func convertDate(s string) string {
	matches := strings.Split(s, "")
	if len(matches) > 0 {
		date := matches[0]
		date = strings.ReplaceAll(date, "", "")
		return date
	}
	return ""

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
