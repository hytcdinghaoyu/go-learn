package main

import (
	"fmt"
)

func bubbleSort(sli []int) []int {
	length := len(sli)

	for i := 1; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}

	return sli
}

func selectSort(sli []int) []int {
	length := len(sli)

	for i := 0; i < length; i++ {
		k := i
		for j := i + 1; j < length-1; j++ {
			if sli[k] > sli[j] {
				k = j
			}
		}

		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}

	}

	return sli
}

func quickSort(sli []int) []int {
	length := len(sli)

	if length <= 1{
		return sli
	}

	base := sli[0]
	leftSli := []int{}
	rightSli := []int{}

	for i := 1; i < length; i++ {
		if sli[i] < base{
			leftSli = append(leftSli, sli[i])
		}else {
			rightSli = append(rightSli, sli[i])
		}
	}

	leftSli = quickSort(leftSli)
	rightSli = quickSort(rightSli)

	leftSli = append(leftSli, base)
	return append(leftSli, rightSli...)

}

func main() {

	var sli = []int{66, 3, 45, 7, 14, 23, 33, 78}

	fmt.Println(quickSort(sli))
}
