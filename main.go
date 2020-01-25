package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tantalic/envconfig"
)

const (
	Version = "3.0.0"
)

type config struct {
	// Twitter Timeline
	Username   string `envconfig:"USERNAME" required:"true"`
	TweetCount int    `envconfig:"TWEET_COUNT" default:"10"`
	Timeline   string `envconfig:"TIMELINE" default:"home"`

	// API Server
	Host string `envconfig:"HOST" default:""`
	Port int    `envconfig:"PORT" default:"80"`

	// Twitter Credentials
	ConsumerKey    string `envconfig:"CONSUMER_KEY" required:"true"`
	ConsumerSecret string `envconfig:"CONSUMER_SECRET" required:"true"`
	AccessToken    string `envconfig:"ACCESS_TOKEN" required:"true"`
	AccessSecret   string `envconfig:"ACCESS_SECRET" required:"true"`
}

func main() {
	c, err := getConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading configuration from environment (%s).\n", err)
		os.Exit(1)
	}

	switch c.Timeline {
	case "home":
		go UpdateHomeTimeline(c)
	case "user":
		go UpdateUserTimeline(c)
	}

	err = StartApiServer(c)
	if err != nil {
		log.Printf("Unable to start http server: %s\n", err)
	}
}

func getConfig() (config, error) {
	var c config
	err := envconfig.Process("TWITTER_TIMELINE_SERVICE", &c)
	return c, err
}
