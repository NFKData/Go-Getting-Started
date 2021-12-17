package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	sliceFromArr := arr[:]
	fmt.Println(&sliceFromArr[0], sliceFromArr)
	arr[2] = 27
	fmt.Println(sliceFromArr)
	sliceFromArr = append(sliceFromArr, 27)
	fmt.Println(&sliceFromArr[0], sliceFromArr)

	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 27)
	fmt.Println(slice)

	firstElementSliceFromArr := arr[0:1]
	secondElementSliceFromArr := arr[1:2]
	lastElementSliceFromArr := arr[2:3]
	twoElementsSliceFromArr := arr[0:2]
	fmt.Println(firstElementSliceFromArr, secondElementSliceFromArr, lastElementSliceFromArr, twoElementsSliceFromArr)

}
