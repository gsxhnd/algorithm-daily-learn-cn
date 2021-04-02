package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if a, ok := m[target-num]; ok {
			return []int{a, i}
		}
		m[num] = i
	}
	return nil
}
