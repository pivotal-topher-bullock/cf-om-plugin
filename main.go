package main

import (
	"code.cloudfoundry.org/cli/plugin"
	flags "github.com/jessevdk/go-flags"
	"github.com/pivotal-topher-bullock/cf-om-plugin/cmd"
)

// OmPlugin is the struct implementing the interface defined by the core CLI. It can
// be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type OmPlugin struct{}

func (c *OmPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	var subCmd cmd.subCmd
	if args[1] == "login" {
		subCmd = cmd.Login{}
	}

	if subCmd != nil {
		args, err := flags.ParseArgs(&subCmd, args)
		subCmd.Run()
	}
}

func (c *OmPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "om",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "om",
				HelpText: "cf login using credentials from ERT in OpsManger",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "login\n  cf om login om-username om-password om-url username",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(OmPlugin))
}
