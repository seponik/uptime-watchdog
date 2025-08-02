package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newTestServer(status int, delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(status)
	}))
}

func TestCheckAll(t *testing.T) {
	upSrv := newTestServer(200, 50*time.Millisecond)
	defer upSrv.Close()

	warnSrv := newTestServer(301, 20*time.Millisecond)
	defer warnSrv.Close()

	downURL := "http://127.0.0.1:65500"

	urls := []string{upSrv.URL, warnSrv.URL, downURL}
	results := CheckAll(urls)

	if len(results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(results))
	}

	if results[0].StatusCode != 200 || results[0].Error != nil {
		t.Errorf("expected UP result, got %+v", results[0])
	}

	if results[1].StatusCode != 301 || results[1].Error != nil {
		t.Errorf("expected WARN result, got %+v", results[1])
	}

	if results[2].Error == nil {
		t.Errorf("expected error for DOWN, got %+v", results[2])
	}
}
