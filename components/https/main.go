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

func NewHttpClient(timeout int) *http.Client {
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
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func Request(method string, url string, header map[string]string, query map[string]string, body io.Reader, resultPointer interface{}) (resBytes []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
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
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	resBytes, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if resultPointer == nil {
		return resBytes, nil
	}
	err = jsoniter.Unmarshal(resBytes, resultPointer)
	if err != nil {
		return nil, err
	}
	return resBytes, nil
}
