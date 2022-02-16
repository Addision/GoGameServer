package master

import (
	. "gf"
	. "global/config"
	nodenetwork "network/proto"
)

type MasterServer struct {
	masterServer IMasterNodeServer
	IBaseServer
}

func NewMasterServer() *MasterServer {
	master := &MasterServer{
		masterServer: NewMasterNodeServer(),
		IBaseServer:  NewBaseServer(),
	}
	// 获取本服信息
	config := GConfig.MasterServer
	master.SetServerInfo(int32(config.NodeId),
		config.NodeName,
		config.NodeIP,
		int32(config.NodePort),
		int32(config.MaxConnect),
		nodenetwork.ServerState_SS_NORMAL,
		1)
	master.masterServer.InitNetServer(config.NodeIP, config.NodePort)
	return master
}
