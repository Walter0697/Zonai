package util

import (
	"bufio"
	"fmt"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/briandowns/spinner"
)

func BuildProject(parent *model.ProjectParentModel, child *model.ProjectChildModel, configuration *model.ProjectConfigurationModel, imageTag string) error {
	cmdList := []string{}
	buildCommandList := strings.Split(configuration.DockerBuildCommand, " ")
	for index, command := range buildCommandList {
		if index == 0 {
			continue
		}
		cmdList = append(cmdList, command)
	}

	// building the image
	imageName := parent.ProjectName + "/" + child.ProjectName
	fullImageName := imageName + ":" + imageTag
	cmdList = append(cmdList, fullImageName, ".")

	cmd := exec.Command("docker", cmdList...)
	cmd.Dir = child.ProjectPath

	s := spinner.New(spinner.CharSets[35], 500*time.Millisecond)
	s.Suffix = " Building " + imageName + "..."
	s.Start()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	cmd.Wait()
	s.Stop()

	fmt.Println("Image built successfully with tag: " + fullImageName)

	// saving into local image
	outputName := parent.ProjectName + "_" + child.ProjectName + "_" + imageTag + ".tar"
	outputFilePath := path.Join(configuration.OutputImagePath, outputName)
	saveImageCmd := exec.Command("docker", "save", "-o", outputFilePath, fullImageName)
	saveImageCmd.Dir = child.ProjectPath

	s2 := spinner.New(spinner.CharSets[43], 500*time.Millisecond)
	s2.Suffix = " Saving " + imageName + " to local file..."
	s2.Start()
	stdout2, _ := saveImageCmd.StdoutPipe()
	saveImageCmd.Start()

	scanner2 := bufio.NewScanner(stdout2)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		m := scanner2.Text()
		fmt.Println(m)
	}

	cmd.Wait()
	s2.Stop()

	fmt.Println("Image saved successfully with name: " + outputName)

	return nil
}

func GetImageName(parent *model.ProjectParentModel, child *model.ProjectChildModel) string {
	return parent.ProjectName + "/" + child.ProjectName
}
