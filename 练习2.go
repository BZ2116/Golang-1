//使用Go标准库中的随机数生成函数，生成100个整数并添加（append）到一个数组中。

package main

import (
	"fmt"
	"math/rand"
	"time"
)
//冒泡排序
func main() {
	var number []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		number = append(number, rand.Intn(1000))
	}
	fmt.Println("before: ", number)
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(number)-i-1; j++ {
			if number[j] > number[j+1] {
				number[j], number[j+1] = number[j+1], number[j]
			}
		}

	}
	fmt.Println("after:", number)
}
