package models

import (
	"gen/zlog"
	"time"
)

type AutoTask struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"column:name"`
	DoTime    string    `json:"DoTime" gorm:"column:dotime"`
	DeviceId  string    `json:"deviceId" gorm:"column:deviceId"`
	DeviceAc  uint      `json:"deviceAc" gorm:"column:deviceAc"`
	Show      uint      `json:"show" gorm:"column:show"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt"`
}

func (AutoTask) TableName() string {
	return "autotask"
}

func GetAutoTasks(page, pageSize int) ([]*AutoTask, error) {
	var autoTasks []*AutoTask
	err := Db.Limit(pageSize).Offset((page - 1) * pageSize).
		Order("id asc").Find(&autoTasks).Error
	if err != nil {
		return nil, err
	}
	return autoTasks, nil
}

// GetAutoTaskItemById 根据id获取自动化任务
func GetAutoTaskItemById(id int) (*AutoTask, error) {
	zlog.Info("根据主键id查找自动化任务")
	var autoItem AutoTask
	err := Db.Where("id = ?", id).First(&autoItem).Error
	if err != nil {
		return nil, err
	}
	return &autoItem, nil
}

// AddNewAutoItem 新增自动化任务
func AddNewAutoItem(AutoTask AutoTask) (*AutoTask, error) {
	zlog.Info("新创建自动化任务")
	err := Db.Create(&AutoTask)
	if err != nil {

	}
	return &AutoTask, nil
}
