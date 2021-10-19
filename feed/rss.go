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
		// shift array
		z := *new([]Article)
		z = append(z, r.Articles[0:len(r.Articles)-1]...)
		// push article to array
		r.Articles[0] = Article{
			ID:      item.GUID,
			Title:   item.Title,
			Content: item.Content,
			Date:    *item.PublishedParsed,
			URL:     item.Link,
			Images:  []string{item.Image.URL},
			FeedID:  r.ID,
		}
		// fill rest of array
		for i, y := range z {
			r.Articles[i+1] = y
		}
	}
	return nil
}

func (r *RSS) Get(t *[]Article) error {
	*t = r.Articles[:]
	return nil
}
