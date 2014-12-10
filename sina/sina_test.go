package sina

import (
	"os"
	"testing"
)

var Url string = "http://github.com/ewangplay/shorturl"

func TestSina(t *testing.T) {
	s := New()
	s.Params["source"] = os.Getenv("SINA_WEIBO_APP_KEY")
	//s.Params["source"] = "4228546392"
	u, err := s.Shorten(Url)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%s => %s\n", Url, u)
}
