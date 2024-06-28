package util

import (
	"fmt"

	"github.com/Walter0697/zonai/model"
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

func ListProject() {
	projectList := ReadProjectList()
	Divider()
	fmt.Println("Project List")
	ListData(projectList)
	Divider()
}

func ListDeployment() {
	deploymentList := ReadDeploymentList()
	Divider()
	fmt.Println("Deployment List")
	ListData(deploymentList)
	Divider()
}
