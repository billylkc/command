package command

import (
	"fmt"
	"sort"
	"strings"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/jinzhu/now"
	"github.com/mmcdole/gofeed"
)

type RedditTopic struct {
	Updated time.Time
	Title   string
	Content string
	Link    string
}

// Reddit reads the rss from the subreddit
func Reddit(subreddit string) ([]RedditTopic, error) {
	var topics []RedditTopic
	fp := gofeed.NewParser()
	endPoint := fmt.Sprintf("http://www.reddit.com/r/%s/.rss", subreddit)
	feed, err := fp.ParseURL(endPoint)
	if err != nil {
		return topics, nil
	}

	// fmt.Println(PrettyPrint(feed))
	for _, f := range feed.Items {
		content := strings.TrimSpace(strip.StripTags(f.Content))
		if strings.HasPrefix(content, "&#32; ") {
			content = ""
		}
		t, _ := now.Parse(f.Updated)
		topic := RedditTopic{
			Title:   f.Title,
			Content: content,
			Link:    f.Link,
			Updated: t,
		}
		topics = append(topics, topic)
	}

	// sort by updated descending order
	sort.Slice(topics, func(i, j int) bool {
		return topics[i].Updated.After(topics[j].Updated)
	})
	return topics, nil
}
