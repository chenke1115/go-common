/*
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-10-21 10:13:31
 * @LastEditTime: 2022-12-07 11:53:12
 * @Description: Do not edit
 */
package file

import (
	"os"
	"path"
	"runtime"
)

/**
 * @description: create dir
 * @param {string} dir
 * @return {*}
 */
func MakeDir(dir string) (err error) {
	if dir == "" {
		dir, _ = os.Getwd()
	}

	err = os.MkdirAll(dir, os.ModePerm)

	return
}

/**
 * @description: Get root path by os
 * @return {*}
 */
func GetRootPath() string {
	rootPath, _ := os.Getwd()
	return rootPath
}

/**
 * @description: Get current ab path By caller
 * @return {*}
 */
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
