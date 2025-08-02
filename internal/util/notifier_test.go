package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newWebhookServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload map[string]string

		json.NewDecoder(r.Body).Decode(&payload)

		if payload["text"] == "" {
			http.Error(w, "missing text", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return httptest.NewServer(handler)
}

func TestSendAlert(t *testing.T) {
	server := newWebhookServer()
	defer server.Close()

	err := sendAlert(server.URL, "http://test.org")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
