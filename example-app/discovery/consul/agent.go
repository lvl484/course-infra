package consul

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
	"net/http"
)

const (
	agentCheckInterval   = 10
	agentCheckTimeout    = 10
	agentDeregisterAfter = 60
)

// AgentConfig is a new consul agent config
func AgentConfig(serviceName string, bindPort int, healthCheck string) *consul.AgentServiceRegistration {
	return &consul.AgentServiceRegistration{
		Name: serviceName,
		Check: &consul.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d%s", serviceName, bindPort, healthCheck),
			Method:                         http.MethodGet,
			Interval:                       fmt.Sprintf("%ds", agentCheckInterval),
			Timeout:                        fmt.Sprintf("%ds", agentCheckTimeout),
			DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", agentDeregisterAfter),
		},
		Address: serviceName,
		Port:    bindPort,
	}
}
