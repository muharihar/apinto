package service_http

import (
	"strings"

	"github.com/eolinker/goku/plugin"

	"github.com/eolinker/eosc"
)

type AnonymousConfig struct {
	Type   string `json:"type"`
	Config string `json:"config"`
}

//Config service_http驱动配置
type Config struct {
	id          string
	Name        string `json:"name"`
	Driver      string `json:"driver"`
	Desc        string `json:"desc"`
	Timeout     int64  `json:"timeout"`
	Retry       int    `json:"retry"`
	Scheme      string `json:"scheme"`
	ProxyMethod string `json:"proxy_method"`

	Upstream          eosc.RequireId   `json:"upstream" skill:"github.com/eolinker/goku/upstream.upstream.IUpstream" require:"false"`
	UpstreamAnonymous *AnonymousConfig `json:"anonymous"`

	PluginConfig map[string]*plugin.Config `json:"plugins"`
}

var validMethods = []string{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
	"HEAD",
	"OPTIONS",
}

var validScheme = []string{
	"HTTP",
	"HTTPS",
}

func (c *Config) rebuild() {
	if c.Retry < 0 {
		c.Retry = 0
	}
	if c.Timeout < 0 {
		c.Timeout = 0
	}
	if !checkValidParams(strings.ToUpper(c.ProxyMethod), validMethods) {
		c.ProxyMethod = ""
	}

	if !checkValidParams(strings.ToUpper(c.Scheme), validScheme) {
		c.Scheme = "http"
	}
}

func checkValidParams(data string, params []string) bool {
	for _, p := range params {
		if data == p {
			return true
		}
	}
	return false
}
