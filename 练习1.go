//输入一个字符串，判断字符串是否是回文串，如果是则返回它的重复部分，如果不是则返回false

package main

import "fmt"

func main() {
	var s string
	fmt.Scanln(&s)
	s1 := []rune(s)
	for i := 0; i < len(s1)/2; i++ {
		if s1[i] != s1[len(s1)-1-i] {
			fmt.Println("false")
			break
		} else if i == len(s1)/2-1 {
			switch {
			case i%2 == 0:
				s1 = s1[0 : len(s1)/2]
			case i%2 == 1:
				s1 = s1[0 : len(s1)/2+1]
			}

			s3 := string(s1)
			fmt.Println(s3)
		}

	}
}
