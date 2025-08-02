package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/seponik/uptime-watchdog/internal/checker"
	"github.com/seponik/uptime-watchdog/internal/config"
	"github.com/seponik/uptime-watchdog/internal/util"
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Println("[ERR] Failed to read the config file.")
		os.Exit(1)
	}

	for {
		results := checker.CheckAll(config.URLs)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		util.ProcessResults(results, config.WebhookURL)

		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}
