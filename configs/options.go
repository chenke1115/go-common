/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-15 15:32:42
 * @LastEditTime: 2022-11-18 14:42:08
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
}

/**
 * @description: default
 * @return {*}
 */
func defaultOptions() (option Options) {
	return
}
