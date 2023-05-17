package models

import "time"

// Location 位置model
type Location struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Longitude string    `json:"longitude" gorm:"column:longitude"`
	Latitude  string    `json:"latitude" gorm:"column:latitude"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (Location) TableName() string {
	return "location"
}

// GetTheLocation 获取最新的一条数据
func GetTheLocation() ([]*Location, error) {
	var locations []*Location
	err := Db.Order("id desc").Limit(1).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}
