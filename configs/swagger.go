/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-25 14:40:56
 * @LastEditTime: 2023-08-01 16:37:28
 * @Description: Do not edit
 */
package configs

type Swagger struct {
	Version  string   `json:"version" yaml:"Version"`
	Host     string   `json:"host" yaml:"Host"`
	BasePath string   `json:"base_path" yaml:"BasePath"`
	Schemes  []string `json:"schemes" yaml:"Schemes"`
	Title    string   `json:"title" yaml:"Title"`
}
