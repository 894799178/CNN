package main

import (
	"fmt"
	"github.com/Comdex/imgo"
	"math/rand"
	"utils"
)

func main() {
	fmt.Println("hello")
	//test()
	//	testRangeNumber()
	//look()
	row := 0
	col := 0
	sum := 0
	//该标记用于处理错误的设计导致的bug.
	flag := true
	img1 := imgo.MustRead("resource/3.png")
	img2 := imgo.MustRead("resource/2.png")
	//
	a:for n := 0; n < len(img1); n++ {
		for m := 0; m < len(img1[n]); m++ {
			if utils.CompareArrayAllValue(img1[n][m], img2[col][row]) {
				fmt.Print("1")
				sum++
				if row >= len(img2[col])-1{
					fmt.Print(img1[n][m])
					fmt.Print(img2[col][row])
					row = 0
					flag = false
					col++
					fmt.Println()
					//	img1[n][m], img2[col][row]
				}
				if col >= len(img2)-1{
					fmt.Print("二者匹配")
					break a
				}
				if flag{
					row++
				}else{
					flag = true
				}
			}else{
				fmt.Print(" ")
				row = 0
			}

		}
	}
	fmt.Println("sum-->",sum)
	//fmt.Print(img1)
}



func test(){
	 arr  := make([]int, 0, 1000000)

	a:for i := 2; i <= 1000000; i++ {
		for n := 2; n <= i; n++ {
			if n == i {
				for n:= range arr {
					if n * i == 707829217{
						fmt.Print("结果")
						fmt.Print(n,"-->")
						fmt.Print(i)
						break a
					}
				}
				arr = append(arr, i)
				fmt.Printf("%d ", i)
			}
			if i%n == 0 && n < i {
				break
			}
		}
	}
}

func look(){
	img1 := imgo.MustRead("resource/2.png")
	fmt.Print("img1-->",len(img1))
	fmt.Print("img1[]-->",len(img1[0]))
	//img2 := imgo.MustRead("resource/2.png")
	for n := 0; n < len(img1); n++ {
		for m := 0; m < len(img1[n]); m++ {
			fmt.Print(img1[n][m])
		}
		fmt.Println()
	}
	fmt.Println()
}
func testRangeNumber() {
	rand.Intn(99)
	for i := 0; i < 4; i++ {
		for n := 0; n < 20; n++ {
			num := rand.Intn(100)
			if num <= 9 {
				fmt.Print("0", num)
			} else {
				fmt.Print(num)
			}
			fmt.Print(",")

		}
		fmt.Println()
	}


}
