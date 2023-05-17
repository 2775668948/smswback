package services

import (
	. "gen/models"
)

type TemperatureService struct {
}

func NewTemperatureService() *TemperatureService {
	return &TemperatureService{}
}

// GetNewTemperature 获取最新的位置数据
func (r TemperatureService) GetNewTemperature() ([]*Temperature, error) {
	temps, err := GetTheTemperature()
	if err != nil {
		return nil, err
	}
	return temps, nil
}
