package versioncontrol

import (
	"os/exec"
	"strings"
)

func GetGitBranch() string {
	// Run git command to get current branch
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}
