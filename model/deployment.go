package model

type DeploymentHistory struct {
	List []DeploymentHistoryModel
}

type DeploymentHistoryModel struct {
	ProjectName string
	History     []string
}
