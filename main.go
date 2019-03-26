package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

var Path string = "runtime/cache/"

func main() {
	fmt.Println(getMacAddrs())
	//Set("1234567", 123213, 3600)
}

func Set(key string, value interface{}, time int64) bool {
	path := genPath(key)
	fmt.Println(path)
	return true
}
func genFileName(name string) string {
	hash := md5.New()
	hash.Write([]byte(name))
	resu := hash.Sum(nil)
	return hex.EncodeToString(resu)
}
func genPath(name string) string {
	str := genFileName(name)
	path := Path + str[:2]+"/"
	os.MkdirAll(path,os.ModePerm)
	fmt.Println(path)
	return ""
}

func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	fmt.Println(netInterfaces)
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}
