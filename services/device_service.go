package services

import (
	"fmt"
	_ "gen/middleware"
	. "gen/models"
	"gen/zlog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/goccy/go-json"
)

type DeviceService struct{}

type Message struct {
	ID int `json:"id"`
}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

// GetAll 获取所有的设备
func (r DeviceService) GetAll(page int) ([]*DeviceConfig, error) {
	const pageSize = 15
	deviceConfigs, err := GetdeviceConfigs(page, pageSize)
	if err != nil {
		return nil, err
	}
	return deviceConfigs, nil
}

var PubMessage = func(message Message) {

	zlog.Info("将数据推送到硬件")
	opts := mqtt.NewClientOptions().AddBroker("tcp://47.106.159.110:1883")
	client := mqtt.NewClient(opts)

	// 连接MQTT代理
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	// 发布一条消息
	text, err := json.Marshal(message)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	token := client.Publish("swpub", 0, false, text)
	token.Wait()
}

// UpdateDeviceSw 更改设备开关
func (r DeviceService) UpdateDeviceSw(deviceId string) (*DeviceConfig, error) {
	deviceConfig, err := GetDeviceById(deviceId)
	// 先拿到原来的状态
	var oldState = deviceConfig.Status
	zlog.Info("原来的状态是--->%d", oldState)
	// 更新

	if oldState == 0 {
		// 更新成打开

		switch deviceId {
		case "103":
			zlog.Info("是照明灯")
			message := Message{ID: 101}
			PubMessage(message)
			break
		case "105":
			zlog.Info("是消毒灯")
			message := Message{ID: 103}
			PubMessage(message)
			break
		case "106":
			zlog.Info("是门")
			message := Message{ID: 107}
			PubMessage(message)
			break
		case "107":
			zlog.Info("是除湿风扇")
			message := Message{ID: 105}
			PubMessage(message)
			break
		}

		zlog.Info("更新成打开")
		err2 := Db.Model(&DeviceConfig{}).Where(deviceConfig).UpdateColumn("status", 1)
		if err2 != nil {
			return nil, err
		}

	} else {

		switch deviceId {
		case "103":
			zlog.Info("是照明灯")
			message := Message{ID: 102}
			PubMessage(message)
			break
		case "105":
			zlog.Info("是消毒灯")
			message := Message{ID: 104}
			PubMessage(message)
			break
		case "106":
			zlog.Info("是门")
			message := Message{ID: 108}
			PubMessage(message)
			break
		case "107":
			zlog.Info("是除湿风扇")
			message := Message{ID: 106}
			PubMessage(message)
			break
		}
		zlog.Info("更新成关闭")
		err2 := Db.Model(&DeviceConfig{}).Where(deviceConfig).UpdateColumn("status", 0)
		if err2 != nil {
			return nil, err
		}

	}
	//// 然后将mqtt命令推送下去
	if err != nil {
		return nil, err
	}
	return deviceConfig, err
}
