package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"gen/config"
	"gen/middleware"
	"gen/models"
	"gen/router"
	"gen/zlog"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "app.ini", "config file path")
	flag.Parse()

	// 加载配置
	cfg, err := config.Load(configFile)
	if err != nil {
		panic(fmt.Sprintf("load config failed, file: %s, error: %s", configFile, err))
	}

	// 初始化日志
	zlog.Init(cfg)
	defer zlog.GetLogger().Sync()

	// 启动mqtt服务
	//go startMqtt(cfg)

	// 初始化数据库
	err = models.Init(cfg)
	if err != nil {
		zlog.Panic("Init db failed, error: %s", err)
	}

	// 启动Web服务
	zlog.Info("Server starting...")
	err = startServer(cfg)
	if err != nil {
		zlog.Panic("Server started failed: %s", err)
	}
}

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic %s: %s\n", message.Topic(), message.Payload())
}

func startMqtt(cfg *config.AppConfig) {
	zlog.Info("mqtt init...")
	opts := mqtt.NewClientOptions().AddBroker("tcp://47.106.159.110:1883").SetClientID("jjjjj")

	// 设置回调函数，用于接收订阅的消息
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	})

	// 创建 MQTT 客户端连接
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅指定的 MQTT 主题
	topic := "swsub"
	qos := 0
	if token := client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

}

func startServer(cfg *config.AppConfig) error {
	server := &http.Server{
		Addr:    ":" + cfg.HttpPort,
		Handler: getEngine(cfg),
	}
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctxFunc context.CancelFunc) {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
		for {
			select {
			case <-signalChan:
				ctxFunc()
				return
			}
		}
	}(cancel)
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			zlog.Error("Failed to shutdown server: %s", err)
		}
	}()
	zlog.Debug("Server started success")
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		zlog.Debug("Server was shutdown gracefully")
		return nil
	}
	return err
}

func getEngine(cfg *config.AppConfig) *gin.Engine {
	gin.SetMode(func() string {
		if cfg.IsDevEnv() {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	engine := gin.New()
	engine.Use(middleware.Cors())
	engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "服务器内部错误，请稍后再试！",
		})
	}))
	router.RegisterRoutes(engine)
	return engine
}
