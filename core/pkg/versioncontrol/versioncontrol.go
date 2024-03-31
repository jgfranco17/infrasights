package versioncontrol

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetGitBranch() (string, error) {
	// Run git command to get current branch
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "unknown", fmt.Errorf("Failed to get branch: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}
