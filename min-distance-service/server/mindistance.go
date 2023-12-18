package server

import (
	"math/rand"
)

func sort(arr []int) []int {

	var isSorted = false

	for !isSorted {
		isSorted = true

		var i = 0
		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
			i++
		}
	}

	return arr
}

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	l, r := 0, len(arr)-1

	p := rand.Int() % len(arr)

	arr[p], arr[r] = arr[r], arr[p]

	for i := range arr {
		if arr[i] < arr[r] {
			arr[l], arr[i] = arr[i], arr[l]
			l++
		}
	}

	arr[l], arr[r] = arr[r], arr[l]

	quickSort(arr[:l])
	quickSort(arr[l+1:])

	return arr
}

func minDistance(arr []int) [][]int {
	arr = quickSort(arr)
	res := [][]int{{arr[0], arr[1]}}
	min := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < min {
			min = diff
			res = [][]int{{arr[i], arr[i+1]}}
		} else if diff == min {
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}

// func main() {
// 	//var arr = []int{1, 49, 3, 9, 27, 15}
// 	var arr = []int{4, 3, 2, 1, 8}
// 	fmt.Println(quickSort(arr))
// 	fmt.Println(minDistance(arr))
// }
