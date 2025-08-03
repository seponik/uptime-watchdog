# 📈 Uptime Watchdog

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/seponik/uptime-watchdog)](https://goreportcard.com/report/github.com/seponik/uptime-watchdog)

A **lightweight**, **concurrent**, and **reliable** uptime monitoring tool built in Go. Monitor multiple endpoints simultaneously and get instant Slack notifications when services go down.

## ✨ Features

- 🔄 **Concurrent Monitoring** - Check multiple URLs simultaneously for optimal performance

- ⚡ **Lightning Fast** - Built with Go's goroutines for maximum efficiency  

- 📱 **Slack Integration** - Instant notifications to your team channels

- 📋 **YAML Configuration** - Simple, human-readable configuration files

- 🎯 **Minimal Resource Usage** - Designed to be lightweight and efficient

- 📊 **Response Time Tracking** - Monitor performance metrics

## 🚀 Quick Start

### Installation

```bash
# Download the latest release
TODO
```

### Configuration

Create a `config.yaml` file:

```yaml
# Uptime Watchdog Configuration
webhook_url: https://hooks.slack.com/services/X/Y/Z

endpoints:
  - name: Portfolio
    url: https://me.seponik.dev
    timeout: 5s
    interval: 1h

  - name: Example API
    url: https://api.example.com/health
    timeout: 5s
    interval: 30m
    expected_status: 200
    
  - name: Example Website
    url: https://www.example.com
    timeout: 10s
    interval: 1h
    expected_status: 200
```

### Usage

```bash
./uwdog -config YOUR_CONFIG.yaml
```

## 📊 Sample Output

<img src="assets/output.png">


## 🔧 Configuration Options

### Application Configuration

| Parameter | Description | Required |
|-----------|-------------|----------|
| `webhook_url` | Slack webhook URL | ✅ |

### Endpoint Configuration

| Parameter | Description | Default | Required |
|-----------|-------------|---------|----------|
| `name` | Friendly name for the endpoint | - | ✅ |
| `url` | URL to monitor | - | ✅ |
| `interval` | Check interval (e.g., "30s", "2m") | "5m" | ❌ |
| `timeout` | Request timeout | "5s" | ❌ |
| `expected_status` | Expected HTTP status code | 200 | ❌ |



## 🐳 Docker Support

```bash
# Pull the image
docker pull seponik/TODO

# Run with mounted config
docker run -v $(pwd)/config.yaml:/uwdog/config.yaml seponik/TODO
```

### Docker Compose

```yaml
version: '3.8'
services:
  uptime-guardian:
    image: seponik/TODO
    volumes:
      - ./config.yaml:/uwdog/config.yaml
    restart: unless-stopped
```

## 🏗️ Building from Source

```bash
git clone https://github.com/seponik/uptime-watchdog.git

cd uptime-watchdog

make build
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.
## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

