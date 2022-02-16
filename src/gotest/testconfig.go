package main

// import (
// 	"fmt"
// 	"utils"
// )
//
// type ServerBaseConfig struct {
// 	NodeId     int
// 	NodeName   string
// 	NodeIP     string
// 	NodePort   int
// 	MaxConnect int
// }
//
// type MasterConfig struct {
// 	ServerBaseConfig
// }
//
// type GameConfig struct {
// 	ServerBaseConfig
// }
//
// type GateConfig struct {
// 	ServerBaseConfig
// }
//
// type WorldConfig struct {
// 	ServerBaseConfig
// }
//
// type LoginConfig struct {
// 	ServerBaseConfig
// }
//
// type ChatConfig struct {
// 	ServerBaseConfig
// }
//
// type GateUserConfig struct {
// 	ServerBaseConfig
// }
//
// type MysqlConfig struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	Password string
// 	DataBase string
// }
//
// type RedisConfig struct {
// 	Host     string
// 	Port     int
// 	Password string
// }
//
// type Config struct {
// 	MasterServer   MasterConfig
// 	WorldServer    WorldConfig
// 	LoginServer    LoginConfig
// 	GateServer     GateConfig
// 	GameServer     GameConfig
// 	ChatServer     ChatConfig
// 	GateUserServer GateUserConfig
//
// 	Mysql MysqlConfig
// 	Redis RedisConfig
// }
//
// func NewConfig() *Config {
// 	return &Config{}
// }
//
// func main() {
// 	configure := utils.NewConfigure()
// 	configure.Load("./config/server.json")
// 	stConfig := NewConfig()
// 	if err := configure.AssignStruct(stConfig); err == nil {
// 		fmt.Println("stConfig==", stConfig.ChatServer)
// 	} else {
// 		fmt.Println("err==", err)
// 	}
// }
