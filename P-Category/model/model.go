package model

import "Publisher/P-Category/constant"

func DefaultCategoryItem() CategoryItem {
	categoryItem := CategoryItem{}
	categoryItem.Id = ""
	categoryItem.Name = ""
	categoryItem.ShowInSlider = true
	categoryItem.ShowInTopbar = true
	categoryItem.IsFree = true
	categoryItem.ThumbUrl = ""
	categoryItem.BannerUrl = ""
	categoryItem.ShortMsg = ""
	categoryItem.SubCategory = append(categoryItem.SubCategory, DefaultSubCategoryItem())
	categoryItem.CategoryLink = "Category google.com"
	categoryItem.NonGroupedItemViewConfig = append(categoryItem.NonGroupedItemViewConfig, DefaultViewConfig())
	categoryItem.ArticleGroup1 = append(categoryItem.ArticleGroup1, DefaultArticleGroup())
	categoryItem.ArticleGroup2 = append(categoryItem.ArticleGroup2, DefaultArticleGroup())
	categoryItem.ArticleGroup3 = append(categoryItem.ArticleGroup3, DefaultArticleGroup())
	categoryItem.ArticleGroup4 = append(categoryItem.ArticleGroup4, DefaultArticleGroup())
	categoryItem.ArticleGroup5 = append(categoryItem.ArticleGroup5, DefaultArticleGroup())
	categoryItem.ArticleGroup6 = append(categoryItem.ArticleGroup6, DefaultArticleGroup())
	categoryItem.ArticleGroup7 = append(categoryItem.ArticleGroup7, DefaultArticleGroup())
	categoryItem.ArticleGroup8 = append(categoryItem.ArticleGroup8, DefaultArticleGroup())
	categoryItem.ArticleGroup9 = append(categoryItem.ArticleGroup9, DefaultArticleGroup())

	return categoryItem
}

type CategoryItem struct {
	Id                       string            `json:"ID" "required"`
	Name                     string            `json:"Name" "required"`
	ShowInSlider             bool              `json:"showInSlider"`
	ShowInTopbar             bool              `json:"showInTopbar"`
	IsFree                   bool              `json:"isFree"`
	CategoryLink             string            `json:"catLink"`
	SubCategory              []SubCategoryItem `json:"subCategory"`
	NonGroupedItemViewConfig []ViewConfig      `json:"nonGroupedItemViewConfig"`
	ThumbUrl                 string            `json:thumbUrl`
	BannerUrl                string            `json:bannerUrl`
	ShortMsg                 string            `json:shortMsg`
	ArticleGroup1            []ArticleGroup    `json:articleGroup1`
	ArticleGroup2            []ArticleGroup    `json:articleGroup2`
	ArticleGroup3            []ArticleGroup    `json:articleGroup3`
	ArticleGroup4            []ArticleGroup    `json:articleGroup4`
	ArticleGroup5            []ArticleGroup    `json:articleGroup5`
	ArticleGroup6            []ArticleGroup    `json:articleGroup6`
	ArticleGroup7            []ArticleGroup    `json:articleGroup7`
	ArticleGroup8            []ArticleGroup    `json:articleGroup8`
	ArticleGroup9            []ArticleGroup    `json:articleGroup9`
}

func DefaultSubCategoryItem() SubCategoryItem {
	subCategoryItem := SubCategoryItem{}
	subCategoryItem.ParentId = ""
	subCategoryItem.Id = ""
	subCategoryItem.Name = ""
	subCategoryItem.ShowInSlider = true
	subCategoryItem.ShowInTopbar = true
	subCategoryItem.IsFree = true
	subCategoryItem.ThumbUrl = ""
	subCategoryItem.BannerUrl = ""
	subCategoryItem.ShortMsg = ""
	subCategoryItem.SubCategoryLink = "Sub google.com"
	subCategoryItem.InCategoryItemViewConfig = DefaultViewConfig()
	subCategoryItem.NonGroupedItemViewConfig = DefaultViewConfig()
	subCategoryItem.ArticleGroup1 = append(subCategoryItem.ArticleGroup1, DefaultArticleGroup())
	subCategoryItem.ArticleGroup2 = append(subCategoryItem.ArticleGroup2, DefaultArticleGroup())
	subCategoryItem.ArticleGroup3 = append(subCategoryItem.ArticleGroup3, DefaultArticleGroup())
	subCategoryItem.ArticleGroup4 = append(subCategoryItem.ArticleGroup4, DefaultArticleGroup())
	subCategoryItem.ArticleGroup5 = append(subCategoryItem.ArticleGroup5, DefaultArticleGroup())
	subCategoryItem.ArticleGroup6 = append(subCategoryItem.ArticleGroup6, DefaultArticleGroup())
	subCategoryItem.ArticleGroup7 = append(subCategoryItem.ArticleGroup7, DefaultArticleGroup())
	subCategoryItem.ArticleGroup8 = append(subCategoryItem.ArticleGroup8, DefaultArticleGroup())
	subCategoryItem.ArticleGroup9 = append(subCategoryItem.ArticleGroup9, DefaultArticleGroup())

	return subCategoryItem
}

type SubCategoryItem struct {
	ParentId                 string         `json:"ParentId"`
	Id                       string         `json:"ID" "required"`
	Name                     string         `json:"Name" "required"`
	ShowInSlider             bool           `json:"showInSlider"`
	ShowInTopbar             bool           `json:"showInTopbar"`
	IsFree                   bool           `json:"isFree"`
	SubCategoryLink          string         `json:"subCatLink"`
	NonGroupedItemViewConfig ViewConfig     `json:"nonGroupedItemViewConfig"`
	InCategoryItemViewConfig ViewConfig     `json:"InCategoryItemViewConfig"`
	ThumbUrl                 string         `json:thumbUrl`
	BannerUrl                string         `json:bannerUrl`
	ShortMsg                 string         `json:shortMsg`
	ArticleGroup1            []ArticleGroup `json:articleGroup1`
	ArticleGroup2            []ArticleGroup `json:articleGroup2`
	ArticleGroup3            []ArticleGroup `json:articleGroup3`
	ArticleGroup4            []ArticleGroup `json:articleGroup4`
	ArticleGroup5            []ArticleGroup `json:articleGroup5`
	ArticleGroup6            []ArticleGroup `json:articleGroup6`
	ArticleGroup7            []ArticleGroup `json:articleGroup7`
	ArticleGroup8            []ArticleGroup `json:articleGroup8`
	ArticleGroup9            []ArticleGroup `json:articleGroup9`
}

func DefaultArticleGroup() ArticleGroup {
	articleGroup := ArticleGroup{}
	articleGroup.ViewConfig = DefaultViewConfig()
	articleGroup.ArticleIds = append(articleGroup.ArticleIds, "-1")
	return articleGroup
}

type ArticleGroup struct {
	ViewConfig ViewConfig `json:"viewConfig"`
	ArticleIds []string   `json:"articleIds"`
}

func DefaultViewConfig() ViewConfig {
	viewConfig := ViewConfig{}
	viewConfig.ViewType = constant.VT_ITEM
	viewConfig.ViewIndex = -1
	viewConfig.ViewTitle = "View Title"
	viewConfig.ActionTitle = "Action Title"
	viewConfig.Redirection = "google.com"
	return viewConfig
}

type ViewConfig struct {
	ViewType    string `json:"viewType"`
	ViewIndex   int    `json:"viewIndex"`
	Redirection string `json:"redirection"`
	ViewTitle   string `json:"viewTitle"`
	ActionTitle string `json:"actionTitle"`
}
