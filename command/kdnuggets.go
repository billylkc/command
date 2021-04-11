package command

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/now"
)

// KDRecord
type KDRecord struct {
	Updated time.Time
	Title   string
	Content string
	Link    string
	Summary string
	Tags    string
}

// KDNuggets retrieves the records for some months
func KDNuggets(d string, nrecords int) ([]KDRecord, error) {
	var results []KDRecord
	var ll []string // list of dates to be looped over, e.g. ["2021-04-01", "2021-03-01"]

	t, err := now.Parse(d)
	if err != nil {
		return results, err
	}

	for i := 1; i <= nrecords; i++ {
		tt := t.AddDate(0, -(i - 1), 0)
		ll = append(ll, tt.Format("2006-01-02"))
	}

	for _, dd := range ll {
		tt, err := now.Parse(dd)
		if err != nil {
			return results, err
		}

		year := tt.Year()
		month := fmt.Sprintf("%02d", tt.Month())
		url := fmt.Sprintf("https://www.kdnuggets.com/%d/%s/index.html", year, month)
		res, err := http.Get(url)
		if err != nil {
			return results, err
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			return results, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return results, err
		}

		// Find the review items
		doc.Find("ul.three_ul.test li").Each(func(i int, s *goquery.Selection) {
			var (
				rec  KDRecord
				tags []string
			)
			title := s.Find("a").Text()
			dateF := strings.ReplaceAll(s.Find("font").Text(), "- ", "")
			fmt.Println(dateF)
			date, _ := time.Parse("Jan 2, 2006", dateF)
			link, _ := s.Find("a").Attr("href")
			snippet := strings.TrimSpace(s.Find("div").Text())
			tag := strings.ReplaceAll(s.Find("p").Text(), "Tags: ", "")
			tt := strings.Split(tag, ",")
			for _, t := range tt {
				t = strings.TrimSpace(t)
				tags = append(tags, t)
			}
			rec = KDRecord{
				Updated: date,
				Title:   title,
				Content: snippet,
				Link:    link,
				Tags:    strings.Join(tags, ", "),
			}
			results = append(results, rec)
		})
	}
	return results, nil
}
