package shorturl

import (
	"errors"
	"github.com/ewangplay/shorturl/adfly"
	"github.com/ewangplay/shorturl/bitly"
	"github.com/ewangplay/shorturl/catchy"
	"github.com/ewangplay/shorturl/cligs"
	"github.com/ewangplay/shorturl/gggg"
	"github.com/ewangplay/shorturl/gitio"
	"github.com/ewangplay/shorturl/googl"
	"github.com/ewangplay/shorturl/isgd"
	"github.com/ewangplay/shorturl/moourl"
	"github.com/ewangplay/shorturl/pendekin"
	"github.com/ewangplay/shorturl/shorl"
	"github.com/ewangplay/shorturl/snipurl"
	"github.com/ewangplay/shorturl/tinyurl"
	"github.com/ewangplay/shorturl/vamu"
	"github.com/ewangplay/shorturl/sina"
	"os"
)

type Client struct {
	Provider string
	Params   map[string]string
}

func NewClient(provider string) *Client {
	return &Client{Provider: provider}
}

func (c *Client) Shorten(u string) ([]byte, error) {
	switch c.Provider {
	case "tinyurl":
		s := tinyurl.New()
		return s.Shorten(u)
	case "isgd":
		s := isgd.New()
		return s.Shorten(u)
	case "gitio":
		s := gitio.New()
		return s.Shorten(u)
	case "bitly":
		s := bitly.New()
		s.Params["login"] = os.Getenv("BITLY_LOGIN")
		s.Params["apiKey"] = os.Getenv("BITLY_API_KEY")
		return s.Shorten(u)
	case "shorl":
		s := shorl.New()
		return s.Shorten(u)
	case "vamu":
		s := vamu.New()
		return s.Shorten(u)
	case "moourl":
		s := moourl.New()
		return s.Shorten(u)
	case "cligs":
		s := cligs.New()
		return s.Shorten(u)
	case "snipurl":
		s := snipurl.New()
		return s.Shorten(u)
	case "adfly":
		s := adfly.New()
		return s.Shorten(u)
	case "googl":
		s := googl.New()
		s.Params["key"] = os.Getenv("GOOGL_API_KEY")
		return s.Shorten(u)
	case "gggg":
		s := gggg.New()
		return s.Shorten(u)
	case "pendekin":
		s := pendekin.New()
		return s.Shorten(u)
	case "catchy":
		s := catchy.New()
		return s.Shorten(u)
    case "sina":
        s := sina.New()
        s.Params["source"] = os.Getenv("SINA_WEIBO_APP_KEY")
        return s.Shorten(u)
	}

	err := errors.New("You should not see this :P")
	return nil, err
}
