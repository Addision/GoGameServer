// master 用于服务器发现功能
// 1、所有服务器启动时先连接master，如果连接不上，则终止程序。
// 2、其他服务器连接上master，优先上报个人服务器信息，master保留连接上服务器信息
// 3、master将保存的服务器信息，同步给所有已经连接上服务器
package master

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	. "network"
	nodenetwork "network/proto"
)

type ServerInfoMap map[int32]*nodenetwork.ServerReport

type IMasterNodeServer interface {
	// 保存服务器信息
	SaveServer(serverInfo *nodenetwork.ServerReport)
	SyncServer(session ISession)

	InitNetServer(ip string, port int)
}

type MasterNodeServer struct {
	serverMap ServerInfoMap
	// 网络服务
	netService INetService
	// 消息处理
	msgHandler IMsgHandler
	// 继承自 nodeservice
	INodeService
}

func NewMasterNodeServer() *MasterNodeServer {
	server := &MasterNodeServer{
		serverMap:    make(ServerInfoMap),
		msgHandler:   NewMsgHandler(),
		INodeService: NewNodeService(),
	}
	server.SetMsgHandler(server.msgHandler)
	server.AddRouter()
	return server
}

func (s *MasterNodeServer) InitNetServer(ip string, port int) {
	s.netService = NewNetServer(s, ip, port)
}

func (s *MasterNodeServer) AddRouter() {
	s.msgHandler.AddRouter(uint32(nodenetwork.ServerNodeMsgID_REPORT_CLIENT_INFO_TO_SERVER), &ServerRouter{})
}

// define master router
type ServerRouter struct {
	BaseRouter
}

func (router *ServerRouter) Handle(request IRequest) MsgErrCode {
	fmt.Println("master receive msg from server")
	msgID := request.GetMsgID()
	if msgID != uint32(nodenetwork.ServerNodeMsgID_REPORT_CLIENT_INFO_TO_SERVER) {
		fmt.Println("receive err msg, msgID=", msgID)
		return MEC_MSGID
	}
	// 保存信息
	iSession := request.GetSession()
	iNodeService := iSession.GetNodeService()
	master := iNodeService.(*MasterNodeServer)
	dataBytes := request.GetData()
	reportInfo := &nodenetwork.ServerReport{}
	proto.Unmarshal(dataBytes, reportInfo)
	master.SaveServer(reportInfo)
	// 同步所有服务器信息给其他服务器
	master.SyncServer(iSession)
	return MEC_OK
}

func (s *MasterNodeServer) SaveServer(serverInfo *nodenetwork.ServerReport) {
	if _, ok := s.serverMap[serverInfo.ServerId]; ok {
		return
	}
	s.serverMap[serverInfo.ServerId] = serverInfo
}

func (s *MasterNodeServer) SyncServer(session ISession) {
	for _, serverInfo := range s.serverMap {
		msgID := uint32(nodenetwork.ServerNodeMsgID_MASTER_REPORT_SERVER_INFO_TO_SERVER)
		if dataBytes, err := proto.Marshal(serverInfo); err == nil {
			session.SendProtoBuffer(msgID, dataBytes)
		}
	}
}
