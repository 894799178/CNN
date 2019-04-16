package main

import (
	"fmt"
	"github.com/Comdex/imgo"
	"math/rand"
	"service"
)

func main() {
	fmt.Println("hello")
	//test()
	//	testRangeNumber()
	//look()
	service.ContainsChildImage("resource/1.png", "resource/2.png")
	//service.Screenshots(5,5,500,500,1)

}

func test() {
	arr := make([]int, 0, 1000000)
a:
	for i := 2; i <= 1000000; i++ {
		for n := 2; n <= i; n++ {
			if n == i {
				for n := range arr {
					if n*i == 707829217 {
						fmt.Print("结果")
						fmt.Print(n, "-->")
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

func look() {
	img1 := imgo.MustRead("resource/2.png")
	fmt.Print("img1-->", len(img1))
	fmt.Print("img1[]-->", len(img1[0]))
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
