package util

import (
	"os"
	"path"
	"strings"

	"github.com/Walter0697/zonai/model"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

func FindComposeAndEdit(imageTag string) (string, *model.DeploymentImageItem) {
	deploymentList := ReadDeploymentList()
	imageInfo := strings.Split(imageTag, ":")
	imageRepoInfo := strings.Split(imageInfo[0], "/")
	parentName := imageRepoInfo[0]
	childName := imageRepoInfo[1]

	folderPath := ""
	for _, parentDeployment := range deploymentList.List {
		if parentDeployment.ProjectName == parentName {
			for _, childDeployment := range parentDeployment.List {
				if childDeployment.ProjectName == childName {
					folderPath = childDeployment.ProjectPath
				}
			}
		}
	}

	if folderPath == "" {
		imageNameDisplay := color.YellowString(imageTag)
		color.Red("Deployment for " + imageNameDisplay + " not found")
		return "", nil
	}

	dockerComposePath := path.Join(folderPath, "docker-compose.yml")
	// check if docker compose exists
	if _, err := os.Stat(dockerComposePath); os.IsNotExist(err) {
		color.Red("docker-compose.yml not found in " + folderPath)
		return "", nil
	}

	// read docker compose
	dockerComposeContent, _ := os.ReadFile(dockerComposePath)

	dockerCompose := make(map[interface{}]interface{})
	err := yaml.Unmarshal([]byte(dockerComposeContent), &dockerCompose)
	if err != nil {
		panic(err)
	}

	serviceName := parentName + "_" + childName

	if _, ok := dockerCompose["services"]; !ok {
		color.Red("services not found in docker-compose.yml")
		return "", nil
	}

	services := dockerCompose["services"].(map[interface{}]interface{})

	if _, ok := services[serviceName]; !ok {
		imageNameDisplay := color.YellowString(serviceName)
		color.Red("Service " + imageNameDisplay + " not found in docker-compose.yml")
		return "", nil
	}

	currentProject := services[serviceName].(map[interface{}]interface{})
	if _, ok := currentProject["image"]; !ok {
		imageNameDisplay := color.YellowString(serviceName)
		color.Red("image not found in " + imageNameDisplay)
		return "", nil
	}

	currentProject["image"] = imageTag

	// write back to docker-compose.yml
	dockerComposeYaml, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dockerComposePath, dockerComposeYaml, 0644)
	if err != nil {
		panic(err)
	}

	var deploymentImageItem model.DeploymentImageItem
	deploymentImageItem.ImageTag = imageTag
	deploymentImageItem.ProjectPath = folderPath
	deploymentImageItem.ProjectName = parentName

	return folderPath, &deploymentImageItem
}
