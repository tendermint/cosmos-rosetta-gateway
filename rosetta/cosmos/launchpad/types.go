package launchpad

import "net/http"

type Launchpad struct {
	endpoint string
	c        *http.Client
}

func NewLaunchpad(c *http.Client, endpoint string) *Launchpad {
	return &Launchpad{
		c:        c,
		endpoint: endpoint,
	}
}
