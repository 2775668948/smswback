package services

import (
	. "gen/models"
)

type DeviceValueService struct{}

func NewDeviceValueService() *DeviceValueService {
	return &DeviceValueService{}
}

// GetAllDeviceValue  根据id返回最新的7条数据
func (r DeviceValueService) GetAllDeviceValue(id string) ([]*DeviceValue, error) {
	deviceValues, err := GetDeviceValues(id)
	if err != nil {
		return nil, err
	}
	return deviceValues, nil
}
