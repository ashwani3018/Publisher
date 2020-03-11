package model

type Category struct {
	Id string `json:"ID" "required"`
	Name string `json:"Name" "required"`
	ShowInSlider bool `json:"showInSlider"`
	ShowInTopbar bool `json:"showInTopbar"`
	IsFree bool `json:"isFree"`
	CategoryLink string `json:"catLink"`
	SubCategory []Category `json:"subCategory"`
	ViewConfiguration []ViewConfig `json:"viewConfig"`
	ThumbUrl string `json:thumbUrl`
	BannerUrl string `json:bannerUrl`
	ShortMsg string `json:shortMsg`
}

type ViewConfig struct {
	ViewType string `json:"viewType"`
	ViewIndex int `json:"viewIndex"`
	redirection string `json:"redirection"`
	ViewTitle string `json:"viewTitle"`
	ActionTitle string `json:"actionTitle"`
}
