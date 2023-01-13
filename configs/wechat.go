/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-12-23 15:37:57
 * @LastEditTime: 2022-12-26 14:10:41
 * @Description: Do not edit
 */
package configs

type Wechat struct {
	AppId       string    `json:"app_id" yaml:"AppId"`
	AppSecret   string    `json:"app_secret" yaml:"AppSecret"`
	AccessToken string    `json:"access_token" yaml:"AccessToken"`
	Tpl         []string  `json:"tpl" yaml:"Tpl"`
	Url         WechatUrl `json:"url" yaml:"Url"`
}

type WechatUrl struct {
	RedirectUri string `json:"redirect" yaml:"Redirect"`
	H5Uri       string `json:"h5" yaml:"H5"`
}
