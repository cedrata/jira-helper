package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RequestHelper struct {
	host            string
	resource        string
	protocol        string
	method          string
	queryParameters map[string]string
	headers         map[string]string
	body            []byte
}

func NewRequestHelper(host string, resource string, method string, queryParameters map[string]string, headers map[string]string, body []byte) *RequestHelper {
	return &RequestHelper{
		host:            host,
		resource:        resource,
		protocol:        "https",
		method:          method,
		queryParameters: queryParameters,
		headers:         headers,
		body:            body,
	}
}

func (rh *RequestHelper) BuildRequest() (*http.Request, error) {

	parsedResource, err := url.Parse(rh.resource)
	if err != nil {
		return nil, err
	}

	queryString := parsedResource.Query()
	for k, v := range rh.queryParameters {
		queryString.Set(k, v)
	}

	parsedResource.RawQuery = queryString.Encode()
	parsedHost, err := url.Parse(fmt.Sprintf("%s://%s", rh.protocol, rh.host))
	if err != nil {
		return nil, err
	}

	completeUrl := parsedHost.ResolveReference(parsedResource)
	bodyBuffer := bytes.NewBuffer(rh.body)

	request, err := http.NewRequest(rh.method, completeUrl.String(), bodyBuffer)
	if err != nil {
		return nil, err
	}

	for k, v := range rh.headers {
		request.Header.Add(k, v)
	}

	return request, nil
}

func PostHeadersWithBearer(token string) map[string]string {
	return map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
}

func GetHeadersWithBearer(token string) map[string]string {
	return map[string]string{
		"Accept":        "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
}

func PrettyHttpReponse(response *http.Response) (string, error) {
	var buf bytes.Buffer
    var err error

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	minified := fmt.Sprintf(
		`{"statusCode":%d,"body":%s}`,
		response.StatusCode, body,
	)

	err = json.Indent(&buf, []byte(minified), "", "  ")

	return buf.String(), err
}
