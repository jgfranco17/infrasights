package main

import (
	log "github.com/jgfranco17/infrasights/core/pkg/logging"
	cmd "github.com/jgfranco17/infrasights/service/pkg/commands"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

func main() {
	projectName := "infrasights"
	projectDescription := "A CLI DevOps assistive utility"
	commandsList := []*cobra.Command{
		cmd.MainCommand(),
		cmd.GetGitDataCommand(),
	}
	command := cmd.NewCommandRegistry(projectName, projectDescription)
	command.RegisterCommands(commandsList)

	err := command.Execute()
	if err != nil {
		log.Error(err.Error())
	}
}
