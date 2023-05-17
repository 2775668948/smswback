package models

import "time"

type Heart struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Value     uint      `json:"value" gorm:"column:value"`
	Pid       uint      `json:"pid" gorm:"column:pid"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (Heart) TableName() string {
	return "heart"
}

// GetTheHeart 获取最新的5条数据
func GetTheHeart() ([]*Heart, error) {
	var hearts []*Heart
	err := Db.Order("id desc").Limit(5).Find(&hearts).Error
	if err != nil {
		return nil, err
	}
	return hearts, nil
}
