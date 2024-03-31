package versioncontrol

import (
	"fmt"
	"os/exec"
	"strconv"
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

func GetGitCommitSHA() (string, error) {
	lastCommitSHA, err := executeGitCommand("rev-parse", "HEAD")
	if err != nil {
		return "", fmt.Errorf("Failed to get last commit SHA: %v", err)
	}
	return strings.TrimSpace(lastCommitSHA), nil
}

func GetGitLastCommitTime() (string, error) {
	lastCommitTime, err := executeGitCommand("show", "-s", "--format=%ci", "HEAD")
	if err != nil {
		return "", fmt.Errorf("Failed to get last commit timestamp: %v", err)
	}
	return strings.TrimSpace(lastCommitTime), nil
}

func GetCommitCountOnCurrentBranch() (int, error) {
	// Execute the Git command to get the commit count
	output, err := executeGitCommand("rev-list", "--count", "HEAD")
	if err != nil {
		return 0, fmt.Errorf("Failed to get commit count: %v", err)
	}

	// Convert the output to an integer
	count, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		return 0, fmt.Errorf("Failed to convert commit count to integer: %v", err)
	}

	return count, nil
}

func executeGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
