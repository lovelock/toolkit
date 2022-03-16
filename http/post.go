package http

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/lovelock/toolkit/handlers"
)

func PostForm(url string, params map[string]interface{}, form map[string]interface{}) (*resty.Response, error) {
	client := resty.New()
	r := client.R()
	if params != nil && len(params) > 0 {
		queryParams := make(map[string]string)
		for k, v := range params {
			queryParams[k] = v.(string)
		}
		r.SetQueryParams(queryParams)
	}

	if form != nil && len(form) > 0 {
		formData := make(map[string]string)
		for k, v := range form {
			formData[k] = v.(string)
		}

		r.SetFormData(formData)
	}

	return r.Post(url)
}

func PostJson(url string, params map[string]interface{}, jsonBody interface{}) (*resty.Response, error) {
	client := resty.New()
	r := client.R()
	if params != nil && len(params) > 0 {
		queryParams := make(map[string]string)
		for k, v := range params {
			queryParams[k] = v.(string)
		}
		r.SetQueryParams(queryParams)
	}

	if jsonBody != nil {
		_, err := json.Marshal(jsonBody)
		if err != nil {
			handlers.LogErrorAndGoon(err)
		} else {
			r.SetBody(jsonBody)
		}
	}

	return r.Post(url)
}
