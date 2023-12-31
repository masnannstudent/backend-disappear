package entities

import (
	"time"
)

type ArticleModels struct {
	ID        uint64     `gorm:"column:id;primaryKey" json:"id"`
	Title     string     `gorm:"column:title;type:varchar(255)" json:"title"`
	Photo     string     `gorm:"column:photo;type:varchar(255)" json:"photo"`
	Content   string     `gorm:"column:content;type:text" json:"content"`
	Author    string     `gorm:"column:author;type:varchar(255)" json:"author"`
	Views     uint64     `gorm:"column:views" json:"views"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP NULL;index" json:"deleted_at"`
}

type ArticleBookmarkModels struct {
	ID        uint64         `gorm:"column:id;primaryKey" json:"id"`
	UserID    uint64         `gorm:"column:user_id" json:"user_id"`
	ArticleID uint64         `gorm:"column:article_id" json:"voucher_id"`
	User      *UserModels    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Article   *ArticleModels `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
}

func (ArticleModels) TableName() string {
	return "articles"
}

func (ArticleBookmarkModels) TableName() string {
	return "article_bookmarks"
}
