package util

import "os/exec"

func IsDockerRunning() bool {
	// check if docker is running
	cmd := exec.Command("docker", "info")
	err := cmd.Run()
	return err == nil
}
