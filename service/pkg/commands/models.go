package commands

import (
	"github.com/spf13/cobra"
)

type ShellInfo struct {
	Name  string
	Value interface{}
}

type CommandRegistry struct {
	rootCmd *cobra.Command
}

// NewCommandRegistry creates a new instance of CommandRegistry
func NewCommandRegistry(name string, description string) *CommandRegistry {
	return &CommandRegistry{
		rootCmd: &cobra.Command{
			Use:   name,
			Short: description,
		},
	}
}

// RegisterCommand registers a new command with the CommandRegistry
func (cr *CommandRegistry) RegisterCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		cr.rootCmd.AddCommand(cmd)
	}
}

// Execute executes the root command
func (cr *CommandRegistry) Execute() error {
	return cr.rootCmd.Execute()
}
