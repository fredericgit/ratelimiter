package ratelimiter

import (
	"fmt"
	"time"
	"strconv"
)

// Config represents the processor config settings within the collector's config.yaml
type Config struct {
	LimitRate    string `mapstructure:"limit_rate"`
	LimitBy      []string `mapstructure:"limit_by"`
}

// Validate checks if the processor configuration is valid
func (cfg *Config) Validate() error {	
	rate = strings.Replace(cfg.LimitRate, " ", "", -1)
	nbOfTraces, interval, found := strings.Cut(rate, "/")
	if !found {
		return fmt.Errorf("invalide rate value. Should be in the form '<integer>/<duration>' (ex: '10000/5m')")
	}
	_, err := time.ParseDuration(interval)
	if err != nil {
		return fmt.Errorf("invalide interval in rate value")
	}
	_, err := strconv.ParseInt(nbOfTraces, 10, 64)
	if err != nil {
		return fmt.Errorf("invalide number of traces in rate value")
	}
	return nil
}
