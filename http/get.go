package http

import (
	"github.com/go-resty/resty/v2"
)

func Get(url string, params map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	r := client.R()
	if params != nil {
		queryParams := make(map[string]string)
		for k, v := range params {
			queryParams[k] = v.(string)
		}
		r.SetQueryParams(queryParams)
	}

	return r.Get(url)
}
