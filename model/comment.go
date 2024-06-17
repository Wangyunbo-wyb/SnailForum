package model

import "gorm.io/gorm"

type Comment struct {
	CommentID int64  `gorm:"index:idx_comment_id;unique;not null" json:"comment_id"`
	Content   string `gorm:"type:text;not null" json:"content"`
	PostID    uint64 `gorm:"not null" json:"post_id"`
	AuthorID  uint64 `gorm:"index:idx_author_id;not null" json:"author_id"`
	ParentID  uint64 `gorm:"default:0;not null" json:"parent_id"`
	Status    uint8  `gorm:"type:tinyint;default:1;not null" json:"status"`
	gorm.Model
}
