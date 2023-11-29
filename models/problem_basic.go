package models

import (
	"gorm.io/gorm"
	"log"
)

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problem_categories"`
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`
	Content           string             `gorm:"column:content;type:text;" json:"content"`
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"max_runtime"`
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"max_mem"`
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"test_cases"` // 关联测试用例表
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}
func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategories").
		Preload("ProblemCategories.CategoryBasic").
		Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}

	log.Println(tx.Statement.SQL.String())
	return tx
}
