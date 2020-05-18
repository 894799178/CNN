/*
@Time : 2019/4/15 11:08
@Author : 一条小咸鱼
@File : image_service
@Software: GoLand
 主要用于提供图像识别的服务
*/
package service

import (
	"cnn/src/utils"
	"fmt"
	"github.com/Comdex/imgo"
)

/**
图像识别,
准确的说应该是判断主图片是否包含子图片;
更准确的说应该是从主矩形中判断是否包含某个其他小于该矩形的矩形
无法识别非矩形图片
*/
func ContainsChildImage(mainFileName string, childFileName string) bool {

	//该标记用于处理错误的设计导致的bug.
	//flag := true
	img1 := imgo.MustRead(mainFileName)
	img2 := imgo.MustRead(childFileName)
	for n := 0; n < len(img1); n++ {
		for m := 0; m < len(img1[n]); m++ {
			//比较像素点 如果是一致的则返回true
			if ComparePixels(img1[n][m], img2[0][0]) {
				if childImageCompare(img1, img2, m, n) {
					fmt.Print("二者匹配")
					return true
				}
				//	//判断是否已经到达当行像素点的最后一个
				//	if row >= len(img2[col])-1{
				//		fmt.Print(img1[n][m])
				//		fmt.Print("n-->",n," m-->",m)
				//		fmt.Println(img2[col][row])
				//		row = 0
				//		flag = false
				//		col++
				//	}
				//	if col >= len(img2)-1{
				//		fmt.Print("二者匹配sum-->",sum)
				//		//匹配则跳出所有循环
				//		return true
				//	}
				//	//处理因为设计问题导致的bug
				//	//保证切换比较下一行的时候两边的都是从第一个开始比较
				//	if flag{
				//		row++
				//	}else{
				//		flag = true
				//	}
				//}else{
				//	//排除某些凑巧一致的像素点
				//	row = 0
			}
		}
	}
	return false
}

func childImageCompare(mainArr, childArr [][][]uint8, mainX, mainY int) bool {
	sum := 0
	for y := 0; y < len(childArr); y++ {
		for x := 0; x < len(childArr[y]); x++ {
			if ComparePixels(mainArr[y+mainY][x+mainX], childArr[y][x]) {
				sum++
				fmt.Print("1")
			} else {
				return false
			}
		}
		fmt.Println()
	}
	return true
}

/**
	比较像素点,像素点按照数组的形式传入
	位置 0 : Red
    位置 1 : Green
    位置 2 : Blue
    位置 3 : 透明值
*/
func ComparePixels(arr1 []uint8, arr2 []uint8) bool {
	return utils.CompareArrayAllValue(arr1, arr2)
}
