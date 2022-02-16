package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	player "gotest/testdata"
	"io/ioutil"
)

func main() {
	// data := &player.PlayerData{
	// 	RoleInfo: &player.RoleInfo{
	// 		Name:   "player",
	// 		Age:    22,
	// 		States: make(map[int32]string),
	// 		Temp:   "Temp",
	// 	},
	// }
	// data.RoleInfo.States = nil
	// info, err := proto.Marshal(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//
	// if err := ioutil.WriteFile("./protodata", info, 0644); err != nil {
	// 	log.Fatalln("Failed to write proto:", err)
	// }
	//
	info := &player.PlayerData{}
	in, err := ioutil.ReadFile("./protodata")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = proto.Unmarshal(in, info)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}
