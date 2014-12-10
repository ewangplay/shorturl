// package sina provides support for t.cn shortening service.
package sina

import (
    "fmt"
	"encoding/json"
	"github.com/ewangplay/shorturl/base"
	"net/http"
)

type Sina struct {
	*base.Service
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

    /*
    新浪短链接服务器返回的数据格式如下:
    {"urls":[
    {"object_type":"","result":true,"url_short":"http://t.cn/RzlP1E0","object_id":"","url_long":"http://github.com/ewangplay/shorturl","type":0}
    ]}
    */
    var result map[string][]map[string]interface{}

	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

    url_short, ok := result["urls"][0]["url_short"].(string)
    if ok {
        return []byte(url_short), nil
    } else {
        return nil, fmt.Errorf("invalid converted short url")
    }
    

}
