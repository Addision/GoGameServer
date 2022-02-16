package game

import (
	. "gf"
	. "global"
	. "global/config"
	. "network"
	node "network/proto"
)

type GameServer struct {
	GFNodeServer *GFNodeServer
	config       *Config
}

func NewSceneServer() *GameServer {
	return &GameServer{
		GFNodeServer: NewGFNodeServer(),
		config:       NewConfig(),
	}
}

func (s *GameServer) OnStart() {
	// 设置路由
	s.AddRouter()
	s.SetServerInfo()
	s.AddConnectServers()
}

func (s *GameServer) OnStop() {

}

// gate send msg to game
func (s *GameServer) onGateRouteToGame() {

}

// world send msg to game
func (s *GameServer) onWorldRouteToGame() {

}

// chat send msg to game
func (s *GameServer) onChatRouteToGame() {

}

// game send msg to gate
func (s *GameServer) SendToGate(serverID int32, playerID uint64, msgID int32, msg []byte) {
	gatePacket := &node.GameToGatePacket{}
	gatePacket.PlayerId = playerID
	gatePacket.MsgId = msgID
	gatePacket.MsgBody = msg
	s.GFNodeServer.NodeClient.SendPbByServerID(serverID, int32(node.ServerNodeMsgID_GAME_ROUTE_TO_GATE), gatePacket)
}

// game send msg to world
func (s *GameServer) SendToWorld(serverID int32, playerID uint64, msgID int32, msg []byte) {
	worldPacket := &node.GameToWorldPacket{}
	worldPacket.PlayerId = playerID
	worldPacket.MsgId = msgID
	worldPacket.MsgBody = msg
	s.GFNodeServer.NodeClient.SendPbByServerID(serverID, int32(node.ServerNodeMsgID_GAME_ROUTE_TO_WORLD), worldPacket)
}

// set server info and report to other servers
func (s *GameServer) SetServerInfo() {
	// 获取本服配置信息
	info := s.config.GameServer
	s.GFNodeServer.SetReportInfo(int32(info.NodeId), info.NodeName, info.NodeIP, int32(info.NodePort), int32(info.MaxConnect), node.ServerState_SS_NORMAL, int32(SERVER_TYPE_GAME))
}

// 启动服务器添加作为客户端需要连接的服务器
func (s *GameServer) AddConnectServers() {
	// gameserver connect maser world gate DB
	s.GFNodeServer.NodeClient.AddConnectServer(SERVER_TYPE_MASTER)
	s.GFNodeServer.NodeClient.AddConnectServer(SERVER_TYPE_GATE)
	s.GFNodeServer.NodeClient.AddConnectServer(SERVER_TYPE_WORLD)
	s.GFNodeServer.NodeClient.AddConnectServer(SERVER_TYPE_CHAT)
}

// define game router
// 处理gate发送给game消息
type OnGateRouter struct {
	BaseRouter
}

func (router *OnGateRouter) Handle(request IRequest) {

}

// 处理world消息
type OnWorldRouter struct {
	BaseRouter
}

func (router *OnWorldRouter) Handle(request IRequest) {

}

// 处理聊天消息
type OnChatRouter struct {
	BaseRouter
}

func (router *OnChatRouter) Handle(request IRequest) {

}

func (s *GameServer) AddRouter() {

}
