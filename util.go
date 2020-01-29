package slackoff

import (
	"fmt"
	"net/http"
	"github.com/pkg/errors"
	"bytes"
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
)

var ErrorColor = color.New(color.FgWhite, color.BgRed, color.Bold)

// StatusCodeError represents an http response error.
// type httpStatusCode interface { HTTPStatusCode() int } to handle it.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("HTTP Response error: %s.", t.Status)
}

func (t statusCodeError) HTTPStatusCode() int {
	return t.Code
}

type HttpPoster interface {
	Post(string, interface{}) (*http.Response, error)
}

type HttpClient struct {}

func (h *HttpClient) Post(url string, jsonPayload interface{}) (*http.Response, error) {
	raw, err := json.Marshal(jsonPayload)

	if err != nil {
		return nil, errors.Wrap(err, "marshal failed")
	}

	response, err := http.Post(url, "application/json", bytes.NewReader(raw));

	if err != nil {
		return response, errors.Wrap(err, "failed to post")
	}

	if response.StatusCode != http.StatusOK {
		return response, statusCodeError{Code: response.StatusCode, Status: response.Status}
	}

	return response, nil
}

// Interface assertion
var _ HttpPoster = (*HttpClient)(nil)

type FileReader interface {
	ReadFile(filename string) (string, error)
}

type IOFileReader struct {}
func (i *IOFileReader) ReadFile(filename string) (string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Interface assertion
var _ FileReader = (*IOFileReader)(nil)
