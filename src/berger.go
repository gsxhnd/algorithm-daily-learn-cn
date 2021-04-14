package main

import "fmt"

func main() {
	parseAryB([]int{1, 2, 3, 4, 5})
	parseAryB([]int{1, 2, 3, 4, 5, 6, 7, 8})
}

func parseAryB(srcAry []int) {
	retSlice := make([][]int, 0)
	var round int

	// 判断单双数
	if len(srcAry)%2 != 0 {
		srcAry = append(srcAry, 0)
	}
	round = len(srcAry) - 1
	var sortAry = make([]int, len(srcAry))
	copy(sortAry, srcAry)

	retSlice = append(retSlice, srcAry)
	for i := 1; i < round; i++ {
		var roundData []int
		if i%2 != 0 {
			start := sortAry[len(srcAry)/2+1 : len(srcAry)-1]
			end := sortAry[:len(srcAry)/2]
			roundData = append(roundData, sortAry[len(srcAry)-1])
			roundData = append(roundData, start...)
			roundData = append(roundData, end...)
			roundData = append(roundData, sortAry[len(srcAry)/2])
		} else {
			start := sortAry[len(srcAry)/2+1:]
			end := sortAry[1 : len(srcAry)/2]
			roundData = append(roundData, sortAry[len(srcAry)/2])
			roundData = append(roundData, start...)
			roundData = append(roundData, end...)
			roundData = append(roundData, sortAry[0])
		}
		copy(sortAry, roundData)
		retSlice = append(retSlice, roundData)
	}

	for _, slc := range retSlice {
		fmt.Println(slc)
	}
	for rid, slc := range retSlice {
		nowLen := len(slc) / 2
		left := slc[:nowLen]
		right := revSlc(slc[nowLen:])
		fmt.Println("round ", rid)
		for eidx, _ := range left {
			fmt.Println(left[eidx], "-", right[eidx])
		}
	}
}

func revSlc(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
