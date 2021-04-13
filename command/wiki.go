// Easy access for wiki page
// Search the terms with duck duck go
// Then Quickly print the first two paragraph of the wiki page queried

package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"astuart.co/goq"
	"github.com/billylkc/command/util"
	"github.com/billylkc/gtoolkits"
)

// Wiki as the result from wikipedia
type Wiki struct {
	Title   string   `goquery:"h1"`
	Content []string `goquery:"div.mw-parser-output p"`
	Summary string
	Link    string
}

type Duck struct {
	Data struct {
		Abstract         string `json:"Abstract"`
		AbstractSource   string `json:"AbstractSource"`
		AbstractText     string `json:"AbstractText"`
		AbstractURL      string `json:"AbstractURL"`
		Answer           string `json:"Answer"`
		AnswerType       string `json:"AnswerType"`
		Definition       string `json:"Definition"`
		DefinitionSource string `json:"DefinitionSource"`
		DefinitionURL    string `json:"DefinitionURL"`
		Entity           string `json:"Entity"`
		Heading          string `json:"Heading"`
		Image            string `json:"Image"`
		Redirect         string `json:"Redirect"`
		RelatedTopics    []struct {
			FirstURL string `json:"FirstURL,omitempty"`
			Result   string `json:"Result,omitempty"`
			Text     string `json:"Text,omitempty"`
			Name     string `json:"Name,omitempty"`
			Topics   []struct {
				FirstURL string `json:"FirstURL"`
				Result   string `json:"Result"`
				Text     string `json:"Text"`
			} `json:"Topics,omitempty"`
		} `json:"RelatedTopics"`
		Results []interface{} `json:"Results"`
		Type    string        `json:"Type"`
		Meta    struct {
			Attribution  interface{} `json:"attribution"`
			Blockgroup   interface{} `json:"blockgroup"`
			CreatedDate  interface{} `json:"created_date"`
			Description  string      `json:"description"`
			Designer     interface{} `json:"designer"`
			DevDate      interface{} `json:"dev_date"`
			DevMilestone string      `json:"dev_milestone"`
			Developer    []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"developer"`
			ExampleQuery    string      `json:"example_query"`
			ID              string      `json:"id"`
			IsStackexchange interface{} `json:"is_stackexchange"`
			JsCallbackName  string      `json:"js_callback_name"`
			LiveDate        interface{} `json:"live_date"`
			Maintainer      struct {
				Github string `json:"github"`
			} `json:"maintainer"`
			Name            string      `json:"name"`
			PerlModule      string      `json:"perl_module"`
			Producer        interface{} `json:"producer"`
			ProductionState string      `json:"production_state"`
			Repo            string      `json:"repo"`
			SignalFrom      string      `json:"signal_from"`
			SrcDomain       string      `json:"src_domain"`
			SrcID           int         `json:"src_id"`
			SrcName         string      `json:"src_name"`
			SrcOptions      struct {
				Directory         string `json:"directory"`
				IsFanon           int    `json:"is_fanon"`
				IsMediawiki       int    `json:"is_mediawiki"`
				IsWikipedia       int    `json:"is_wikipedia"`
				Language          string `json:"language"`
				MinAbstractLength string `json:"min_abstract_length"`
				SkipAbstract      int    `json:"skip_abstract"`
				SkipAbstractParen int    `json:"skip_abstract_paren"`
				SkipEnd           string `json:"skip_end"`
				SkipIcon          int    `json:"skip_icon"`
				SkipImageName     int    `json:"skip_image_name"`
				SkipQr            string `json:"skip_qr"`
				SourceSkip        string `json:"source_skip"`
				SrcInfo           string `json:"src_info"`
			} `json:"src_options"`
			SrcURL interface{} `json:"src_url"`
			Status string      `json:"status"`
			Tab    string      `json:"tab"`
			Topic  []string    `json:"topic"`
			Unsafe int         `json:"unsafe"`
		} `json:"meta"`
	} `json:"data"`
	DuckbarTopic string `json:"duckbar_topic"`
	Meta         struct {
		Attribution  interface{} `json:"attribution"`
		Blockgroup   interface{} `json:"blockgroup"`
		CreatedDate  interface{} `json:"created_date"`
		Description  string      `json:"description"`
		Designer     interface{} `json:"designer"`
		DevDate      interface{} `json:"dev_date"`
		DevMilestone string      `json:"dev_milestone"`
		Developer    []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"developer"`
		ExampleQuery    string      `json:"example_query"`
		ID              string      `json:"id"`
		IsStackexchange interface{} `json:"is_stackexchange"`
		JsCallbackName  string      `json:"js_callback_name"`
		LiveDate        interface{} `json:"live_date"`
		Maintainer      struct {
			Github string `json:"github"`
		} `json:"maintainer"`
		Name            string      `json:"name"`
		PerlModule      string      `json:"perl_module"`
		Producer        interface{} `json:"producer"`
		ProductionState string      `json:"production_state"`
		Repo            string      `json:"repo"`
		SignalFrom      string      `json:"signal_from"`
		SrcDomain       string      `json:"src_domain"`
		SrcID           int         `json:"src_id"`
		SrcName         string      `json:"src_name"`
		SrcOptions      struct {
			Directory         string `json:"directory"`
			IsFanon           int    `json:"is_fanon"`
			IsMediawiki       int    `json:"is_mediawiki"`
			IsWikipedia       int    `json:"is_wikipedia"`
			Language          string `json:"language"`
			MinAbstractLength string `json:"min_abstract_length"`
			SkipAbstract      int    `json:"skip_abstract"`
			SkipAbstractParen int    `json:"skip_abstract_paren"`
			SkipEnd           string `json:"skip_end"`
			SkipIcon          int    `json:"skip_icon"`
			SkipImageName     int    `json:"skip_image_name"`
			SkipQr            string `json:"skip_qr"`
			SourceSkip        string `json:"source_skip"`
			SrcInfo           string `json:"src_info"`
		} `json:"src_options"`
		SrcURL interface{} `json:"src_url"`
		Status string      `json:"status"`
		Tab    string      `json:"tab"`
		Topic  []string    `json:"topic"`
		Unsafe int         `json:"unsafe"`
	} `json:"meta"`
	Model     string `json:"model"`
	PixelID   string `json:"pixel_id"`
	Signal    string `json:"signal"`
	Templates struct {
		Item string `json:"item"`
	} `json:"templates"`
}

// GetWiki will search first in wiki, if not found, use gogoduck, then wiki again
func GetWiki(query string) (Wiki, error) {

	result, err := tryWiki(query)

	// if cant find in first try, use duck duck go
	if err != nil {
		fmt.Println(err)
		link, err := duck(query)
		if err != nil {
			return result, err
		}
		result = wiki(link)

	}
	err = result.deriveSummary()
	if err != nil {
		fmt.Println("Can not connect to gRPC server.")
	}
	return result, nil
}

// deriveSummary will send request to gRPC server to extract the keywords
// keywrods will be highlighted before print for readability
func (w *Wiki) deriveSummary() error {
	var content string
	if len(w.Content) >= 3 {
		content = strings.Join(w.Content[0:3], "\n\n")
	} else {
		content = strings.Join(w.Content, "\n")
	}

	keywords, err := gtoolkits.GetKeywords(content, 20)
	if err != nil {
		w.Summary = content
		return err
	}
	for _, k := range keywords {
		content = util.InsensitiveReplace(content, k, true)
	}
	w.Summary = content
	return nil
}

// duck will return the wiki link after searching for the query string
func duck(q string) (string, error) {
	query := strings.ReplaceAll(q, "_", "+")
	response, err := http.Get(fmt.Sprintf("https://duckduckgo.com/?q=%s", query))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	// s := string(body)

	re, _ := regexp.Compile(`DDG\.duckbar\.add\((.*)\);}\);`)
	matches := re.FindAllSubmatch(body, -1)
	var duck Duck
	if len(matches) == 1 {
		byt := matches[0][1]
		if err := json.Unmarshal(byt, &duck); err != nil {
			return "", err
		}
	} else {
		return "", fmt.Errorf("not matched")
	}

	var url string
	abstract := duck.Data.AbstractText
	if abstract == "" {
		url = fmt.Sprintf("https://en.wikipedia.org/wiki%s", duck.Data.RelatedTopics[0].FirstURL)
	} else {
		url = duck.Data.AbstractURL
	}
	return url, nil
}

// wiki gets the wikipedia link, and print the first few paragraph from the page
func wiki(link string) Wiki {

	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var wiki Wiki
	err = goq.NewDecoder(res.Body).Decode(&wiki)
	if err != nil {
		log.Fatal(err)
	}
	wiki.Link = link
	return wiki
}

func tryWiki(q string) (Wiki, error) {
	var result Wiki
	link := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", q)
	result = wiki(link)

	// check if contains "may refer to:"
	if result.Content == nil {
		return result, fmt.Errorf("(search in duck duck go)")
	} else {
		if len(result.Content) == 1 { // Wiki suggestions
			if strings.Contains(result.Content[0], "may refer to:") {
				return result, fmt.Errorf("(search in duck duck go)")
			}
			if strings.Contains(result.Content[0], "may also refer to:") {
				return result, fmt.Errorf("(search in duck duck go)")
			}
		}
	}
	return result, nil
}
