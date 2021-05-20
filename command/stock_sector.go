package command

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Sector for industry sector overview
type Sector struct {
	Sector      string
	ChangePct   float64
	PchangePct  float64
	Turnover    int
	AvgTurnover int
	AvgPE       float64
	ZoneA       int // > +2%
	ZoneB       int // 0 - +2%
	ZoneC       int // 0%
	ZoneD       int // 0 - -2%
	ZoneE       int // < -2%
	ZoneN       int // total no of stocks
}

// GetSectorOverview gets the overview of the industry
func GetSectorOveriew() ([]Sector, error) {
	var results []Sector
	link := "http://www.aastocks.com/en/stocks/market/industry/industry-performance.aspx"

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var (
		sector      string
		changePct   float64
		pchangePct  float64
		turnover    int
		avgTurnover int
		avgPE       float64
		zone        string // from dist string `0,2,2,9,5`
		zoneA       int    // > +2%
		zoneB       int    //  0 - +2%
		zoneC       int    // 0%
		zoneD       int    // -2% - 0
		zoneE       int    // < -2%
		zoneN       int    // No of stocks
	)

	doc.Find("table.indview_tbl").Each(func(i int, s *goquery.Selection) {
		s.Find("tr.indview_tr").Each(func(j int, tr *goquery.Selection) {
			var elements []string
			tr.Find("td").Each(func(k int, td *goquery.Selection) {
				dist, exists := td.Find("div.jsPerfDistBar").Attr("def")
				if exists {
					zone = dist
					zones := strings.Split(zone, ",")
					zoneA, _ = strconv.Atoi(zones[0])
					zoneB, _ = strconv.Atoi(zones[1])
					zoneC, _ = strconv.Atoi(zones[2])
					zoneD, _ = strconv.Atoi(zones[3])
					zoneE, _ = strconv.Atoi(zones[4])
					zoneN = zoneA + zoneB + zoneC + zoneD + zoneE
				}
				elements = append(elements, td.Text())
			})
			if len(elements) >= 6 {
				sector = strings.TrimSpace(elements[0])
				changePct, _ = ParseF(elements[1])
				pchangePct, _ = ParseF(elements[2])
				turnover, _ = ParseI(elements[3])
				avgTurnover, _ = ParseI(elements[4])
				avgPE, _ = ParseF(elements[5])
			}
			s := Sector{
				Sector:      sector,
				ChangePct:   changePct,
				PchangePct:  pchangePct,
				Turnover:    turnover,
				AvgTurnover: avgTurnover,
				AvgPE:       avgPE,
				ZoneA:       zoneA,
				ZoneB:       zoneB,
				ZoneC:       zoneC,
				ZoneD:       zoneD,
				ZoneE:       zoneE,
				ZoneN:       zoneN,
			}
			results = append(results, s)
		})
	})
	return results, nil
}

// ParseF parses from string to float, ignoring % sign
func ParseF(s string) (float64, error) {
	s = strings.ReplaceAll(s, "%", "")
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, nil
	}
	return num, err
}

// ParseI parses from string to integer, also handles numbers like 1K, 1M, 1B, etc
func ParseI(s string) (int, error) {
	var (
		num int
		f   float64
		err error
	)

	if strings.Contains(s, "N/A") {
		num = 0
	}
	if strings.Contains(s, "K") {
		s = strings.ReplaceAll(s, "K", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000)
	}
	if strings.Contains(s, "M") {
		s = strings.ReplaceAll(s, "M", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000_000)
	}
	if strings.Contains(s, "B") {
		s = strings.ReplaceAll(s, "B", "")
		f, err = strconv.ParseFloat(s, 64)
		num = int(f * 1_000_000_000)
	}
	if err != nil {
		return 0, err
	}

	return num, nil
}
