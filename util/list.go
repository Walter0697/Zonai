package util

import (
	"fmt"
	"strings"

	"github.com/Walter0697/zonai/model"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
)

func ListData(projectList model.ProjectList) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"#", "Project Name", "Child Name", "Flag", "Project Path"})
	for i, project := range projectList.List {
		if len(project.List) > 0 {
			firstProject := project.List[0]
			tw.AppendRow([]interface{}{i + 1, project.ProjectName, firstProject.ProjectName, firstProject.Flag, firstProject.ProjectPath})
		} else {
			tw.AppendRow([]interface{}{i + 1, project.ProjectName, "", "", ""})
		}

		for index, child := range project.List {
			if index == 0 {
				continue
			}
			tw.AppendRow([]interface{}{"", "", child.ProjectName, child.Flag, child.ProjectPath})
		}
	}
	fmt.Println(tw.Render())
}

func findAndReplace(compareText string, searchText string) (string, bool) {
	compareTextLower := strings.ToLower(compareText)
	searchTextLower := strings.ToLower(searchText)

	if strings.Contains(compareTextLower, searchTextLower) {
		// // find all index of searchText in compareText
		indexed := []int{}
		for i := 0; i < len(compareTextLower); i++ {
			valid := true
			for j := 0; j < len(searchTextLower); j++ {
				if i+j >= len(compareTextLower) {
					valid = false
					break
				}
				if compareTextLower[i+j] != searchTextLower[j] {
					valid = false
					break
				}
			}
			if valid {
				indexed = append(indexed, i)
			}
		}

		// replace all searchText with colored searchText
		finalText := compareText
		colorTextList := []interface{}{}
		replaceText := ""
		for i := 0; i < len(searchText); i++ {
			replaceText += "|"
		}
		for _, currentIndex := range indexed {
			finalText = finalText[:currentIndex] + replaceText + finalText[currentIndex+len(searchText):]
			colorTextList = append(colorTextList, color.YellowString(compareText[currentIndex:currentIndex+len(searchText)]))
		}

		finalText = strings.ReplaceAll(finalText, replaceText, "%s")
		finalText = fmt.Sprintf(finalText, colorTextList...)

		return finalText, true
	}

	return compareText, false
}

func FilterList(projectList model.ProjectList, searchKey *string) model.ProjectList {
	var filteredList model.ProjectList
	for _, project := range projectList.List {
		valid := false
		changed, v := findAndReplace(project.ProjectName, *searchKey)
		if v {
			project.ProjectName = changed
			valid = true
		}

		var filteredChildList []model.ProjectChildModel
		oneChildValid := false
		for _, child := range project.List {
			validChild := false
			changed, v := findAndReplace(child.ProjectName, *searchKey)
			if v {
				child.ProjectName = changed
				validChild = true
			}

			changed, v = findAndReplace(child.ProjectPath, *searchKey)
			if v {
				child.ProjectPath = changed
				validChild = true
			}

			if validChild || valid {
				filteredChildList = append(filteredChildList, child)
			}

			if validChild {
				oneChildValid = true
			}
		}

		if oneChildValid {
			valid = true
		}

		project.List = filteredChildList

		if valid {
			filteredList.List = append(filteredList.List, project)
		}
	}
	return filteredList
}

func ListProject(searchKey *string) {
	projectList := ReadProjectList()
	if searchKey != nil {
		projectList = FilterList(projectList, searchKey)
	}
	Divider()
	fmt.Println("Project List")
	ListData(projectList)
	Divider()
}

func ListDeployment(searchKey *string) {
	deploymentList := ReadDeploymentList()
	if searchKey != nil {
		deploymentList = FilterList(deploymentList, searchKey)
	}
	Divider()
	fmt.Println("Deployment List")
	ListData(deploymentList)
	Divider()
}
