package util

import "github.com/Walter0697/zonai/model"

func AnalysisFlag(childString string, list []model.ProjectChildModel) string {
	// there should be no duplicated flag
	currentIndex := 1
	var currentFlag string = ""
	currentFlag = currentFlag + string(childString[0])
	needRecheck := true

	for needRecheck {
		needRecheck = false
		for _, child := range list {
			if child.Flag == currentFlag {
				currentFlag = currentFlag + string(childString[currentIndex])
				currentIndex++
				needRecheck = true
				break
			}
		}
		if needRecheck {
			break
		}
	}

	return currentFlag
}

func FindProject(projectName string) *model.ProjectParentModel {
	projectList := ReadProjectList()
	for _, project := range projectList.List {
		if project.ProjectName == projectName {
			return &project
		}
	}
	return nil
}

func GetAllProjectFlags(parent *model.ProjectParentModel) []string {
	var flags []string
	for _, child := range parent.List {
		flags = append(flags, child.Flag)
	}
	return flags
}
