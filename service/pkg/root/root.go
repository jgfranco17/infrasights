package cmd

import (
	"fmt"
	"time"

	infra "github.com/jgfranco17/infrasights/core/pkg/infrastructure"
	vcs "github.com/jgfranco17/infrasights/core/pkg/versioncontrol"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Logger = logrus.New()

var rootCmd = &cobra.Command{
	Use:   "devops",
	Short: "A CLI tool for DevOps tasks",
	Long: `A CLI tool for DevOps tasks that prints out shell environment
information including git branch, Kubernetes namespace, and current time.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Time:", time.Now().Format("2006-01-02 15:04:05"))

		// Print current git branch
		gitBranch := vcs.GetGitBranch()
		fmt.Printf("Git Branch: %s\n", gitBranch)

		// Print current Kubernetes namespace
		k8sNamespace := infra.GetKubeNamespace()
		fmt.Printf("Kubernetes Namespace: %s\n", k8sNamespace)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
