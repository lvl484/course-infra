package consul

import consul "github.com/hashicorp/consul/api"

// ServiceRegister is used to register a new service with the agent
func ServiceRegister(client *consul.Client, agent *consul.AgentServiceRegistration) error {
	return client.Agent().ServiceRegister(agent)
}
