package request

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Request struct {
	URL         string
	Body        string
	RequestType string //GET/POST
}

func New() *Request {
	r := &Request{}
	return r
}

func (r *Request) SetURL(url string) *Request {
	r.URL = url
	return r
}

func (r *Request) SetBody(body string) *Request {
	r.Body = body
	return r
}

func (r *Request) SetRequestType(rtype string) *Request {
	r.RequestType = rtype
	return r
}

func (r *Request) Execute() ([]byte, error) {
	payload := strings.NewReader(r.Body)
	req, err := http.NewRequest(r.RequestType, r.URL, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)

}
