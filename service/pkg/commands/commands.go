package commands

import (
	"fmt"
	"os"
	"runtime"
	"time"

	dock "github.com/jgfranco17/infrasights/core/pkg/containerization"
	env "github.com/jgfranco17/infrasights/core/pkg/environment"
	infra "github.com/jgfranco17/infrasights/core/pkg/infrastructure"
	log "github.com/jgfranco17/infrasights/core/pkg/logging"
	vcs "github.com/jgfranco17/infrasights/core/pkg/versioncontrol"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func printCli(data []ShellInfo) {
	for _, entity := range data {
		fmt.Printf("%s: %v\n", entity.Name, entity.Value)
	}
}

func MainCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "basic",
		Short: "Display basic devops-related information",
		Long:  `Print out information about Git, Kubernetes, Docker, and other infra components.`,
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

			printCli(data)
		},
	}
}

func GetGitDataCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "git",
		Short: "Display Git information",
		Long:  `Print out information about current Git workspace.`,
		Run: func(cmd *cobra.Command, args []string) {
			var data []ShellInfo

			gitBranch, err := vcs.GetGitBranch()
			if err != nil {
				log.Fatal(err.Error())
			}
			commitCount, err := vcs.GetCommitCountOnCurrentBranch()
			if err != nil {
				log.Fatal(err.Error())
			}
			data = append(data, ShellInfo{
				Name:  "Current branch",
				Value: fmt.Sprintf("%s (%v commits)", gitBranch, commitCount),
			})
			commitSha, err := vcs.GetGitCommitSHA()
			if err != nil {
				log.Fatal(err.Error())
			}
			commitTime, err := vcs.GetGitLastCommitTime()
			if err != nil {
				log.Fatal(err.Error())
			}
			data = append(data, ShellInfo{
				Name:  "Last commit",
				Value: fmt.Sprintf("[%s] %s", commitTime, commitSha),
			})

			printCli(data)
		},
	}
}

func GetMachineInfo() *cobra.Command {
	return &cobra.Command{
		Use:   "device",
		Short: "Display local environment data",
		Long:  `Print out information about currently-used local machine.`,
		Run: func(cmd *cobra.Command, args []string) {
			var data []ShellInfo
			tc := cases.Title(language.English)
			data = append(data, ShellInfo{
				Name:  "Operating system",
				Value: fmt.Sprintf("%s (%s)", tc.String(runtime.GOOS), runtime.GOARCH),
			})
			hostName, err := os.Hostname()
			if err != nil {
				log.Fatal(err.Error())
			}
			data = append(data, ShellInfo{
				Name:  "Hostname",
				Value: hostName,
			})
			data = append(data, ShellInfo{
				Name:  "Go root",
				Value: runtime.GOROOT(),
			})
			data = append(data, ShellInfo{
				Name:  "CPU count",
				Value: runtime.NumCPU(),
			})
			workingDir, err := env.GetWorkingDir()
			if err != nil {
				log.Fatal(err.Error())
			}
			data = append(data, ShellInfo{
				Name:  "Working directory",
				Value: workingDir,
			})
			printCli(data)
		},
	}
}
