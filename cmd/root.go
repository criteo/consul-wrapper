package cmd

import (
	"os"
	"runtime"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
	"github.com/tionebsalocin/consul-wrapper/runner"
)

var token string
var service string
var frequency time.Duration
var port int
var consulAddr string

func init() {
	RootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Consul token used for registration")
	RootCmd.PersistentFlags().StringVarP(&consulAddr, "consul-addr", "c", os.Getenv("CONSUL_HTTP_ADDR"), "Consul HTTP address")
	RootCmd.Flags().StringVarP(&service, "service", "s", "", "Consul Service Name")
	RootCmd.Flags().IntVarP(&port, "port", "p", 0, "Service port")
	RootCmd.PersistentFlags().DurationVarP(&frequency, "frequency", "f", 30*time.Second, "Health Check Frequency (in seconds)")
}

var RootCmd = &cobra.Command{
	Use:  "consul-wrapper",
	Long: "Run a program and register it in consul",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		shell := "sh"
		shellFlag := "-c"
		if runtime.GOOS == "windows" {
			shell = "cmd"
			shellFlag = "/c"
		}
		runner.Run(runner.Config{
			ConsulToken:        token,
			ConsulAddr:         consulAddr,
			CommandLine:        []string{shell, shellFlag, args[0]},
			SelfCheckFrequency: frequency,
			Definition: &api.AgentServiceRegistration{
				Name: service,
			},
		})
	},
}
