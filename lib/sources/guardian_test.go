package sources

import "testing"

func TestGetNews(t *testing.T) {
	want := ""
	source := Guardian{}
	if got := source.GetNews(); got != want {
		t.Errorf("Guardian: want %q, got %q", want, got)
	}
}
