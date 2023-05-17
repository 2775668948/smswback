package models

import (
	"time"
)

type Threshold struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Did       uint      `json:"did" gorm:"column:did"`
	Uint      string    `json:"uint" gorm:"column:uint"`
	Top       string    `json:"top" gorm:"column:top"`
	Low       string    `json:"low" gorm:"column:low"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (Threshold) TableName() string {
	return "threshold"
}

// GetTheThreshold 获取最新的2条数据
func GetTheThreshold() ([]*Threshold, error) {
	var theres []*Threshold
	err := Db.Order("id desc").Limit(2).Find(&theres).Error
	if err != nil {
		return nil, err
	}
	return theres, nil
}

//
//func GetDeviceById(id string) (*DeviceConfig, error) {
//	zlog.Info("根据id查找设备")
//	var device DeviceConfig
//	err := Db.Where("deviceId = ?", id).First(&device).Error
//	if err != nil {
//		return nil, err
//	}
//	return &device, nil
//}
