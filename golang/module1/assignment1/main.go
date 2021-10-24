package main

import "fmt"

// 给定字符串数组
// ['I',‘am','stupid','and','weak']
// 用for循环遍历该数组，并且修改为
// ['I',"am","smart","and","strong"]
func main() {
	mySlice := []string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("mySlice %+v\n", mySlice)

	for i, v := range mySlice {
		if v == "stupid" {
			mySlice[i] = "smart"
		} else if v == "weak" {
			mySlice[i] = "strong"
		}
	}
	fmt.Printf("mySlice %+v\n", mySlice)

}
