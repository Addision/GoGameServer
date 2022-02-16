package gf

import (
	. "network"
	nodenetwork "network/proto"
)

type IBaseServer interface {
	Start()
	Serve()
	Stop()

	ReadConfig() // 读取配置：读取redis mysql配置
	StartLog()   // 启动日志管理

	SetOnStart(onStart func() bool)
	SetOnStop(onStop func() bool)

	// 设置服务器信息，用于广播给其他节点
	SetServerInfo(serverID int32, serverName string, serverIP string, serverPort int32, maxConnect int32, serverState nodenetwork.ServerState, serverType int32)
	GetServerInfo() nodenetwork.ServerReport
	AddRouter(msgID int32, router IRouter)
}

// GFNodeServer 包含服务器端和客户端
type BaseServer struct {
	// 服务器启动关闭
	onStart func() bool
	onStop  func() bool
	// 本服务器详细信息
	ServerInfo nodenetwork.ServerReport
}

func NewBaseServer() *BaseServer {
	return &BaseServer{
		onStart:    nil,
		onStop:     nil,
		ServerInfo: nodenetwork.ServerReport{},
	}
}

// 启动服务器
func (s *BaseServer) Start() {
	s.onStart()
}

// 关闭服务器
func (s *BaseServer) Stop() {
	s.onStop()
}

// 运行服务器执行
func (s *BaseServer) Serve() {

}

// 读取配置
func (s *BaseServer) ReadConfig() {

}

//
func (s *BaseServer) SetOnStart(onStart func() bool) {
	if onStart != nil {
		s.onStart = onStart
	}
}

func (s *BaseServer) SetOnStop(onStop func() bool) {
	if onStop != nil {
		s.onStop = onStop
	}
}

func (s *BaseServer) StartLog() {

}

func (s *BaseServer) SetServerInfo(serverID int32, serverName string, serverIP string, serverPort int32, maxConnect int32, serverState nodenetwork.ServerState, serverType int32) {
	s.ServerInfo.ServerId = serverID
	s.ServerInfo.ServerName = []byte(serverName)
	s.ServerInfo.ServerIp = []byte(serverIP)
	s.ServerInfo.ServerPort = serverPort
	s.ServerInfo.MaxOnline = maxConnect
	s.ServerInfo.CurOnline = 0
	s.ServerInfo.ServerState = serverState
	s.ServerInfo.ServerType = serverType
}

func (s *BaseServer) GetServerInfo() nodenetwork.ServerReport {
	return s.ServerInfo
}

func (s *BaseServer) AddRouter(msgID int32, router IRouter) {

}
