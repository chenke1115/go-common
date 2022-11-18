/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-18 14:25:12
 * @LastEditTime: 2022-11-18 14:30:29
 * @Description: Do not edit
 */
package configs

type App struct {
	Name string `json:"name" yaml:"Name"`
	User User   `json:"user" yaml:"User"`
}
