package upstage

import (
	"fmt"
	"github.com/eolinker/eosc"
)

type Config struct {
	APIKey       string `json:"upstage_api_key"`
	Organization string `json:"organization"`
}

func checkConfig(v interface{}) (*Config, error) {
	conf, ok := v.(*Config)
	if !ok {
		return nil, eosc.ErrorConfigType
	}
	if conf.APIKey == "" {
		return nil, fmt.Errorf("api_key is required")
	}
	return conf, nil
}