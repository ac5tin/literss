package feed

import (
	"fmt"
	"time"

	uf "github.com/ac5tin/usefulgo"
)

const MAX_ARTICLES uint8 = 15

type FeedType string

const (
	FeedTypeRSS FeedType = "RSS"
)

type Article struct {
	ID      string
	Title   string
	Content string
	Date    time.Time
	URL     string
	FeedID  string
	Images  []string
}

type Feed interface {
	Fetch() error
	Get(t *[]Article) error
	AutoFetch() error
}

type feedBase struct {
	ID       string
	URL      string
	Name     string
	Articles [MAX_ARTICLES]Article
	Feed
}

func NewFeedBase(url, name string) feedBase {
	return feedBase{
		ID:       uf.GenUUIDV4(),
		URL:      url,
		Name:     name,
		Articles: *new([MAX_ARTICLES]Article),
	}
}

func (f *feedBase) Fetch() error {
	return fmt.Errorf("Fetch() not implemented")
}

func (f *feedBase) Get(t *[]Article) error {
	return fmt.Errorf("Get() not implemented")
}

func (f *feedBase) AutoFetch() error {
	return fmt.Errorf("AutoFetch() not implemented")
}
