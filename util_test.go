// +build unit

package slackoff

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
)

func TestPostWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()
	h := HttpClient{}
	r, err := h.Post(ts.URL, &WebhookMessage{
		Text: "My Test",
	})
	if err == nil {
		t.Errorf("Post() didn’t return an error")
	}

	if r == nil {
		t.Errorf("Post() didn't return a response")
	}
}

func TestPostBodyIsJson(t *testing.T) {
	var receivedPayload WebhookMessage

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&receivedPayload)
		if err != nil {
			t.Errorf("Request contained invalid JSON, %s", err)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	h := HttpClient{}
	r, err := h.Post(ts.URL, &WebhookMessage{
		Text: "My Test",
	})
	if err != nil {
		t.Errorf("Post() returned an error")
	}

	if r == nil {
		t.Errorf("Post() didn't return a response")
	}
}

func TestNoResponse(t *testing.T) {

	h := HttpClient{}
	_, err := h.Post("http://foo.invalid", &WebhookMessage{
		Text: "My Test",
	})

	if err == nil {
		t.Errorf("Post() didn't return an error")
	}
}
