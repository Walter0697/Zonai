package util

import (
	"os/exec"
)

func ExecuteFetchAll(path string) (string, error) {
	cmd := exec.Command("git", "fetch", "--all")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func ExecuteGitListTags(path string) (string, error) {
	cmd := exec.Command("git", "tag")
	cmd.Dir = path
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ExecuteGitCheckout(path string, tag string) (string, error) {
	cmd := exec.Command("git", "checkout", "-f",  tag)
	cmd.Dir = path
	result, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(result), nil
}