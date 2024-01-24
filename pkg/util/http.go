package util

import (
	"crypto/tls"
	"io"
	"net/http"
)

func Request(url string, headers map[string][]string, method string, sslVerify bool) ([]byte, error) {
	// new request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// set request headers
	req.Header = headers

	// raise GET request
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: sslVerify},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// read the data of response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
