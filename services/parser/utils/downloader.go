package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// DownloadFile - file from endpoint via http and save it to path
func DownloadFile(path string, endpoint string) error {

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	c := http.DefaultClient
	c.Timeout = time.Second * 4
	// Make request
	to, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	req := http.Request{
		Method: http.MethodGet,
		URL:    to,
	}
	resp, err := c.Do(&req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
