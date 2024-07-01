package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/Walter0697/zonai/model"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func BuildProjectWithImageList(currentProject *model.ProjectParentModel, flags []string, configuration *model.ProjectConfigurationModel, history *model.BuildHistory, now string, compressFlag bool, currentEnvironment string) {
	image_list := []string{}
	for _, projects := range currentProject.List {
		for _, flag := range flags {
			if projects.Flag == flag {
				imageName := GetImageName(currentProject, &projects)
				version := 1
				found := false
				for hindex, h := range history.List {
					if h.ImageName == imageName {
						if h.BuildDate == now {
							version = h.BuildVersion + 1
							h.BuildVersion = version
						} else {
							version = 1
							h.BuildVersion = version
							h.BuildDate = now
						}
						found = true
						history.List[hindex] = h
						break
					}
				}
				if !found {
					history.List = append(history.List, model.BuildItem{
						ImageName:    imageName,
						BuildDate:    now,
						BuildVersion: 1,
					})
				}
				imageTag := now
				if version != 1 {
					imageTag = imageTag + "-" + fmt.Sprintf("%d", version)
				}
				imageFilename := BuildProject(currentProject, &projects, configuration, imageTag, currentEnvironment)
				if imageFilename != "" {
					image_list = append(image_list, imageFilename)
				}
				break
			}
		}
	}

	if compressFlag {
		CompressImageList(image_list, currentProject, configuration, currentEnvironment)
	}
}

func BuildProject(parent *model.ProjectParentModel, child *model.ProjectChildModel, configuration *model.ProjectConfigurationModel, imageTag string, currentEnvironment string) string {
	// if the current environment has any file, we will copy and replace the file
	if currentEnvironment != "" {
		environmentFolder := parent.ProjectName + "_" + child.ProjectName
		environmentParentPath := path.Join(configuration.EnviromentPath, environmentFolder)
		evironmentPath := path.Join(environmentParentPath, currentEnvironment)

		if _, err := os.Stat(evironmentPath); err == nil {
			// read all files from env path
			files, err := os.ReadDir(evironmentPath)
			if err != nil {
				panic(err)
			}

			// copy all files to the project path
			for _, file := range files {
				source := path.Join(evironmentPath, file.Name())
				destination := path.Join(child.ProjectPath, file.Name())
				// if the destination has a env file, copy one for backup
				if _, err := os.Stat(destination); err == nil {
					backupDestination := path.Join(child.ProjectPath, file.Name()+".backup")
					err := moveFile(destination, backupDestination)
					if err != nil {
						fmt.Println(err)
					}
				}

				err := copyFile(source, destination)
				if err != nil {
					fmt.Println(err)
				}

				color.Green("Copied " + file.Name() + " to " + child.ProjectPath)
			}

		}

		color.Blue("Current Environment: " + currentEnvironment)
	}

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

	imageNameDisplay := color.YellowString(fullImageName)
	fmt.Println("Image built successfully with tag: " + imageNameDisplay)

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

	savedImageDisplay := color.YellowString(outputName)
	fmt.Println("Image saved successfully with name: " + savedImageDisplay)

	return outputName
}

func GetImageName(parent *model.ProjectParentModel, child *model.ProjectChildModel) string {
	return parent.ProjectName + "/" + child.ProjectName
}

func copyFile(src, dst string) (err error) {
	// ignore when src does not exist
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return nil
	}

	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()

	_, err = io.Copy(out, in)
	return
}

func moveFile(src, dst string) (err error) {
	err = copyFile(src, dst)
	if err != nil {
		return
	}
	err = os.Remove(src)
	return
}
