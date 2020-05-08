package main

import (
	"fmt"
	"strconv"
	"strings"
)

//@link https://www.cnblogs.com/adoontheway/p/6047918.html
//@todo 排列组合

var a []int = []int{1, 2, 3, 4, 5, 6, 7}
var rawLen int
var combineLen int

func main() {
	rawLen = len(a)
	combineLen = 4
	combine()
}

func combine() {
	fmt.Println("start combine")
	arrLen := len(a) - combineLen
	for i := 0; i <= arrLen; i++ {
		result := make([]int, combineLen)
		result[0] = a[i]
		//fmt.Println(i, arrLen)
		doProcess(result, i, 1)
	}
}

func doProcess(result []int, rawIndex int, curIndex int) {
	var choice int = rawLen - rawIndex + curIndex - combineLen
	//fmt.Printf("Choice: %d, rawLen: %d, rawIndex : %d, curIndex : %d \r\n", choice, rawLen, rawIndex, curIndex)
	var tResult []int
	for i := 0; i < choice; i++ {
		if i != 0 {
			tResult := make([]int, combineLen)
			copyArr(result, tResult)
		} else {
			tResult = result
		}
		//fmt.Println(curIndex)
		tResult[curIndex] = a[i+1+rawIndex]

		if curIndex+1 == combineLen {
			PrintIntArr(tResult)
			continue
		} else {
			doProcess(tResult, rawIndex+i+1, curIndex+1)
		}

	}
}

func copyArr(rawArr []int, target []int) {
	for i := 0; i < len(rawArr); i++ {
		target[i] = rawArr[i]
	}
}

func PrintIntArr(arr []int) {
	valuesText := []string{}
	for i := range arr {
		number := arr[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	fmt.Println(strings.Join(valuesText, ","))
}