package logic

import (
	"SnailForum/common"
	"SnailForum/config"
	"SnailForum/model"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

func GetCommunityCategory() []model.Category {
	db := config.GetDB()
	category := make([]model.Category, 0)
	db.Where(&model.Category{}).Find(&category)
	return category
}

func GetCategoryById(cid int32) model.Category {
	db := config.GetDB()
	var category model.Category
	db.Where(cid).Find(&category)
	return category
}

// GetCommunityList 查询分类社区列表
func GetCommunityList() (communityList []*model.Community, err error) {
	db := config.GetDB()
	err = db.Select("community_id, community_name").Find(&communityList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return communityList, err
}

func GetCommunityNameByID(idStr string) (community *model.Community, err error) {
	db := config.GetDB()
	community = new(model.Community)
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	err = db.Select("community_id, community_name").Where("community_id = ?", id).First(community).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("query community failed", zap.Error(err))
			return nil, errors.New(common.ErrorInvalidID)
		}
		zap.L().Error("query community failed", zap.Error(err))
		return nil, errors.New(common.ErrorQueryFailed)
	}
	return community, nil
}

// GetCommunityByID 根据ID查询分类社区详情
func GetCommunityByID(id uint64) (*model.CommunityDetailRes, error) {
	db := config.GetDB()
	community := new(model.CommunityDetail)
	err := db.Select("community_id, community_name, introduction, create_time").Where("community_id = ?", id).First(community).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("query community failed", zap.Error(err))
			return nil, errors.New(common.ErrorInvalidID)
		}
		zap.L().Error("query community failed", zap.Error(err))
		return nil, errors.New(common.ErrorQueryFailed)
	}
	return &model.CommunityDetailRes{
		CommunityID:   community.CommunityID,
		CommunityName: community.CommunityName,
		Introduction:  community.Introduction,
	}, nil
}
