package simplehttptestwrapper

import (
	"net/http"
	"net/http/httptest"
)

// HTTPTestRequestConfig allows for customizing an http
// test request
type HTTPTestRequestConfig struct {
	URLQueryParameters []map[string]string
	Headers            []map[string]string
}

// HTTPTestResponseFormat defines what a test response is
type HTTPTestResponseFormat struct {
	Code int
	Body []byte
}

// HTTPTestHandler returns the tested handlers response to be
func HTTPTestHandler(customHandler http.HandlerFunc, requestType string, endpoint string, conf HTTPTestRequestConfig) (HTTPTestResponseFormat, error) {
	req, err := http.NewRequest(requestType, endpoint, nil)
	if err != nil {
		return HTTPTestResponseFormat{}, err
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customHandler)
	handler.ServeHTTP(rr, req)

	ret := HTTPTestResponseFormat{
		Code: rr.Code,
		Body: rr.Body.Bytes(),
	}
	return ret, nil
}
