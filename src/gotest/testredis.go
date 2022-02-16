package main

import (
	. "db"
	. "netnode/gfservice"
	"time"
)

func main() {
	//DbRedis.DelKey("haha")
	// DbRedis.SetString("haha", "55")
	// DbRedis.SetKeyExpire("haha", 5)
	// DbRedis.Incrby("haha", 10)
	// var err error
	// var str string
	// str, err = DbRedis.GetString("haha")
	// if err == nil {
	// 	fmt.Println(str)
	// }

	//DbRedis.HSet("myhash", "aa", "bbb")
	// fmt.Println(DbRedis.HGet("myhash", "cc"))
	//result, _ := DbRedis.HGetAll("FFFF")
	// fmt.Println(DbRedis.HLen("FFFF"))
	//DbRedis.SAdd("hehe", "ssss")
	// fmt.Println(DbRedis.SIsMember("hehe", "ssss"))
	// DbRedis.ZRevRange("runoobkey", 0, 10, true)
	// var len int64
	// len, err = DbRedis.StrLen("haha")
	// if err == nil {
	// 	fmt.Println(len)
	// } else {
	// 	fmt.Println(err)
	// }
	StartRedis()
	key := ServiceFactory.ServerKey() + ".1.userupdate"
	DbRedis.SAdd(key, "ssss")
	DbRedis.SAdd(key, "12341")
	DbRedis.SAdd(key, "34t324")
	DbRedis.SAdd(key, "34t3444444444")
	// memberList, _ := DbRedis.SMembers("v4.dev.1.dataupdate")
	// for _, v := range memberList {
	// 	fmt.Println(v)
	// }
	time.Sleep(time.Second * 1)
}
