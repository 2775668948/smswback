package services

import (
	. "gen/models"
)

type HeartService struct {
}

func NewHeartService() *HeartService {
	return &HeartService{}
}

// GetNewHeartValue 获取最新的位置数据
func (r HeartService) GetNewHeartValue() ([]*Heart, error) {
	hearts, err := GetTheHeart()
	if err != nil {
		return nil, err
	}
	return hearts, nil
}
