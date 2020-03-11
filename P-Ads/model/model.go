package model



/*type AdsConfig struct {
	ID  int  `json:"ID"`
	Ads Ads    `json:"ads"`
}*/

type Ads struct {
	ID  int  `json:"ID"`
	Category    [] Category     `json:"category"`
	SubCategory  []SubCategory  `json:"subCategory"`
	BottomBanner BottomBanner `json:"bottomBanner"`
	FullScreen   FullScreen   `json:"fullScreen"`
}

type Category struct {
	ID          string    `json:"id"`
	CategoryAds []AdsItem `json:"categoryAds"`
}

type SubCategory struct {
	ID          string    `json:"id"`
	CategoryAds []AdsItem `json:"categoryAds"`
}

type BottomBanner struct {
	AdsID       string `json:"adsId"`
	TotalRequest int `json:"totalRequest"`
}

type FullScreen struct {
	AdsID         string `json:"adsId"`
	TotalRequest  int `json:"totalRequest"`
	ShowInDetail  bool `json:"showInDetail"`
	ShowInListing bool `json:"showInListing"`
}

type AdsItem struct {
	Position int `json:"position"`
	AdsID    string `json:"adsId"`
	Type     string `json:"type"`
}


///////////////////////////
type TestItem struct {
	ID int
	Year   int
	Title  string
	Plot   string
	Rating float64
}