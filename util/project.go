package util

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
	"github.com/fatih/color"
)

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

func AddProject(projectList model.ProjectList, projectName, childName, projectPath string) model.ProjectList {

	parentIndex := -1
	var parentProject *model.ProjectParentModel
	for i, project := range projectList.List {
		if project.ProjectName == projectName {
			parentIndex = i
			parentProject = &project
			break
		}
	}

	if parentIndex == -1 {
		parentProject = &model.ProjectParentModel{
			ProjectName: projectName,
			List:        []model.ProjectChildModel{},
		}
	}

	childIndex := -1
	var childProject *model.ProjectChildModel
	for i, child := range parentProject.List {
		if child.ProjectName == childName {
			childIndex = i
			childProject = &child
			break
		}
	}

	flag := AnalysisFlag(childName, parentProject.List)
	if childIndex == -1 {
		childProject = &model.ProjectChildModel{
			ProjectName: childName,
			ProjectPath: projectPath,
			Flag:        flag,
		}
	}

	addedChild := false
	addedParent := false

	if childIndex == -1 {
		parentProject.List = append(parentProject.List, *childProject)
		addedChild = true
	} else {
		parentProject.List[childIndex] = *childProject
	}

	if parentIndex == -1 {
		projectList.List = append(projectList.List, *parentProject)
		addedParent = true
	} else {
		projectList.List[parentIndex] = *parentProject
	}

	if addedParent {
		color.Cyan("--> Added New Project " + projectName)
	} else {
		color.Cyan("--> Updated Project " + projectName)
	}
	if addedChild {
		color.Cyan("--> Added New Child Project " + childName)
		fmt.Printf("%v %v %v\n", color.CyanString("--> You can use flag"), color.YellowString(flag), color.CyanString("for Child Project "+childName))
	} else {
		color.Cyan("--> Updated Child Project " + childName)
	}

	return projectList
}

func RemoveProject(projectList model.ProjectList, parentName, childName string) model.ProjectList {
	parentIndex := -1
	var parentProject *model.ProjectParentModel
	for i, project := range projectList.List {
		if project.ProjectName == parentName {
			parentIndex = i
			parentProject = &project
			break
		}
	}

	if parentIndex == -1 {
		color.Red("--> Parent Project " + parentName + " does not exist")
		return projectList
	}

	childIndex := -1
	for i, child := range parentProject.List {
		if child.ProjectName == childName {
			childIndex = i
			break
		}
	}

	if childIndex == -1 {
		color.Red("--> Child Project " + childName + " does not exist")
		return projectList
	}

	parentProject.List = append(parentProject.List[:childIndex], parentProject.List[childIndex+1:]...)
	projectList.List[parentIndex] = *parentProject

	if len(parentProject.List) == 0 {
		projectList.List = append(projectList.List[:parentIndex], projectList.List[parentIndex+1:]...)
	}

	color.Cyan("--> Removed Child Project " + childName)
	return projectList
}

func RemoveWholeParentProject(projectList model.ProjectList, parentName string) model.ProjectList {
	parentIndex := -1
	for i, project := range projectList.List {
		if project.ProjectName == parentName {
			parentIndex = i
			break
		}
	}

	if parentIndex == -1 {
		color.Red("--> Parent Project " + parentName + " does not exist")
		return projectList
	}

	projectList.List = append(projectList.List[:parentIndex], projectList.List[parentIndex+1:]...)

	color.Cyan("--> Removed Parent Project " + parentName)
	return projectList
}
