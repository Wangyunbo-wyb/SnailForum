package model

import (
	"gorm.io/gorm"
)

type Category struct {
	Name        string `gorm:"index:idx_name;unique;type:varchar(256);not null" json:"name"`
	Description string `gorm:"type:varchar(256)" json:"description"`
	gorm.Model
}

type Community struct {
	CommunityID   uint64 `gorm:"type:int(10);not null" json:"community_id"`
	CommunityName string `gorm:"type:varchar(128)" json:"community_name"`
	gorm.Model
}

// CommunityDetail 社区详情model
type CommunityDetail struct {
	CommunityID   uint64 `gorm:"type:int(10);not null" json:"community_id"`
	CommunityName string `gorm:"type:varchar(128)" json:"community_name"`
	Introduction  string `gorm:"type:varchar(256);not null" json:"introduction,omitempty"` // omitempty 当Introduction为空时不展示
	gorm.Model
}

// CommunityDetailRes 社区详情model
type CommunityDetailRes struct {
	CommunityID   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction  string `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	gorm.Model
}
