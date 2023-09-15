package core

import (
	"io"
	"fmt"
	"bytes"
	"net/http"
)

type Engine struct {
	Response *http.Response
}

func (e *Engine) Get(urlRequested string, userAgent string) *Engine {
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlRequested, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	e.Response = resp
	return e
}

func (e *Engine) Body() string {
	if e.Response != nil {
		defer e.Response.Body.Close()
		body, _ := io.ReadAll(e.Response.Body)
		return string(body)
	}
	return ""
}

func (e *Engine) Cookies() []string {
	cookieVals := make([]string, len(e.Response.Cookies()))
	for i, cookie := range e.Response.Cookies() {
		cookieVals[i] = fmt.Sprintf("%v=%v", cookie.Name, cookie.Value)
	}
	return cookieVals
}

func DoPostRequest(urlRequested string, bodyOfData []byte, addSomeHeaders map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", urlRequested, bytes.NewBuffer(bodyOfData))
	if err != nil {
		return nil, err
	}

	for headerKey, headerValue := range addSomeHeaders {
		req.Header.Set(headerKey, headerValue)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}