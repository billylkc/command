package command

import (
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/billylkc/myutil"
)

// ListenTweets listens to twitch through developer api
func ListenTweets(tags ...string) {
	consumerKey, _ := myutil.GetEnv("TWITTER_CONSUMER_KEY")
	consumerSecret, _ := myutil.GetEnv("TWITTER_CONSUMER_SECRET")
	accessToken, _ := myutil.GetEnv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret, _ := myutil.GetEnv("TWITTER_ACCESS_TOKEN_SECRET")

	fmt.Printf("Listening to tags - %v\n\n", tags)

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	stream := api.PublicStreamFilter(url.Values{
		"track": tags,
	})
	defer stream.Stop()

	for v := range stream.C {
		t, ok := v.(anaconda.Tweet)
		if !ok {
			fmt.Printf("received unexpected value of type %T\n", v)
			continue
		}
		fmt.Println("------------------------")
		fmt.Println(t.FullText)
	}
}
