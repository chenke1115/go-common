/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-09-30 10:50:33
 * @LastEditTime: 2023-05-25 16:01:01
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
	switch arr := array.(type) {
	case []string:
		v, _ := val.(string)
		return InArray(v, arr)
	case []int, []int8, []int16, []int32, []int64, []uint8, []uint16, []uint32, []uint64:
		return InIntArray(val, arr)
	case float32, float64:
		return InFloatArray(val, arr)
	case map[int]interface{}, map[string]interface{}:
		return InMapArray(val, arr)
	default:
		return false
	}
}

// 字符型数组包含元素判断
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

// 整型数组包含元素判断
func InIntArray(val interface{}, array interface{}) bool {
	switch arr := array.(type) {
	case []int:
		v, _ := val.(int)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []int8:
		v, _ := val.(int8)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []int16:
		v, _ := val.(int16)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []int32:
		v, _ := val.(int32)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []int64:
		v, _ := val.(int64)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []uint8:
		v, _ := val.(uint8)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []uint16:
		v, _ := val.(uint16)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []uint32:
		v, _ := val.(uint32)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []uint64:
		v, _ := val.(uint64)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

// 浮点数组包含元素判断
func InFloatArray(val interface{}, array interface{}) bool {
	switch arr := array.(type) {
	case []float32:
		v, _ := val.(float32)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case []float64:
		v, _ := val.(float64)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}

// Map数组包含元素判断
func InMapArray(val interface{}, array interface{}) bool {
	switch arr := array.(type) {
	case map[string]interface{}:
		v, _ := val.(string)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	case map[int]interface{}:
		v, _ := val.(int)
		for _, item := range arr {
			if v == item {
				return true
			}
		}
	default:
		return false
	}

	return false
}
