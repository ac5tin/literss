package feed

import (
	"github.com/mmcdole/gofeed"
)

type RSS struct {
	*feedBase
}

func NewRSS(url, name string) RSS {
	fd := NewFeedBase(url, name)
	return RSS{
		&fd,
	}
}

func (r *RSS) Fetch(num uint8, t *[]Article) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(r.URL)
	if err != nil {
		return err
	}
	for _, item := range feed.Items {
		r.Articles = append(r.Articles, Article{
			ID:      item.GUID,
			Title:   item.Title,
			Content: item.Content,
			Date:    *item.PublishedParsed,
			URL:     item.Link,
			FeedID:  r.ID,
		})
	}
	return nil
}
