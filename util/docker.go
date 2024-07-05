package util

import (
	"os/exec"
)

func IsDockerRunning() bool {
	// check if docker is running
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	return err == nil
}

func DockerPs(projectName string, allFlag bool) (string, error) {
	cmdList := []string{"ps"}
	if allFlag {
		cmdList = append(cmdList, "-a")
	}
	cmdList = append(cmdList, "-f")
	cmdList = append(cmdList, "name="+projectName)
	cmd := exec.Command("docker", cmdList...)

	result, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(result), nil
}
