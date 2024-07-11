package util

import (
	"errors"
	"os/exec"
	"strings"
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

func GetContainerId(projectName string, childName string) (string, string, error) {
	output, _ := DockerPs(projectName, false)
	targetName := projectName + "/" + childName

	outputList := strings.Split(output, "\n")
	for index, line := range outputList {
		if index == 0 {
			// 0 is the header
			continue
		}
		infoSplit := strings.Split(line, "   ")
		if len(infoSplit) >= 2 {
			imageName := infoSplit[1]
			imageInfo := strings.Split(imageName, ":")[0]
			if imageInfo == targetName {
				containerId := infoSplit[0]
				return containerId, targetName, nil
			}
		}
	}

	return "", targetName, errors.New("No image found")
}
