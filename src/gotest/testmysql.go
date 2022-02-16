package main

import (
	. "db"
	"fmt"
)

func main() {

	//fmt.Println("filePath:", Config.ConfigPath)
	// if err := DbMysql.Connect(); err != nil {
	// 	fmt.Println("connect mysql error, ", err.Error())
	// }
	StartMysql()
	//db.DbMysql.Close()
	if DbMysql.IsConnected() {
		fmt.Println("mysql connection ok")
	} else {
		fmt.Println("mysql connection closed")
	}

}
