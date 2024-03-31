package cmd

import (
	"fmt"
	"time"

	dock "github.com/jgfranco17/infrasights/core/pkg/containerization"
	env "github.com/jgfranco17/infrasights/core/pkg/environment"
	infra "github.com/jgfranco17/infrasights/core/pkg/infrastructure"
	log "github.com/jgfranco17/infrasights/core/pkg/logging"
	vcs "github.com/jgfranco17/infrasights/core/pkg/versioncontrol"
	"github.com/spf13/cobra"
)

type ShellInfo struct {
	Name  string
	Value interface{}
}

var rootCmd = &cobra.Command{
	Use:   "devops",
	Short: "A CLI tool for DevOps tasks",
	Long: `A CLI tool for DevOps tasks that prints out shell environment
information including git branch, Kubernetes namespace, and current time.`,
	Run: func(cmd *cobra.Command, args []string) {
		var data []ShellInfo

		data = append(data, ShellInfo{
			Name:  "Current time",
			Value: time.Now().Format("2006-01-02 15:04:05"),
		})
		workingDir, err := env.GetWorkingDir()
		if err != nil {
			log.Fatal(err.Error())
		}
		data = append(data, ShellInfo{
			Name:  "Working directory",
			Value: workingDir,
		})

		imageCount, err := dock.CountDockerImages()
		if err != nil {
			log.Fatal(err.Error())
		}
		data = append(data, ShellInfo{
			Name:  "Docker image count",
			Value: imageCount,
		})
		gitBranch, err := vcs.GetGitBranch()
		if err != nil {
			log.Fatal(err.Error())
		}
		data = append(data, ShellInfo{
			Name:  "Git branch",
			Value: gitBranch,
		})
		data = append(data, ShellInfo{
			Name:  "Kubernetes namespace",
			Value: infra.GetKubeNamespace(),
		})

		for _, entity := range data {
			fmt.Printf("%s: %v\n", entity.Name, entity.Value)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
