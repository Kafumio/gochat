package server

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"gochat/common/constant"
	"gochat/config"
	"gochat/pkg/kafka"
	"gochat/pkg/log"
	"gochat/pkg/protocal"
)

type Client struct {
	Conn *websocket.Conn
	Name string
	Send chan []byte
}

func (c *Client) Read() {
	defer func() {
		MyServer.Ungister <- c
		c.Conn.Close()
	}()

	for {
		c.Conn.PongHandler()
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			log.Logger.Error("client read message error", log.Any("client read message error", err.Error()))
			MyServer.Ungister <- c
			c.Conn.Close()
			break
		}

		msg := &protocal.Message{}
		proto.Unmarshal(message, msg)

		// pong
		if msg.Type == constant.HEAT_BEAT {
			pong := &protocal.Message{
				Content: constant.PONG,
				Type:    constant.HEAT_BEAT,
			}
			pongByte, err := proto.Marshal(pong)
			if err != nil {
				log.Logger.Error("client marshal message error", log.Any("client marshal message error", err.Error()))
			}
			c.Conn.WriteMessage(websocket.BinaryMessage, pongByte)
		} else {
			if config.GetConfig().MsgChannelType.ChannelType == constant.KAFKA {
				kafka.Send(message)
			} else {
				MyServer.Broadcast <- message
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		c.Conn.Close()
	}()

	for message := range c.Send {
		c.Conn.WriteMessage(websocket.BinaryMessage, message)
	}
}
