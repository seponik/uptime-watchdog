package monitor

import (
	"time"

	"github.com/seponik/uptime-watchdog/internal/checker"
	"github.com/seponik/uptime-watchdog/internal/config"
	"github.com/seponik/uptime-watchdog/internal/util"
)

func Monitor(endpoint config.Endpoint, notifier util.Notifier) {

	if endpoint.Interval <= 0 {
		endpoint.Interval = config.Duration(time.Minute * 5) // Default to 5 min if not set
	}

	if endpoint.Timeout <= 0 {
		endpoint.Timeout = config.Duration(time.Second * 5) // Default to 5 seconds if not set
	}

	if endpoint.ExpectedStatus <= 0 {
		endpoint.ExpectedStatus = 200 // Default to 200 if not set
	}

	checker := checker.NewChecker(endpoint)

	for {
		result := checker.Check()

		notifier.ProcessResult(result)

		time.Sleep(time.Duration(endpoint.Interval))
	}
}
