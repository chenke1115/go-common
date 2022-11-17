/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-30 10:50:33
 * @LastEditTime: 2022-11-17 15:27:00
 * @Description: Do not edit
 */
package array

import (
	"sort"
)

/**
 * @description: Judge the value in the array
 * @param {interface{}} val
 * @param {interface{}} array
 * @return {*}
 */
func In(val interface{}, array interface{}) bool {
	// type of []string
	if strArr, ok := array.([]string); ok {
		// string
		str, _ := val.(string)
		return InArray(str, strArr)
	}

	// type of []int
	if intArr, ok := array.([]int); ok {
		i, _ := val.(int)
		return InIntArray(i, intArr)
	}

	// type of map[string]interface{}
	if strArr, ok := array.(map[string]interface{}); ok {
		str, _ := val.(string)
		return InMapArray(str, strArr)
	}

	return false
}

/**
 * @description: Judge str in strArr
 * @param {string} str
 * @param {[]string} strArr
 * @return {*}
 */
func InArray(str string, strArr []string) bool {
	// sort
	sort.Strings(strArr)
	// judge
	index := sort.SearchStrings(strArr, str)
	if index < len(strArr) && strArr[index] == str {
		return true
	}

	return false
}

/**
 * @description: []int
 * @param {int} item
 * @param {[]int} items
 * @return {*}
 */
func InIntArray(item int, items []int) bool {
	for _, val := range items {
		if val == item {
			return true
		}
	}

	return false
}

/**
 * @description: Judge str in mapArr
 * @param {string} str
 * @param {map[string]interface{}} mapArr
 * @return {*}
 */
func InMapArray(str string, mapArr map[string]interface{}) bool {
	for _, v := range mapArr {
		if v == str {
			return true
		}
	}

	return false
}