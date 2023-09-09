package models

type TagModel struct {
	MODEL
	Title    string         `gorm:"size:16" json:"title"`
	Articles []ArticleModel `gorm:"many2many:article_tag_models" json:"-"`
}
