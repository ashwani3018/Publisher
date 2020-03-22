package model

type Category struct {
	CategoryName string
	SubCategories map[string]SubCategory
}

type SubCategory struct {
	SubCategoryName string
	SubSubCategories map[string]SubSubCategory
}

type SubSubCategory struct {
	SubSubCategoryName string
}
