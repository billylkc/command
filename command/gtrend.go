package command

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/billylkc/myutil"
	"github.com/groovili/gogtrends"
	. "github.com/logrusorgru/aurora"
)

type Trend struct {
	Topic   string
	Title   string
	Traffic string
	URL     string
	Snippet string
}

// gtrend returns the google search keywords trend
func Gtrend(country string) ([]Trend, error) {
	ctx := context.Background()
	if country == "" { // default HK
		country = "HK"
	}
	country = strings.Title(country)
	return daily(ctx, country)
}

// daily prints daily trends from country
func daily(ctx context.Context, country string) ([]Trend, error) {
	var results []Trend
	log.Println("Daily trending searches:")
	dailySearches, err := gogtrends.Daily(ctx, "en", country)
	if err != nil {
		return results, err
	}

	for _, d := range dailySearches {

		topic := fmt.Sprintf("%v", *d.Title)
		traffic := d.FormattedTraffic

		for _, a := range d.Articles {

			title := fmt.Sprintf("%s", fmt.Sprintf("[%s] - %s", a.Source, Magenta(a.Title)))
			url := fmt.Sprintf(a.URL)
			snippet := a.Snippet

			t := Trend{
				Topic:   topic,
				Traffic: traffic,
				Title:   fmt.Sprintf("%s\n\n%s", title, myutil.BreakLongStr(snippet, 80, 130)),
				URL:     url,
				// Snippet: snippet,
			}
			results = append(results, t)
		}
	}
	return results, nil
}

// getDateFromHours gets the date from string like "4h ago, 1d ago"
// returns yyyy-mm-dd
func getDateFromHours(s string) (string, error) {
	var res string
	if strings.Contains(s, "d") {
		res = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	} else {
		res = time.Now().Format("2006-01-02")
	}
	return res, nil
}
