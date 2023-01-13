/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-26 14:19:52
 * @LastEditTime: 2023-01-03 18:27:05
 * @Description: Do not edit
 */
package configs

type Log struct {
	Dir        string `json:"dir" yaml:"Dir"`
	MaxSize    int    `json:"max_size" yaml:"MaxSize"`
	MaxBackups int    `json:"max_backups" yaml:"MaxBackups"`
	MaxAge     int    `json:"max_age" yaml:"MaxAge"`
	Compress   bool   `json:"compress" yaml:"Compress"`
}
