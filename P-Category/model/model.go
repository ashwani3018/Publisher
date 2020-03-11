package model

type Category struct {
	Id string `json:"ID" "required"`
	Name string `json:"Name" "required"`
	ShowInSlider bool `json:"showInSlider"`
	ShowInTopbar bool `json:"showInTopbar"`
	IsFree bool `json:"isFree"`
	SubCategory []string `json:"subCategory"`
}
