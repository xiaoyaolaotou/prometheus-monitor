package global

import (
	"monitor/config"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	PromConfig *config.PromConfig
	APIs       map[string]v1.API
)
