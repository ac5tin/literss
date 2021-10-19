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

func (r *RSS) Fetch() error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(r.URL)
	if err != nil {
		return err
	}

	articles := feed.Items[:MAX_ARTICLES]
	for pos := 0; pos < len(articles); pos++ {
		item := feed.Items[len(articles)-pos-1]
		// shift array
		z := *new([]Article)
		z = append(z, r.Articles[0:len(r.Articles)-1]...)
		// push article to array
		images := new([]string)
		if item.Image != nil {
			*images = append(*images, item.Image.URL)
		}

		r.Articles[0] = Article{
			ID:      item.GUID,
			Title:   item.Title,
			Content: item.Content,
			Date:    *item.PublishedParsed,
			URL:     item.Link,
			Images:  *images,
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
