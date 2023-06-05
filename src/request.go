package src

import (
	"io"
	"fmt"
	"net/http"
)

func DoGetRequest(urlRequested string, userAgent string) string {
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
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}