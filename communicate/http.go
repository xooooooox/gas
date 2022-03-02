package communicate

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

// NewRequest full request
func NewRequest(reqMethod string, reqUrl string, reqHeader map[string]string, reqBody []byte) (resStatusCode int, resHeader http.Header, resBody []byte, err error) {
	var request *http.Request
	request, err = http.NewRequest(reqMethod, reqUrl, bytes.NewReader(reqBody))
	if err != nil {
		return
	}
	length := len(reqHeader)
	if length > 0 {
		for k, v := range reqHeader {
			request.Header.Set(k, v)
		}
	}
	client := &http.Client{
		Timeout: time.Second * 15,
	}
	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	resBody, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	resStatusCode = response.StatusCode
	resHeader = response.Header
	return
}

// GetHeader get request with response header
func GetHeader(reqUrl string, reqHeader ...map[string]string) (resHeader http.Header, resBody []byte, err error) {
	var header map[string]string
	length := len(reqHeader)
	if length > 0 {
		header = reqHeader[length-1]
	}
	_, resHeader, resBody, err = NewRequest(http.MethodGet, reqUrl, header, nil)
	return
}

// PostHeader post request with response header
func PostHeader(reqUrl string, reqBody []byte, reqHeader ...map[string]string) (resHeader http.Header, resBody []byte, err error) {
	var header map[string]string
	length := len(reqHeader)
	if length > 0 {
		header = reqHeader[length-1]
	}
	_, resHeader, resBody, err = NewRequest(http.MethodPost, reqUrl, header, reqBody)
	return
}

// Get get request
func Get(reqUrl string, reqHeader ...map[string]string) (resBody []byte, err error) {
	_, resBody, err = GetHeader(reqUrl, reqHeader...)
	return
}

// Post post request
func Post(reqUrl string, reqBody []byte, reqHeader ...map[string]string) (resBody []byte, err error) {
	_, resBody, err = PostHeader(reqUrl, reqBody, reqHeader...)
	return
}
