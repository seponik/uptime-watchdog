package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/seponik/uptime-watchdog/internal/config"
	"github.com/seponik/uptime-watchdog/internal/monitor"
	"github.com/seponik/uptime-watchdog/internal/util"
)

func main() {
	cfgPath := flag.String("config", "/uwdog/config.yaml", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Println("[ERR] Failed to find/read the config file.")
		os.Exit(1)
	}

	err = util.ValidateConfig(cfg)
	if err != nil {
		fmt.Printf("[ERR] Configuration validation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[INFO] Starting uptime watchdog...")

	fmt.Println("[INFO] Loaded configuration successfully.")

	notifier := util.NewNotifier(cfg.WebhookURL)

	var wg sync.WaitGroup

	for _, endpoint := range cfg.Endpoints {
		wg.Add(1)

		go func(ep config.Endpoint, nt util.Notifier) {
			defer wg.Done()

			monitor.Monitor(ep, nt)
		}(endpoint, notifier)
	}

	wg.Wait()
}
