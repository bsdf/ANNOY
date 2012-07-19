package main

import (
	"flag"
	"fmt"
	"github.com/bsdf/twitter"
	"strings"
	"time"
)

var timeout int

func init() {
	flag.IntVar(&timeout, "t", 5, "timeout between tweets in seconds")
}

func main() {
	flag.Parse()

	t := twitter.Twitter{
		ConsumerKey:      "YOUR_CONSUMER_KEY",
		ConsumerSecret:   "YOUR_CONSUMER_SECRET",
		OAuthToken:       "YOUR_OAUTH_TOKEN",
		OAuthTokenSecret: "YOUR_OAUTH_TOKEN_SECRET",
	}

	str := strings.Join(flag.Args(), " ")

	for {
		tweet, err := t.Tweet(str)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		time.Sleep(time.Duration(timeout) * time.Second)
		_, err = t.Destroy(tweet.Id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
