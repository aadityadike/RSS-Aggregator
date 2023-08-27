package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string     `xml:"title"`
		Description string     `xml:"description"`
		Link        string     `xml:"link"`
		Language    string     `xml:"language"`
		Items       []RssItems `xml:"item"`
	} `xml:"channel"`
}

type RssItems struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Creator     string `xml:"dc:creator"`
	PubDate     string `xml:"pubDate"`
	Cover_image string `xml:"cover_image"`
}

func getAllFeeds(url string) (RssFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RssFeed{}, err
	}

	storeData := RssFeed{}

	err = xml.Unmarshal(data, &storeData)
	if err != nil {
		return RssFeed{}, err
	}

	return storeData, nil
}
