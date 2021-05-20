package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Quandl struct {
	token     string
	limit     int
	startDate string // e.g. 2019-01-01
	endDate   string
	order     string
}

type option func(*Quandl)

func main() {

	quandl := NewQuandl()
	quandl.Option(SetLimit(10))
	quandl.Option(SetEndDate("2021-01-02"))
	quandl.Option(SetOrder("desc"))

	result, err := quandl.GetHistoricalPrice(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(PrettyPrint(result))
}

type HistoricalPrice struct {
	DatasetData struct {
		Limit       interface{}     `json:"limit"`
		Transform   interface{}     `json:"transform"`
		ColumnIndex interface{}     `json:"column_index"`
		ColumnNames []string        `json:"column_names"`
		StartDate   string          `json:"start_date"`
		EndDate     string          `json:"end_date"`
		Frequency   string          `json:"frequency"`
		Data        [][]interface{} `json:"data"`
		Collapse    interface{}     `json:"collapse"`
		Order       string          `json:"order"`
	} `json:"dataset_data"`
}

func NewQuandl() Quandl {

	token, err := getQuanToken()
	if err != nil {
		log.Fatal(err.Error())
	}

	return Quandl{
		limit:     100,
		token:     token,
		startDate: "2015-01-01",                    // Hard code for now, as there is a row limit for api call
		endDate:   time.Now().Format("2006-01-02"), // Default for today, in yyyy-mm-dd format
		order:     "desc",
	}
}

func (q *Quandl) GetHistoricalPrice(code int) (HistoricalPrice, error) {
	var data HistoricalPrice
	c := fmt.Sprintf("%05d", code)
	endpoint, err := q.getQuanEndPoint("HKEX", c, "historicalPrice")
	resp, err := http.Get(endpoint)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data, nil
}

// getQuanEndPoint returns the endpoint of the api callOB
func (q *Quandl) getQuanEndPoint(db, code, api string) (string, error) {
	var endpoint string
	switch api {
	case "historicalPrice":
		endpoint = "https://www.quandl.com/api/v3/datasets/HKEX/00005/data.json?api_key=%s&start_date=%s&end_date=%s&order=%s&limit=%d"
	default:
		return "", fmt.Errorf("no api endpoint - &s", api)
	}
	endpoint = fmt.Sprintf(endpoint, q.token, q.startDate, q.endDate, q.order, q.limit)

	return endpoint, nil
}

// getQuanToken returns the
func getQuanToken() (string, error) {
	token, ok := os.LookupEnv("QUANDL_TOKEN")
	if !ok {
		return token, errors.New("quandl token not set, please check your env variable QUANDL_TOKEN")
	}
	return token, nil
}

// Option sets the options specified.
func (q *Quandl) Option(opts ...option) {
	for _, opt := range opts {
		opt(q)
	}
}

// SetLimit sets Foo's verbosity level to v.
func SetLimit(v int) option {
	return func(q *Quandl) {
		q.limit = v
	}
}

// SetEndDate sets end date for the query in yyyy-mm-dd format
func SetEndDate(v string) option {
	return func(q *Quandl) {
		q.endDate = v
	}
}

// SetStartDate sets start date for the query in yyyy-mm-dd format
func SetStartDate(v string) option {
	return func(q *Quandl) {
		q.startDate = v
	}
}

// SetOrder sets set the order of the query, e.g. asc/desc
func SetOrder(v string) option {
	return func(q *Quandl) {
		q.order = v
	}
}
