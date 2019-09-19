package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	//s1[startIndex:endIndex:capIndex]，其中startIndex为闭区间，endIndex和capIndex为开区间
	s2 := s1[2:6:7]

	//这里append会影响到s1和slice
	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
