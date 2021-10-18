package feed

import "testing"

func TestRSSInterface(t *testing.T) {
	var _ Feed = (*RSS)(nil)
	var _ Feed = &RSS{}
	var _ Feed = new(RSS)
	t.Log("RSS successfully implements Feed")
}
