package main

import (
	"fmt"
	"strings"
)

//{"from":"face002","name":"yitianjian","to":"wx_001"}
var a []string = []string{"from","to","name"}
var rawLen int
var combineLen int

func main() {
	rawLen = len(a)
	combineLen = 2//取两个进行组合
	combine()
}

func combine() {
	arrLen := len(a) - combineLen //3 - 2
	for i := 0; i <= arrLen; i++ {
		result := make([]string, combineLen)
		result[0] = a[i]
		doProcess(result, i, 1)
	}
}

func doProcess(result []string, rawIndex int, curIndex int) {
	var choice int = rawLen - rawIndex + curIndex - combineLen
	//fmt.Printf("Choice: %d, rawLen: %d, rawIndex : %d, curIndex : %d \r\n", choice, rawLen, rawIndex, curIndex)
	var tResult []string
	for i := 0; i < choice; i++ {
		if i != 0 {
			tResult := make([]string, combineLen)
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

func copyArr(rawArr []string, target []string) {
	for i := 0; i < len(rawArr); i++ {
		target[i] = rawArr[i]
	}
}

func PrintIntArr(arr []string) {
	//valuesText := []string{}
	//for i := range arr {
	//	number := arr[i]
	//	text := strconv.Itoa(number)
	//	valuesText = append(valuesText, text)
	//}

	fmt.Println(strings.Join(arr, ","))
}

