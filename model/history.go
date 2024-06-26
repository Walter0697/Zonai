package model

type BuildHistory struct {
	List []BuildItem
}

type BuildItem struct {
	ImageName    string
	BuildDate    string
	BuildVersion int
}
