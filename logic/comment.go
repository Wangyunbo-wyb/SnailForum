package logic

import (
	"SnailForum/config"
	"SnailForum/model"
)

func CreateComment(comment *model.Comment) error {
	db := config.GetDB()
	if err := db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func GetCommentListByIDs(ids []uint64) (commentList []*model.Comment, err error) {
	db := config.GetDB()
	err = db.Where("comment_id IN ?", ids).Find(&commentList).Error
	return
}
