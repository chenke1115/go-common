/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-18 14:27:07
 * @LastEditTime: 2022-11-18 14:28:56
 * @Description: Do not edit
 */
package configs

type User struct {
	Super    []string `json:"super" yaml:"Super"`
	Password Password `json:"password" yaml:"Password"`
}

type Password struct {
	Salt string `json:"salt" yaml:"Salt"`
	Init string `json:"init" yaml:"Init"`
}
