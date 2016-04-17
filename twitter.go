package main

import (
	"log"
	"sync"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	tweetsMutex sync.RWMutex
	tweets      []twitter.Tweet
)

type tweetHandlerFunc func(tweet *twitter.Tweet)

func GetTweets() []twitter.Tweet {
	tweetsMutex.RLock()
	defer tweetsMutex.RUnlock()

	return tweets
}

func UpdateUserTimeline(c config) {
	client := newTwitterClient(c)

	latestTweets, err := fetchLatestTweets(c.Username, c.TweetCount, client)
	if err != nil {
		log.Printf("Error fetching latest tweets (%v)", err)
	}

	tweetsMutex.Lock()
	tweets = latestTweets
	tweetsMutex.Unlock()

	streamNewTweets(client, func(tweet *twitter.Tweet) {
		tweetsMutex.Lock()
		tweets = append([]twitter.Tweet{*tweet}, tweets[:c.TweetCount-1]...)
		tweetsMutex.Unlock()

		log.Printf("Recieved new tweet: %s\n", tweet.IDStr)
	})

}

func fetchLatestTweets(username string, count int, client *twitter.Client) ([]twitter.Tweet, error) {
	log.Println("Fetching latest tweets from REST API")

	latestTweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      username,
		Count:           count,
		ExcludeReplies:  twitter.Bool(true),
		IncludeRetweets: twitter.Bool(false),
	})

	log.Printf("Fetched %d tweets", len(latestTweets))

	return latestTweets, err
}

func streamNewTweets(client *twitter.Client, tweetHandler tweetHandlerFunc) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = tweetHandler

	log.Println("Starting stream api")

	userParams := &twitter.StreamUserParams{
		StallWarnings: twitter.Bool(true),
		With:          "user",
		Language:      []string{"en"},
	}
	stream, err := client.Streams.User(userParams)
	if err != nil {
		log.Fatal(err)
	}

	demux.HandleChan(stream.Messages)
}

func newTwitterClient(c config) *twitter.Client {
	oauthConfig := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
