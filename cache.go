// Copyright 2019 xuzili1994 Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cache

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var Dir string = "runtime/cache/"
var Error string

type Cache struct {
	Time     int64
	Value    interface{}
	Expires  int64
	PathFile string
}

/**
key get cache
 */
func Get(key string) interface{} {
	fileName := getFilePath(key)
	c := Cache{}
	f, err := os.Open(fileName)
	if err != nil {
		Error = err.Error()
		return nil
	}
	r, _ := ioutil.ReadAll(f)
	json.Unmarshal(r, &c)
	if c.Expires < time.Now().Unix() {
		return nil
	}
	return c.Value
}

/**
结构体解码
get struct decoding
 */
func GetDecoding(key string, value interface{}) bool {
	fileName := getFilePath(key)
	c := Cache{}
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		Error = err.Error()
		return false
	}
	r, _ := ioutil.ReadAll(f)
	json.Unmarshal(r, &c)
	if c.Expires < time.Now().Unix() {
		return false
	}
	json.Unmarshal([]byte(c.Value.(string)), value)
	return true
}

/**
set cache
 */
func Set(key string, value interface{}, timeNum int64) bool {
	pathfile := getFilePath(key)
	dir, _ := path.Split(pathfile)
	if mkdirPath(dir) == false {
		return false
	}
	c := Cache{
		Time:     timeNum,
		Value:    value,
		Expires:  time.Now().Unix() + timeNum,
		PathFile: pathfile,
	}
	if setFile(c) == false {
		return false
	}
	return true
}

/**
set struct encoding
 */
func SetEncoding(key string, value interface{}, timeNum int64) bool {
	pathfile := getFilePath(key)
	dir, _ := path.Split(pathfile)
	if mkdirPath(dir) == false {
		return false
	}
	v, _ := json.Marshal(value)
	c := Cache{
		Time:     timeNum,
		Value:    string(v),
		Expires:  time.Now().Unix() + timeNum,
		PathFile: pathfile,
	}
	if setFile(c) == false {
		return false
	}
	return true
}

/**
key get file name
 */
func genFileName(name string) string {
	hash := md5.New()
	hash.Write([]byte(name))
	resu := hash.Sum(nil)
	return hex.EncodeToString(resu)
}

/**
key get file path
 */
func getFilePath(key string) string {
	fimeName := genFileName(key)
	filePath := Dir + fimeName[:2] + "/" + fimeName[2:] + ".txt"
	return filePath
}

/**
mkdir
 */
func mkdirPath(dir string) bool {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		Error = err.Error()
		return false
	}
	return true
}

/**
set cache file
 */
func setFile(cache Cache) bool {
	c, _ := json.Marshal(cache)
	file, err := os.Create(cache.PathFile)
	defer file.Close()
	if err != nil {
		Error = err.Error()
		return false
	}
	_, err = file.Write(c)
	if err != nil {
		Error = err.Error()
		return false
	}
	return true
}

/**
key is exists
 */
func IsExist(key string) bool {
	filePath := getFilePath(key)
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil && os.IsNotExist(err) {
		return false
	}
	body, _ := ioutil.ReadAll(f)
	c := Cache{}
	json.Unmarshal(body, &c)
	if c.Expires < time.Now().Unix() {
		os.Remove(filePath)
		return false
	}
	return true
}

/**
delete cache file
 */
func Delete(key string) bool {
	filePath := getFilePath(key)
	err := os.Remove(filePath)
	if err != nil {
		Error = err.Error()
		return false
	}
	return true
}
