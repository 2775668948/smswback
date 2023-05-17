package models

import (
	"gen/zlog"
	"time"
)

type DeviceConfig struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	DeviceId  string    `json:"deviceId" gorm:"column:deviceId"`
	Name      string    `json:"name" gorm:"column:name"`
	IconUrl   string    `json:"iconUrl" gorm:"column:iconUrl"`
	Status    uint      `json:"status" gorm:"column:status"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

type UpdateDeviceConfigCommand struct {
	DeviceId string
}

func (DeviceConfig) TableName() string {
	return "deviceconfig"
}

func GetdeviceConfigs(page, pageSize int) ([]*DeviceConfig, error) {
	var deviceConfigs []*DeviceConfig
	err := Db.Limit(pageSize).Offset((page - 1) * pageSize).
		Order("id asc").Find(&deviceConfigs).Error
	if err != nil {
		return nil, err
	}
	return deviceConfigs, nil
}

func GetDeviceById(id string) (*DeviceConfig, error) {
	zlog.Info("根据id查找设备")
	var device DeviceConfig
	err := Db.Where("deviceId = ?", id).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}
