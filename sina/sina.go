// package sina provides support for t.cn shortening service.
package sina

import (
	"encoding/json"
	"github.com/ewangplay/shorturl/base"
	"net/http"
)

type Sina struct {
	*base.Service
}

type response struct {
	UrlShort    string
	UrlLong     string
    Type        int
	Result      bool
}

func New() *Sina {
	return &Sina{&base.Service{
		Scheme: "https",
		Host:   "api.weibo.com",
		Method: "GET",
		Path:   "2/short_url/shorten.json",
		Field:  "url_long",
		Code:   http.StatusOK,
		Params: map[string]string{
			"source": "",
		},
	}}
}

func (s *Sina) Shorten(u string) ([]byte, error) {
	res, err := s.Request(u)
	if err != nil {
		return nil, err
	}

	b, err := s.Read(res)
	if err != nil {
		return nil, err
	}

	r := &response{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}

	return []byte(r.UrlShort), nil
}
