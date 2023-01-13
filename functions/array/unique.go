/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-11-17 15:18:50
 * @LastEditTime: 2022-11-17 15:20:41
 * @Description: Do not edit
 */
package array

/**
 * @description: Duplicate removal
 * @param {[]int} arr
 * @return {*}
 */
func UniqueArray(arr []int) []int {
	newArr := make([]int, 0)
	tempArr := make(map[int]bool, len(newArr))
	for _, v := range arr {
		if !tempArr[v] {
			tempArr[v] = true
			newArr = append(newArr, v)
		}
	}

	return newArr
}
