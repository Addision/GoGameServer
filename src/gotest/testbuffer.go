package main

import (
	"fmt"
	"utils"
)

func main() {
	// buffer := utils.NewSingleBuffer()
	// bytes := []byte("ertgergegedaedsd")
	// buffer.Write(bytes)
	//
	// fmt.Println(buffer.Len())
	// fmt.Println(buffer.DataLen())
	// fmt.Println(buffer.ReaderOffset())
	// fmt.Println(buffer.WriterOffset())
	//
	// bytes = []byte("ertgergegedaedsd")
	// buffer.Write(bytes)
	//
	// fmt.Println(buffer.Len())
	// fmt.Println(buffer.DataLen())
	// fmt.Println(buffer.ReaderOffset())
	// fmt.Println(buffer.WriterOffset())
	//
	// buffer.Read(10)
	// fmt.Println(buffer.Len())
	// fmt.Println(buffer.DataLen())
	// fmt.Println(buffer.ReaderOffset())
	// fmt.Println(buffer.WriterOffset())
	//
	// buffer.Read(15)
	// fmt.Println(buffer.Len())
	// fmt.Println(buffer.DataLen())
	// fmt.Println(buffer.ReaderOffset())
	// fmt.Println(buffer.WriterOffset())
	//
	// bytes = []byte("ertgergegedaedsd")
	// buffer.Write(bytes)
	//
	// fmt.Println(buffer.Len())
	// fmt.Println(buffer.DataLen())
	// fmt.Println(buffer.ReaderOffset())
	// fmt.Println(buffer.WriterOffset())

	buffer := utils.NewBufferList()
	bytes := []byte("ertgergegedaedsd")
	buffer.Write(bytes)
	buffer.Write(bytes)
	// buffer.ReadAll()
	data, len := buffer.Read(10)
	fmt.Println(string(data), len)
	data, len = buffer.Read(10)
	fmt.Println(string(data), len)
	data, len = buffer.Read(10)
	fmt.Println(string(data), len)
	data, len = buffer.Read(10)
	fmt.Println(string(data), len)
}
