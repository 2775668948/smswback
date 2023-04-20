package models

import "time"

type DeviceValue struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	DeviceId  string    `json:"deviceId" gorm:"column:deviceId"`
	Value     string    `json:"value" gorm:"column:value"`
	Unit      string    `json:"unit" gorm:"column:unit"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (DeviceValue) TableName() string {
	return "devicevalue"
}

// GetDeviceValues 根据id返回最后的7条数据
func GetDeviceValues(id string) ([]*DeviceValue, error) {
	var deviceValues []*DeviceValue
	err := Db.Where("deviceId = ?", id).Order("id desc").Limit(4).Find(&deviceValues).Error
	if err != nil {
		return nil, err
	}
	return deviceValues, nil
}
