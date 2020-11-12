package https

import (
	"crypto/tls"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives:   false,
		Proxy:               http.ProxyFromEnvironment,
		MaxIdleConns:        30 * 2,
		MaxConnsPerHost:     30,
		MaxIdleConnsPerHost: 30,
		IdleConnTimeout:     30 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	},
	Timeout: 30 * time.Second,
}

func NewHttpClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			Proxy:               http.ProxyFromEnvironment,
			MaxIdleConns:        30 * 2,
			MaxConnsPerHost:     30,
			MaxIdleConnsPerHost: 30,
			IdleConnTimeout:     60 * time.Second,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: timeout,
	}
}

func Request(method string, url string, header map[string]string, query map[string]string, body io.Reader, resultPointer interface{}) (resBytes []byte, res *http.Response, err error) {
	return RequestByClient(httpClient, method, url, header, query, body, resultPointer)
}

func RequestByClient(httpClient *http.Client, method string, url string, header map[string]string, query map[string]string, body io.Reader, resultPointer interface{}) (resBytes []byte, res *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}
	//header
	if header != nil && len(header) > 0 {
		for k, v := range header {
			mh := textproto.MIMEHeader(req.Header)
			mh[k] = []string{v}
		}
	}
	//query
	if query != nil && len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	//send
	res, err = httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	resBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}
	if resultPointer == nil {
		return resBytes, res, nil
	}
	err = jsoniter.Unmarshal(resBytes, resultPointer)
	if err != nil {
		return nil, res, err
	}
	return resBytes, res, nil
}
