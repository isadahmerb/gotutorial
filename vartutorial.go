package main

import "fmt"

func main() {
	intArr := [3]int32{1, 2, 3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The leght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n the legth is %v with capacity %v", len(intSlice), cap(intSlice))

}
