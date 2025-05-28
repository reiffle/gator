package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	//"os"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	postReq, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil) /*
		create a new request, we will still need  to actually send the request
		"nil" is because we are requesting info, not sending it
	*/
	if err != nil {
		fmt.Println("couldn't fetch feed")
		return nil, err
	}

	postReq.Header.Set("User-Agent", "gator") /*
		tells server which program or user is makeing a request.  May get
		different results from different identifiers
	*/

	resp, err := http.DefaultClient.Do(postReq)
	// actually send the request to the server

	if err != nil {
		fmt.Println("coudln't process request")
		return nil, err
	}

	defer resp.Body.Close() /*now we're going to open the file and read
	the response, so we need to close it when we're done*/

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("couldn't read response")
		return nil, err
	}

	//parse the response into the go struct
	var feed RSSFeed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Error umarshalling XML")
		return nil, err
	}

	//format the raw html string for the given fields
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	//do the same for the struct map inside the main struct
	for index, _ := range feed.Channel.Item {
		curr_item := &feed.Channel.Item[index]
		curr_item.Title = html.UnescapeString(curr_item.Title)
		curr_item.Description = html.UnescapeString((curr_item.Description))
	}

	return &feed, nil
}
