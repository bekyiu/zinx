package util

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

// zinx 全局配置
// zinx.json暴露给用户

type GlobalObj struct {
	server ziface.IServer
	Host   string
	Port   int
	Name   string

	Version    string
	MaxConn    int    // 允许的最大连接数
	MaxPkgSize uint32 // 一个数据包data区域的最大值
}

var GlobalConfig *GlobalObj

func init() {
	GlobalConfig = &GlobalObj{
		server:     nil,
		Host:       "0.0.0.0",
		Port:       8999,
		Name:       "ZinxServer",
		Version:    "V0.4",
		MaxConn:    3,
		MaxPkgSize: 4096,
	}
	// 用户自定义配置
	GlobalConfig.Load()
}

func (g *GlobalObj) Load() {
	// todo 相对路径
	config, err := ioutil.ReadFile("/Users/bekyiu/dev/goProjects/zinx/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(config, &GlobalConfig)
}
