/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-12-12 15:32:01
 * @LastEditTime: 2022-12-13 14:49:00
 * @Description: Do not edit
 */
package date

import (
	"time"
)

/**
 * @description: 根据开始日期和结束日期计算出时间段内所有小时整点
 * @param {*} stime exp:2022-01-01 00:00:00
 * @param {string} etime exp:2022-01-01 24:00:00
 * @param {...string} args exp: 15:04:05
 * @return {*}
 */
func GetBetweenHours(stime, etime string, args ...string) (hours []string) {
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(stime) {
		timeFormatTpl = timeFormatTpl[0:len(stime)]
	}

	// 解析开始日期
	st, err := time.Parse(timeFormatTpl, stime)
	if err != nil {
		return
	}
	// 解析结束日期
	et, err := time.Parse(timeFormatTpl, etime)
	if err != nil {
		return
	}
	// 比较开始结束日期
	if et.Before(st) {
		return
	}

	// 输出日期格式
	format := "2006-01-02 15:04:05"
	if len(args) > 0 {
		format = args[0]
	}

	etStr := et.Format(format)
	hours = append(hours, st.Format(format))
	for {
		st = time.Date(st.Year(), st.Month(), st.Day(), st.Hour()+1, 0, 0, 0, time.Local)
		stStr := st.Format(format)
		hours = append(hours, stStr)
		if stStr == etStr {
			break
		}
	}

	return
}

/**
 * @description: 根据开始日期和结束日期计算出时间段内所有日期
 * @param {*} sdate 如：2020-01-01
 * @param {string} edate 如：2020-01-01
 * @param {string} fomart 如：2020-01-01
 * @return {*}
 */
func GetBetweenDates(sdate, edate string, args ...string) (dates []string) {
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}

	// 解析开始日期
	sd, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		return
	}
	// 解析结束日期
	ed, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		return
	}
	// 比较开始结束日期
	if ed.Before(sd) {
		return
	}

	// 输出日期格式
	format := "2006-01-02"
	if len(args) > 0 {
		format = args[0]
	}

	edStr := ed.Format(format)
	dates = append(dates, sd.Format(format))
	for {
		sd = sd.AddDate(0, 0, 1)
		sdStr := sd.Format(format)
		dates = append(dates, sdStr)
		if sdStr == edStr {
			break
		}
	}

	return
}

/**
 * @description: 根据开始日期和结束日期计算出时间段内所有月份
 * @param {*} sMoth exp:2022-01
 * @param {string} eMoth exp:2022-01
 * @param {...string} args exp:2006-01|200601
 * @return {*}
 */
func GetBetweenMoths(sMoth, eMoth string, args ...string) (moths []string) {
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sMoth) {
		timeFormatTpl = timeFormatTpl[0:len(sMoth)]
	}
	// 解析开始日期
	sm, err := time.Parse(timeFormatTpl, sMoth)
	if err != nil {
		return
	}
	// 解析结束日期
	em, err := time.Parse(timeFormatTpl, eMoth)
	if err != nil {
		return
	}
	// 比较开始结束日期
	if em.Before(sm) {
		return
	}

	// 输出日期格式
	format := "2006-01"
	if len(args) > 0 {
		format = args[0]
	}

	emStr := em.Format(format)
	moths = append(moths, sm.Format(format))
	for {
		sm = sm.AddDate(0, 1, 0)
		smStr := sm.Format(format)
		moths = append(moths, smStr)
		if smStr == emStr {
			break
		}
	}

	return
}

/**
 * @description: 获取日期当前月份的第一天和最后一天
 * @param {string} date exp: 2022-01
 * @param {...string} args exp:2006-01|200601
 * @return {*}
 */
func GetMonthFristAndLast(date string, args ...string) map[string]string {
	timeFormatTpl := "2006-01-02 15:04:05"
	if args[0] == "200601" {
		timeFormatTpl = "20060102150405"
	}
	if len(timeFormatTpl) != len(date) {
		timeFormatTpl = timeFormatTpl[0:len(date)]
	}

	// 解析日期
	d, err := time.Parse(timeFormatTpl, date)
	if err != nil {
		return map[string]string{}
	}

	year := d.Year()
	month := d.Month()
	frist := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	last := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Format("2006-01-02")
	result := map[string]string{"frist": frist, "last": last}
	return result
}
