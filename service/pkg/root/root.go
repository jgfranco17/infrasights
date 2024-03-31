package cmd

import (
	"fmt"
	"os"
	"time"

	vcs "github.com/jgfranco17/infrasights/core/pkg/versioncontrol"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devops",
	Short: "A CLI tool for DevOps tasks",
	Long: `A CLI tool for DevOps tasks that prints out shell environment
information including git branch, Kubernetes namespace, and current time.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Print current time
		fmt.Println("Current Time:", time.Now().Format("2006-01-02 15:04:05"))

		// Print current git branch
		gitBranch := vcs.GetGitBranch()
		fmt.Println("Git Branch:", gitBranch)

		// Print current Kubernetes namespace
		k8sNamespace := os.Getenv("NAMESPACE")
		fmt.Println("Kubernetes Namespace:", k8sNamespace)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()
}
