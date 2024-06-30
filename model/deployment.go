package model

type DeploymentHistory struct {
	List []DeploymentItemModel
}

type DeploymentItemModel struct {
	FileName  string
	BuildTime string
	ImageList []DeploymentImageItem
}

type DeploymentImageItem struct {
	ImageTag    string
	ProjectName string
	ProjectPath string
}
