/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-18 14:38:51
 * @LastEditTime: 2022-11-18 14:41:22
 * @Description: Do not edit
 */
package configs

type Session struct {
	Driver string `json:"driver" yaml:"Driver"`
	Name   string `json:"name" yaml:"Name"`
	Secret string `json:"secret" yaml:"Secret"`
}
