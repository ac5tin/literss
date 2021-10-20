package feed

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func TestRSSInterface(t *testing.T) {
	var _ Feed = (*RSS)(nil)
	var _ Feed = &RSS{}
	var _ Feed = new(RSS)
	t.Log("RSS successfully implements Feed")
}

func TestArrayShift(t *testing.T) {
	a := *new([3]string)

	a[0] = "first"
	a[1] = "second"
	a[2] = "third"

	b := []string{"fourth", "fifth"}

	for _, it := range b {
		z := *new([]string)
		z = append(z, a[0:len(a)-1]...)
		a[0] = it
		for i, y := range z {
			a[i+1] = y
		}

	}

	log.Printf("final : %v", a)
	if reflect.DeepEqual(a, [3]string{"fifth", "fourth", "first"}) {
		t.Log("Array shift implemented correctly")
		return
	}
	t.Errorf("array shift not equal")
}

func TestRSSFetchGet(t *testing.T) {
	a := NewRSS("http://feeds.bbci.co.uk/news/world/rss.xml", "BBC")
	if err := a.Fetch(); err != nil {
		t.Errorf(err.Error())
	}

	v := new([]Article)
	if err := a.Get(v); err != nil {
		t.Errorf(err.Error())
	}

	if err := a.AutoFetch(); err != nil {
		t.Errorf(err.Error())
	}

	log.Println("ending RSS Feed in 10 seconds")
	time.Sleep(10 * time.Second)
	log.Println("ending RSS Feed now")
	a.cc()

	time.Sleep(5 * time.Second)
	log.Println("Ended RSS")

	return
}
