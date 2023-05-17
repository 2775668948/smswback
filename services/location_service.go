package services

import (
	. "gen/models"
)

type LocationService struct {
}

func NewLocationService() *LocationService {
	return &LocationService{}
}

// GetNewLocation 获取最新的位置数据
func (r LocationService) GetNewLocation() ([]*Location, error) {
	locations, err := GetTheLocation()
	if err != nil {
		return nil, err
	}
	return locations, nil
}
