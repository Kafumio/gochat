# gochat
go-chat.使用Go基于WebSocket开发的web聊天应用。单聊，群聊。文字，图片，语音，视频消息，屏幕共享，剪切板图片，基于WebRTC的P2P语音通话，视频聊天。

## 功能列表：
- 登录注册
- 修改头像
- 群聊天
- 群好友列表
- 单人聊天
- 添加好友
- 添加群组
- 文本消息
- 剪切板图片
- 图片消息
- 文件发送
- 语音消息
- 视频消息
- 屏幕共享（基于图片）
- 视频通话（基于WebRTC的p2p视频通话）
- 分布式部署（通过kafka全局消息队列，统一消息传递，可以水平扩展系统）

## 技术栈
- web框架Gin
- 长连接WebSocket
- 日志框架Uber的zap
- 配置管理viper
- ORM框架gorm
- 通讯协议Google的proto buffer
- 消息队列kafka
- makefile 的编写
- 数据库MySQL
- 图片文件二进制操作

## 快速运行
### 运行go程序
go环境的基本配置
...

拉取后端代码
```shell
git clone https://github.com/kafumio/gochat
```

进入目录
```shell
cd gochat
```

拉取程序所需依赖
```shell
go mod download
```

MySQL创建数据库
```mysql
CREATE DATABASE chat;
```

修改数据库配置文件
```shell
vim config.toml

[mysql]
host = "127.0.0.1"
name = "chat"
password = "root1234"
port = 3306
table_prefix = ""
user = "root"

修改用户名user，密码password等信息。
```

创建表
```shell
将chat.sql里面的sql语句复制到控制台创建对应的表。
```

在user表里面添加初始化用户
```shell
手动添加用户。
```

运行程序
```shell
go run main.go
```