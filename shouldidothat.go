package main

import (
  "fmt"
  "net/url"
  "log"
  "time"
  "github.com/darkhelmet/twitterstream"
  "github.com/ChimeraCoder/anaconda"
  "os"
  "encoding/json"
)

func decode(conn *twitterstream.Connection) {
  api := anaconda.NewTwitterApi(conf.AccessToken, conf.AccessTokenSecret)
  for {
    if tweet, err := conn.Next(); err == nil {
      log.Printf("%s said: %s\n", tweet.User.ScreenName, tweet.Text)
      response := fmt.Sprintf("@%s yeah, why not?", tweet.User.ScreenName)
      log.Printf("replying: %s", response)
      opts := url.Values{}
      opts.Set("in_reply_to_status_id_str", string(tweet.Id))
      api.PostTweet(response, opts)
    } else {
      log.Printf("Failed decoding tweet: %s", err)
      return
    }
  }
}

type Configuration struct {
  ApiKey string
  ApiSecret string
  AccessToken string
  AccessTokenSecret string
}

var conf = Configuration{}

func main() {
  file, _ := os.Open("conf.json")
  decoder := json.NewDecoder(file)
  err := decoder.Decode(&conf)
  if err != nil {
    fmt.Println("error loading config:", err)
    os.Exit(1)
  }

  anaconda.SetConsumerKey(conf.ApiKey)
  anaconda.SetConsumerSecret(conf.AccessTokenSecret)
  streaming_client := twitterstream.NewClient(conf.ApiKey, conf.ApiSecret, conf.AccessToken, conf.AccessTokenSecret)
  streaming_client.Timeout = 0

  for {
    log.Println("tracking...")
    conn, err := streaming_client.Track("@shouldidothat")
    if err != nil {
      log.Println(err)
      log.Println("Tracking failed, sleeping for 1 minute")
      time.Sleep(1 * time.Minute)
      continue
    }
    decode(conn)
  }
}

