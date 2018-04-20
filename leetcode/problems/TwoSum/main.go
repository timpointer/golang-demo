package twosum

import "log"

func twoSum(nums []int, target int) []int {
	for i, v1 := range nums {
		for j, v2 := range nums {
			if v1+v2 == target && i != j {
				log.Printf("%v,%v", v1, v2)
				list := []int{i, j}
				return list
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		result := target - v
		log.Printf("result,%v,%v", result, m[result])
		if j, ok := m[result]; ok == true {
			return []int{j, i}
		}
		m[v] = i
		log.Printf("%v,%v", v, m)
	}
	return []int{}
}
