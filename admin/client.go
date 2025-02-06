package admin

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type AdminClient struct {
	apiUrl string
	token  string
}

func NewAdminClient(apiUrl string, token string) *AdminClient {
	return &AdminClient{
		apiUrl: apiUrl,
		token:  token,
	}
}

func (s *AdminClient) send(method string, path string, params *map[string]string, body *[]byte) (string, error) {
	var bodyBuffer *bytes.Buffer
	if body == nil {
		bodyBuffer = bytes.NewBuffer([]byte{})
	} else {
		bodyBuffer = bytes.NewBuffer(*body)
	}
	requestUrl, _ := url.JoinPath(s.apiUrl, path)
	request, _ := http.NewRequest(method, requestUrl, bodyBuffer)
	request.Header.Set("Authorization", fmt.Sprintf("Token %s", s.token))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	if params != nil {
		q := request.URL.Query()
		for key, value := range *params {
			q.Add(key, value)
		}
		request.URL.RawQuery = q.Encode()
	}

	httpClient := http.Client{Timeout: 30 * time.Second}

	res, err := httpClient.Do(request)

	if err != nil {
		return "", err
	}

	if res.StatusCode != 200 {
		return "", &InvalidResponseError{res.Status, res.StatusCode}
	}

	defer res.Body.Close()

	response, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(response), nil
}
