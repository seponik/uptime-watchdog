package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/seponik/uptime-watchdog/internal/config"
)

func newTestServer(status int, delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(status)
	}))
}

func TestChecker(t *testing.T) {
	upSrv := newTestServer(200, 50*time.Millisecond)
	defer upSrv.Close()

	downSrv := newTestServer(500, 3*time.Second)
	defer downSrv.Close()

	endpointUp := config.Endpoint{
		URL:     upSrv.URL,
		Timeout: config.Duration(2 * time.Second),
	}

	endpointDown := config.Endpoint{
		URL:     downSrv.URL,
		Timeout: config.Duration(2 * time.Second),
	}

	checkerUp := NewChecker(endpointUp)
	resultUp := checkerUp.Check()

	checkerDown := NewChecker(endpointDown)
	resultDown := checkerDown.Check()

	if resultUp.StatusCode != 200 || resultUp.Error != nil {
		t.Errorf("expected UP result, got %+v", resultUp)
	}

	if resultDown.Error == nil {
		t.Errorf("expected DOWN result, got %+v", resultDown)
	}
}
