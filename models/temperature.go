package models

import "time"

type Temperature struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Value     uint      `json:"value" gorm:"column:value"`
	Pid       uint      `json:"pid" gorm:"column:pid"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (Temperature) TableName() string {
	return "temperature"
}

// GetTheTemperature 获取最新的5条数据
func GetTheTemperature() ([]*Temperature, error) {
	var tempers []*Temperature
	err := Db.Order("id desc").Limit(5).Find(&tempers).Error
	if err != nil {
		return nil, err
	}
	return tempers, nil
}
