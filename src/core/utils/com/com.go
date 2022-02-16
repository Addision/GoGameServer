package com

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/OneOfOne/xxhash"
	"hash/crc32"
	. "logger"
	"net"
	"runtime"
	"runtime/debug"
	"strconv"
	"time"
)

type Func func()
type FuncArgs func(args ...interface{})

func TimeLoopArgs(idleTime time.Duration, f FuncArgs, args ...interface{}) {
	loop := func() {
		defer func() {
			err := recover()
			if err != nil {
				LogError("func TimeLoop error, %s\n stack:%s", err, debug.Stack())
			}
		}()
		f(args...)
	}
	go func() {
		for {
			loop()
			time.Sleep(idleTime)
		}
	}()
}

func TimeLoop(idleTime time.Duration, f Func) {
	loop := func() {
		defer func() {
			err := recover()
			if err != nil {
				LogError("func TimeLoop error, %s\n stack:%s", err, debug.Stack())
			}
		}()
		f()
	}
	go func() {
		for {
			loop()
			time.Sleep(idleTime)
		}
	}()
}

func GoRunArgs(f FuncArgs, args ...interface{}) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				LogError("func GoRunArgs error, %s\n stack:%s", err, debug.Stack())
				return
			}
		}()
		f(args...)
	}()
}

func GoRun(f Func) {
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				LogError("func GoRun error, %s\n stack:%s", err, debug.Stack())
				return
			}
		}()
		f()
	}()
}

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string {
	c := sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// RS Hash Function
func RSHash(str []byte) uint32 {
	var b uint32 = 378551
	var a uint32 = 63689
	var hash uint32 = 0
	for i := 0; i < len(str); i++ {
		hash = hash*a + uint32(str[i])
		a *= b
	}
	return hash
}

func Hash(input []byte) uint64 {
	h := xxhash.New64()
	h.Write(input)
	return h.Sum64()
}

func GetLocalIPV4() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "unknown"
}

// get go routine id
func GoID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func init() {

}
