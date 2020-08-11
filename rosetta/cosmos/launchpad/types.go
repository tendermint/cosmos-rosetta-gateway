package launchpad

type Launchpad struct {
	endpoint string
}

func NewLaunchpad(endpoint string) *Launchpad {
	return &Launchpad{endpoint: endpoint}
}
