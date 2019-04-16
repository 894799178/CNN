/*
@Time : 2019/4/16 11:40
@Author : Tester
@File : 一条小咸鱼
@Software: GoLand
*/
package service

import (
	"syscall"
)

var (
	mode32        = syscall.NewLazyDLL("lib/jietu.dll")
	procGetScreen = mode32.NewProc("getScreen")
)

func Screenshots(left, top, right, bottom, path int64) {
	//utils.StringToUintptr(path)
	getScreen(left, top, right, bottom, path)
}

func getScreen(left, top, right, bottom, path int64) {
	_, _, _ = procGetScreen.Call(uintptr(left), uintptr(top), uintptr(right), uintptr(bottom), uintptr(path))
}
