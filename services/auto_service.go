package services

import (
	. "gen/models"
	"gen/zlog"
)

type AutoTaskService struct{}

func NewAutoTaskService() *AutoTaskService {
	return &AutoTaskService{}
}

// GetAllAutoTask 获取所有的自动化任务
func (r AutoTaskService) GetAllAutoTask(page int) ([]*AutoTask, error) {
	const pageSize = 15
	autotasks, err := GetAutoTasks(page, pageSize)
	if err != nil {
		return nil, err
	}
	return autotasks, nil
}

// UpdateAutoItemState 更改自动化任务开关
func (r AutoTaskService) UpdateAutoItemState(id int) (*AutoTask, error) {
	autoTaskItem, err := GetAutoTaskItemById(id)
	//	先拿到原来的状态
	var oldShow = autoTaskItem.Show
	zlog.Info("原来的状态是------>%d", oldShow)
	//	直接更新
	if oldShow == 0 {
		err2 := Db.Model(&AutoTask{}).Where(autoTaskItem).UpdateColumn("show", 1)
		if err2 != nil {
			return nil, err
		}
	} else {
		err2 := Db.Model(&AutoTask{}).Where(autoTaskItem).UpdateColumn("show", 0)
		if err2 != nil {
			return nil, err
		}
	}
	return autoTaskItem, err
}

// AddNewAutoItem 新增自动化任务
func (r AutoTaskService) AddNewAutoItem(task AutoTask) error {
	zlog.Info("提交新的自动化任务")
	_, err := AddNewAutoItem(task)
	return err
}

// FindAutoItem 根据执行时间去查找自动化任务
func (r AutoTaskService) FindAutoItem(dotime string) (*AutoTask, error) {
	var autoItem AutoTask
	zlog.Debug("当前获取系统时间是-->", dotime)
	err := Db.Where("dotime = ?", dotime).Where("show", 1).First(&autoItem).Error
	if err != nil {
		return nil, err
	}

	// 有的话就执行这个自动化任务
	// 先拿到关联的设备id
	var deviceID = autoItem.DeviceId
	// 这里把类型判断放到服务端来做,也可以放在硬件端去做,只是放在服务端性能更好
	//101 打开照明
	//102 关闭照明
	//103 打开消毒
	//104 关闭消毒
	//105 打开风扇
	//106 关闭风扇
	//107 打开门 这个时候消毒灯关闭照明灯打开
	//108 关闭门
	switch deviceID {
	case "103":
		zlog.Info("设备是照明灯")
		if autoItem.DeviceAc == 1 {
			//		1是关,那就发mqtt消息去关
			zlog.Info("完成了mq信息的发送,设备执行关的指令")
			PubMessage(Message{ID: 102})
		} else {
			//	0是开,那就发mqtt消息去开
			zlog.Info("完成了mq信息的发送,设备执行开的指令")
			PubMessage(Message{ID: 101})
		}
		break
	case "105":
		zlog.Info("设备是消毒灯")
		if autoItem.DeviceAc == 1 {
			//		1是关,那就发mqtt消息去关
			zlog.Info("完成了mq信息的发送,设备执行关的指令")
			PubMessage(Message{ID: 104})
		} else {
			//	0是开,那就发mqtt消息去开
			zlog.Info("完成了mq信息的发送,设备执行开的指令")
			PubMessage(Message{ID: 103})
		}
		break
	case "106":
		zlog.Info("设备是门")
		if autoItem.DeviceAc == 1 {
			//		1是关,那就发mqtt消息去关
			zlog.Info("完成了mq信息的发送,设备执行关的指令")
			PubMessage(Message{ID: 108})
		} else {
			//	0是开,那就发mqtt消息去开
			zlog.Info("完成了mq信息的发送,设备执行开的指令")
			PubMessage(Message{ID: 107})
		}
		break
	case "107":
		zlog.Info("设备是除湿风扇")
		if autoItem.DeviceAc == 1 {
			//		1是关,那就发mqtt消息去关
			zlog.Info("完成了mq信息的发送,设备执行关的指令")
			PubMessage(Message{ID: 106})
		} else {
			//	0是开,那就发mqtt消息去开
			zlog.Info("完成了mq信息的发送,设备执行开的指令")
			PubMessage(Message{ID: 105})
		}
		break
	}
	return &autoItem, err
}
