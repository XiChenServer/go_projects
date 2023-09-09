package models

type AdverModel struct {
	MODEL
	Title  string `gorm:size:32 json:"title"`
	Href   string `json:"href"`
	Images string `json:"images"`
	IsShow bool   `json:"is_show"`
}
