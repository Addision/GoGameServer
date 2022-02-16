package gfservice

// 框架服务器启动停止管理器

type IGFService interface {
	Start()
	Serve()
	Stop()
}
