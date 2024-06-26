package util

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/Walter0697/zonai/model"
)

const (
	DataFile             = ".zonai"
	ProjectConfiguration = "project_configuration.json"
	BuildHistory         = "build_history.json"
	ProjectList          = "project_list.json"

	// configuration default
	DefaultDockerBuildCommand = "docker build -t"
)

func getFilePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return path.Join(exPath, DataFile)
}

// begin read file
func ReadConfiguration() model.ProjectConfigurationModel {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, ProjectConfiguration)
	configuration := LoadJsonFile[model.ProjectConfigurationModel](filePath)
	return configuration
}

func ReadBuildHistory() model.BuildHistory {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, BuildHistory)
	history := LoadJsonFile[model.BuildHistory](filePath)
	return history
}

func ReadProjectList() model.ProjectList {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, ProjectList)
	projectList := LoadJsonFile[model.ProjectList](filePath)
	return projectList
}

// end read file

// begin save file
func SaveConfiguration(configuration model.ProjectConfigurationModel) {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, ProjectConfiguration)
	SaveJsonFile(configuration, filePath)
}

func SaveBuildHistory(history model.BuildHistory) {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, BuildHistory)
	SaveJsonFile(history, filePath)
}

func SaveProjectList(projectList model.ProjectList) {
	folderPath := getFilePath()

	filePath := path.Join(folderPath, ProjectList)
	SaveJsonFile(projectList, filePath)
}

// end save file

func InitializeFolder() {
	folderPath := getFilePath()
	mkdirerr := os.MkdirAll(folderPath, 0755)
	if mkdirerr != nil {
		panic(mkdirerr)
	}

	configFolder := path.Join(folderPath, ProjectConfiguration)
	if _, err := os.Stat(configFolder); errors.Is(err, os.ErrNotExist) {
		var defaultConfiguration = model.ProjectConfigurationModel{
			OutputImagePath:    "",
			DockerBuildCommand: DefaultDockerBuildCommand,
		}
		SaveConfiguration(defaultConfiguration)
	}

	historyFolder := path.Join(folderPath, BuildHistory)
	if _, err := os.Stat(historyFolder); errors.Is(err, os.ErrNotExist) {
		var defaultHistory = model.BuildHistory{
			List: []model.BuildItem{},
		}
		SaveBuildHistory(defaultHistory)
	}

	projectListFolder := path.Join(folderPath, ProjectList)
	if _, err := os.Stat(projectListFolder); errors.Is(err, os.ErrNotExist) {

		var defaultProjectList = model.ProjectList{
			List: []model.ProjectParentModel{},
		}

		SaveProjectList(defaultProjectList)
	}
}