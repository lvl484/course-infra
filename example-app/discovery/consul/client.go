package consul

import consul "github.com/hashicorp/consul/api"

// NewClient new consul client
func NewClient(address string) (*consul.Client, error) {
	consulCfg := consul.DefaultConfig()
	consulCfg.Address = address

	return consul.NewClient(consulCfg)
}
