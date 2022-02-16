package gfservice

import (
	. "global"
	. "logger"
	"network"
)

var ServiceFactory NodeServiceFactory

type NodeServiceFactory struct {
}

func (nsf *NodeServiceFactory) NewNodeService(serviceType ServiceType, serverType ServerType) IGFService {
	if ServerConf == nil {
		return nil
	}
	servConfig := ServerConf.GetNodeConfig(serverType)
	if servConfig == nil {
		LogError("get server config failed")
		return nil
	}
	if serviceType == ST_SERVER {
		return network.NewNodeServer(servConfig)
	}
	if serviceType == ST_CLIENT {
		return network.NewNodeClient(servConfig)
	}
	return nil
}
