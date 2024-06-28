package util

import (
	"os"
	"path"

	"github.com/Walter0697/zonai/model"
)

func GetAllEnvironments(configuration *model.ProjectConfigurationModel, parentProject *model.ProjectParentModel) []string {
	env_set := map[string]bool{}
	for _, child := range parentProject.List {
		environmentFolder := parentProject.ProjectName + "_" + child.ProjectName
		environmentParentPath := path.Join(configuration.EnviromentPath, environmentFolder)

		// read all folders in the environment path
		if _, err := os.Stat(environmentParentPath); err == nil {
			folders, err := os.ReadDir(environmentParentPath)
			if err != nil {
				panic(err)
			}
			for _, file := range folders {
				if file.IsDir() {
					env_set[file.Name()] = true
				}
			}
		}
	}

	// convert set to slice
	var environments []string
	for k := range env_set {
		environments = append(environments, k)
	}
	return environments
}
