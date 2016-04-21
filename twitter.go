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

func UpdateUserTimeline(c config) {
	client := newTwitterClient(c)

	latestTweets, err := fetchLatestUserTweets(c.Username, c.TweetCount, client)
	if err != nil {
		log.Printf("Error fetching latest tweets (%v)", err)
	}

	appendTweets(latestTweets, c.TweetCount)

	streamNewTweets(client, "user", func(tweet *twitter.Tweet) {
		log.Printf("Recieved new tweet: %s\n", tweet.IDStr)
		appendTweets([]twitter.Tweet{*tweet}, c.TweetCount)
	})
}

func UpdateHomeTimeline(c config) {
	client := newTwitterClient(c)

	latestTweets, err := fetchLatestHomeTweets(c.TweetCount, client)
	if err != nil {
		log.Printf("Error fetching latest tweets (%v)", err)
	}

	appendTweets(latestTweets, c.TweetCount)

	streamNewTweets(client, "followings", func(tweet *twitter.Tweet) {
		log.Printf("Recieved new tweet: %s\n", tweet.IDStr)
		appendTweets([]twitter.Tweet{*tweet}, c.TweetCount)
	})
}

func GetTweets() []twitter.Tweet {
	tweetsMutex.RLock()
	defer tweetsMutex.RUnlock()

	return tweets
}

func fetchLatestUserTweets(username string, count int, client *twitter.Client) ([]twitter.Tweet, error) {
	log.Println("Fetching latest user tweets from REST API")

	latestTweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName:      username,
		Count:           count,
		ExcludeReplies:  twitter.Bool(true),
		IncludeRetweets: twitter.Bool(false),
	})

	log.Printf("Fetched %d tweets", len(latestTweets))

	return latestTweets, err
}

func fetchLatestHomeTweets(count int, client *twitter.Client) ([]twitter.Tweet, error) {
	log.Println("Fetching latest home tweets from REST API")

	latestTweets, _, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: count,
	})

	log.Printf("Fetched %d tweets", len(latestTweets))

	return latestTweets, err
}

func streamNewTweets(client *twitter.Client, with string, tweetHandler tweetHandlerFunc) {
	demux := twitter.NewSwitchDemux()
	demux.Tweet = tweetHandler

	log.Println("Starting stream api")

	userParams := &twitter.StreamUserParams{
		StallWarnings: twitter.Bool(true),
		With:          with,
		Language:      []string{"en"},
	}
	stream, err := client.Streams.User(userParams)
	if err != nil {
		log.Fatal(err)
	}

	demux.HandleChan(stream.Messages)
}

func appendTweets(newTweets []twitter.Tweet, maxLength int) {
	tweetsMutex.Lock()
	defer tweetsMutex.Unlock()

	tweets = append(newTweets, tweets...)
	tweets = tweets[:maxLength]
}

func newTwitterClient(c config) *twitter.Client {
	oauthConfig := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := oauthConfig.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}
