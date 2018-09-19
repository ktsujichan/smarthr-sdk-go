package smarthr

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

type Config struct {
	Name    string `yaml:"name"`
	Token   string `yaml:"token"`
	Sandbox bool   `yaml:"sandbox"`
}

func (c *Config) host() string {
	if c.Sandbox {
		return "daruma.space"
	} else {
		return "smarthr.jp"
	}
}

func (c *Config) URL() (*url.URL, error) {
	endpoint := fmt.Sprintf("https://%s.%s/api", c.Name, c.host())
	u, err := url.ParseRequestURI(endpoint)
	return u, errors.WithStack(err)
}
