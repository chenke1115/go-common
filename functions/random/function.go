/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-10 15:05:50
 * @LastEditTime: 2022-12-07 10:45:21
 * @Description: Do not edit
 */
package random

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

/**
 * @description: get random string with length
 * @param {int} l
 * @return {*}
 */
func GetRandomString(l int, args ...string) string {
	str := args[0]
	if str == "" {
		str = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/**
 * @description: 字符串补零
 * @param {string} str 需要操作的字符串
 * @param {int} resultLen 结果字符串的长度
 * @param {bool} reverse true 为前置补零，false 为后置补零
 * @return {*}
 */
func ZeroFillByStr(str string, resultLen int, reverse bool) string {
	if len(str) > resultLen || resultLen <= 0 {
		return str
	}
	if reverse {
		return fmt.Sprintf("%0*s", resultLen, str) //不足前置补零
	}

	result := str
	for i := 0; i < resultLen-len(str); i++ {
		result += "0"
	}

	return result
}

/**
 * @description: Get number from str
 * @param {string} str
 * @return {*}
 */
func GetNumByStr(str string) (strArr []string) {
	re := regexp.MustCompile("[0-9]+")
	strArr = re.FindAllString(str, -1)
	return
}
