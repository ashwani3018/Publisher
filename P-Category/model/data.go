package model

import "Publisher/P-Category/constant"


func GetCategoriesName() []string {
	return []string{"Home", "National", "International",  "State"}
}

func GetCategoriesIds() []string {
	return []string{"Home_id", "National_id", "International_id",  "State_id"}
}

func GetSubCategoriesName() []string {
	return []string{"Economy", "Markets", "Industry", "Bengaluru"}
}

func GetViewType() []string {
	return []string{constant.VT_ITEM, constant.VT_H_LIST, constant.VT_V_LIST, constant.VT_GRID, constant.VT_CARD, constant.VT_H_PAGER,
		constant.VT_V_PAGER, constant.VT_GALLERY, constant.VT_WEB, constant.VT_SubCategory}
}
