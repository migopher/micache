package cache

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var Path string ="runtime/cache/"

type CachenIterface interface {
	Set()bool
	Get() string
	//IsExists()
}


type File struct {
	Type string
	Expire string
	Prefix string
	Path string
}

func (f File)Get(name string) string {
	return  name
}

func (f File) Set(key string ,value interface{},time int64) bool{
	path:=genPath(key)
	fmt.Println(path)
	return true
}
func genFileName(name string) string {
	hash:=md5.New()
	hash.Write([]byte(name))
	resu:=hash.Sum(nil)
	return hex.EncodeToString(resu)
}
func genPath(name string)string  {
	str:=genFileName(name)
	path:=Path+str[:2]+"/"+str[2:]
	fmt.Println(path)
}




//func Connect() interface{} {
//
//}

//func NewCache() *Cache {
//	//return
//}

//func (c *File)Set(key string, value string,time int64)  {
//
//}
//
//func (c *File )IsExists() bool {
//
//}