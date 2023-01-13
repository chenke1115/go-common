/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-03 15:47:04
 * @LastEditTime: 2022-12-22 15:04:03
 * @Description: Do not edit
 */
package date

import (
	"strings"
	"time"
)

/**
 * @description: 当前时间戳
 * @return {*}
 */
func DateUnix() int {
	t := time.Now().Local().Unix()
	return int(t)
}

/**
 * @description:当前年月日时分秒(time类型)
 * @return {*}
 */
func DateNowFormat() time.Time {
	tm := time.Now()
	return tm
}

/**
 * @description: 当前时间 Y-m
 * @return {*}
 */
func DateYmFormat() string {
	tm := time.Now()
	return tm.Format("2006-01")
}

/**
 * @description: 当前 Y-m-d
 * @return {*}
 */
func DateYmdFormat() string {
	tm := time.Now()
	return tm.Format("2006-01-02")
}

/**
 * @description: 当前 Y-m-d H:00:00
 * @return {*}
 */
func DateYmdHFormat() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:00:00")
}

/**
 * @description: 当前时间 Y-m-d H:i:s
 * @return {*}
 */
func DateNowFormatStr() string {
	tm := time.Now()
	return tm.Format("2006-01-02 15:04:05")
}

/**
 * @description: 当前第几周
 * @return {*}
 */
func DateWeek() int {
	_, week := time.Now().ISOWeek()
	return week
}

/**
 * @description: 当前年、月、日
 * @return {*}
 */
func DateYMD() (int, int, int) {
	year, month, day := DateYmdInts()
	return year, month, day
}

/**
 * @description: 当前年、月、日
 * @return {*}
 */
func DateYmdInts() (int, int, int) {
	timeNow := time.Now()
	year, month, day := timeNow.Date()
	return year, int(month), day
}

/**
 * @description: 时间戳转 Y-m-d H:i:s
 * @param {int} timestamp
 * @return {*}
 */
func DateFormat(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02 15:04:05")
}

/**
 * @description: 时间戳转 Y-m-d
 * @param {int} timestamp
 * @return {*}
 */
func DateFormatYmd(timestamp int) string {
	tm := time.Unix(int64(timestamp), 0)
	return tm.Format("2006-01-02")
}

/**
 * @description: 日期time.Time格式化
 * @param {string} date exp:2022-12-01
 * @return {*}
 */
func DateFormatFormStr(date string) (t time.Time, err error) {
	timeFormatTpl := "20060102150405"
	if strings.Contains(date, "-") {
		timeFormatTpl = "2006-01-02 15:04:05"
	} else if strings.Contains(date, "/") {
		timeFormatTpl = "2006/01/02 15:04:05"
	}

	if len(timeFormatTpl) != len(date) {
		timeFormatTpl = timeFormatTpl[0:len(date)]
	}

	// 解析日期
	t, err = time.Parse(timeFormatTpl, date)
	return
}

/**
 * @description: 日期格式互转
 * @param {string} date 日期字符串 exp:2022-12-01
 * @param {...string} args 转换目标格式 exip:2022/12/01
 * @return {*}
 */
func DateFormatTransfers(date string, args ...string) string {
	// 解析日期
	t, err := DateFormatFormStr(date)
	if err != nil {
		return ""
	}

	// 输出日期格式
	format := "2006-01"
	if len(args) > 0 {
		format = args[0]
	}
	return t.Format(format)
}

/**
 * @description: 上一年的月份
 * @param {string} moth exp:2022-01|202201
 * @param {...string} args exp:2006-01|200601
 * @return {*}
 */
func LastYearMoth(moth string, args ...string) string {
	// 解析日期
	t, err := DateFormatFormStr(moth)
	if err != nil {
		return ""
	}

	// 输出日期格式
	format := "2006-01"
	if len(args) > 0 {
		format = args[0]
	}
	return t.AddDate(-1, 0, 0).Format(format)
}
