package main

import (
	log "github.com/jgfranco17/infrasights/core/pkg/logging"
	cmd "github.com/jgfranco17/infrasights/service/pkg/commands"
	sys "github.com/jgfranco17/infrasights/service/pkg/system"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

func main() {
	projectName := "infrasights"
	projectDescription := "A CLI DevOps assistive utility.\nDeveloped by Joaquin Franco."
	commandsList := []*cobra.Command{
		cmd.GetBasicDataCommand(),
		cmd.GetGitDataCommand(),
		cmd.GetMachineInfoCommand(),
		sys.GetVersionCommand(),
	}
	command := cmd.NewCommandRegistry(projectName, projectDescription)
	command.RegisterCommands(commandsList)

	err := command.Execute()
	if err != nil {
		log.Error(err.Error())
	}
}
