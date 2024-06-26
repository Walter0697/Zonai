package model

type ProjectList struct {
	List []ProjectParentModel
}

type ProjectParentModel struct {
	ProjectName string
	List        []ProjectChildModel
}

type ProjectChildModel struct {
	ProjectName string
	ProjectPath string
	Flag        string
}
