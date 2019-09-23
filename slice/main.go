package main

import (
	"fmt"
	"unsafe"
)

const (
	StatusFinish int = 3
	StatusStart  int = 4
)

var (
	sliMapVar = map[int]struct{}{
		StatusFinish: struct{}{},
		StatusStart:  struct{}{},
	}
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	//s1[startIndex:endIndex:capIndex]，其中startIndex为闭区间，endIndex和capIndex为开区间
	s2 := s1[2:6:7]

	//这里append会影响到s1和slice
	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(sliMapVar)

	//in_array的替代方案
	sliMap := map[int]struct{}{
		1: struct{}{},
		2: struct{}{},
	}
	_, ok := sliMap[3]
	fmt.Println(sliMap, ok)

	//Go中的空结构体不占用内存
	fmt.Println(unsafe.Sizeof(struct{}{}))

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}

// Go中没有in_array，可以转成map判断key是否存在
func MapKeyInIntSlice(haystack []int, needle int) bool {
	set := make(map[int]struct{}, len(haystack))

	for _, e := range haystack {
		set[e] = struct{}{}
	}

	_, ok := set[needle]
	return ok
}
