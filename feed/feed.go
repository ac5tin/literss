package feed

import (
	"fmt"
	"time"

	uf "github.com/ac5tin/usefulgo"
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
	Fetch(num uint8, t *[]Article) error
	Get(t *[]Article) error
}

type feedBase struct {
	ID       string
	URL      string
	Name     string
	Articles []Article
	Feed
}

func NewFeedBase(url, name string) feedBase {
	return feedBase{
		ID:       uf.GenUUIDV4(),
		URL:      url,
		Name:     name,
		Articles: make([]Article, 0),
	}
}

func (f *feedBase) Fetch(num uint8, t *[]Article) error {
	return fmt.Errorf("Fetch() not implemented")
}

func (f *feedBase) Get(t *[]Article) error {
	return fmt.Errorf("Get() not implemented")
}
