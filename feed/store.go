package feed

import (
	"crypto/sha512"
	"fmt"
	"log"
	"sync"
	"time"
)

type FeedStore struct {
	Lock  sync.Mutex
	Feeds map[string]Feed
}

func NewFeedStore() FeedStore {
	return FeedStore{
		Feeds: make(map[string]Feed),
	}
}

func (f *FeedStore) UpdateFeed() error {
	for _, v := range f.Feeds {
		go func(feed Feed) {
			if err := feed.Fetch(); err != nil {
				log.Printf("Error: %s", err.Error())
			}
		}(v)
	}
	return nil
}

func (f *FeedStore) AddRSSFeed(rssURL, name string) error {
	hash := fmt.Sprintf("%x", sha512.Sum512([]byte(rssURL)))

	// check if already exist, return if true
	if _, ok := f.Feeds[hash]; ok {
		return nil
	}
	rssFeed := NewRSS(rssURL, name)
	// fetch feeds in the background
	go rssFeed.Fetch()
	// auto fetch
	go rssFeed.AutoFetch()
	// assign feed to store
	f.Lock.Lock()
	f.Feeds[string(hash)] = &rssFeed
	f.Lock.Unlock()
	return nil
}

func (f *FeedStore) GetRSSFeed(rssURL, name string) (Feed, error) {
	hash := fmt.Sprintf("%x", sha512.Sum512([]byte(rssURL)))

	// check if already exist, return if true
	if feed, ok := f.Feeds[hash]; ok {
		return feed, nil
	}

	// create new if feed doesnt exist
	rssFeed := NewRSS(rssURL, name)
	// fetch feeds in the background
	if err := rssFeed.Fetch(); err != nil {
		return nil, err
	}
	// assign feed to store
	f.Lock.Lock()
	f.Feeds[string(hash)] = &rssFeed
	f.Lock.Unlock()

	return &rssFeed, nil
}

func (f *FeedStore) AutoReport() {
	for {
		log.Println("============================\nReport:\n----------")
		log.Printf("Current Timestamp: %v", time.Now())
		log.Printf("Total feeds: %d\n", len(f.Feeds))
		log.Println("============================")
		time.Sleep(15 * time.Minute)
	}
}

var FS *FeedStore
