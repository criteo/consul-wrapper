package runner

import (
	"time"

	"github.com/hashicorp/consul/api"
)

type Config struct {
	ConsulAddr         string
	ConsulToken        string
	CommandLine        []string
	Definition         *api.AgentServiceRegistration
	SelfCheckFrequency time.Duration
}
