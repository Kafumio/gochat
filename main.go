package main

import (
	"gochat/common/constant"
	"gochat/config"
	"gochat/pkg/kafka"
	"gochat/pkg/log"
	"gochat/router"
	"gochat/websocket/server"
	"net/http"
	"time"
)

func main() {
	log.InitLogger(config.GetConfig().Log.Path, config.GetConfig().Log.Level)
	log.Logger.Info("config", log.Any("config", config.GetConfig()))

	if config.GetConfig().MsgChannelType.ChannelType == constant.KAFKA {
		kafka.InitProducer(config.GetConfig().MsgChannelType.KafkaTopic, config.GetConfig().MsgChannelType.KafkaHosts)
		kafka.InitConsumer(config.GetConfig().MsgChannelType.KafkaHosts)
		go kafka.ConsumerMsg(server.ConsumerKafkaMsg)
	}

	log.Logger.Info("start server", log.String("start", "start web sever..."))

	newRouter := router.InitRouter()

	go server.MyServer.Start()

	s := &http.Server{
		Addr:           ":8888",
		Handler:        newRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Logger.Error("server error", log.Any("serverError", err))
	}
}
