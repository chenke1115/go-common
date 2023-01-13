/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-15 15:32:42
 * @LastEditTime: 2022-12-23 15:40:54
 * @Description: Do not edit
 */
package configs

type Options struct {
	Debug    bool     `json:"debug" yaml:"Debug"`
	App      App      `json:"app" yaml:"App"`
	Server   Sever    `json:"server" yaml:"Server"`
	Log      Log      `json:"log" yaml:"Log"`
	Database Database `json:"database" yaml:"Database"`
	Redis    Redis    `json:"redis" yaml:"Redis"`
	Session  Session  `json:"session" yaml:"Session"`
	Swagger  Swagger  `json:"swagger" yaml:"Swagger"`
	Upload   Upload   `json:"upload" yaml:"Upload"`
	Wechat   Wechat   `json:"wechat" yaml:"Wechat"`
}

/**
 * @description: default
 * @return {*}
 */
func defaultOptions() (option Options) {
	return
}
