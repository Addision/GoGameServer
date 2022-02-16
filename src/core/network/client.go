package network

type IClient interface {
}

type Client struct {
	netService INetService
}
