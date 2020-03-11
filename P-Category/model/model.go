package model

type Category struct {
	Id                       string        `json:"ID" "required"`
	Name                     string        `json:"Name" "required"`
	ShowInSlider             bool          `json:"showInSlider"`
	ShowInTopbar             bool          `json:"showInTopbar"`
	IsFree                   bool          `json:"isFree"`
	CategoryLink             string        `json:"catLink"`
	SubCategory              SubCategory `json:"subCategory"`
	NonGroupedItemViewConfig []ViewConfig  `json:"viewConfig"`
	ThumbUrl                 string        `json:thumbUrl`
	BannerUrl                string        `json:bannerUrl`
	ShortMsg                 string        `json:shortMsg`
	ArticleGroup1 []string   `json:articleGroup1`
	ArticleGroup2 []string   `json:articleGroup2`
	ArticleGroup3 []string   `json:articleGroup3`
	ArticleGroup4 []string   `json:articleGroup4`
	ArticleGroup5 []string   `json:articleGroup5`
	ArticleGroup6 []string   `json:articleGroup6`
	ArticleGroup7 []string   `json:articleGroup7`
	ArticleGroup8 []string   `json:articleGroup8`
	ArticleGroup9 []string   `json:articleGroup9`
}

type SubCategory struct {
	Id            string     `json:"ID" "required"`
	Name          string     `json:"Name" "required"`
	ShowInSlider  bool       `json:"showInSlider"`
	ShowInTopbar  bool       `json:"showInTopbar"`
	IsFree        bool       `json:"isFree"`
	CategoryLink  string     `json:"catLink"`
	ViewConfig    ViewConfig `json:"viewConfig"`
	ThumbUrl      string     `json:thumbUrl`
	BannerUrl     string     `json:bannerUrl`
	ShortMsg      string     `json:shortMsg`
	ArticleGroup1 []string   `json:articleGroup1`
	ArticleGroup2 []string   `json:articleGroup2`
	ArticleGroup3 []string   `json:articleGroup3`
	ArticleGroup4 []string   `json:articleGroup4`
	ArticleGroup5 []string   `json:articleGroup5`
	ArticleGroup6 []string   `json:articleGroup6`
	ArticleGroup7 []string   `json:articleGroup7`
	ArticleGroup8 []string   `json:articleGroup8`
	ArticleGroup9 []string   `json:articleGroup9`
}

type ViewConfig struct {
	ViewType    string `json:"viewType"`
	ViewIndex   int    `json:"viewIndex"`
	redirection string `json:"redirection"`
	ViewTitle   string `json:"viewTitle"`
	ActionTitle string `json:"actionTitle"`
}
