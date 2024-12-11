package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	result := GroupAnagrams(strs)

	fmt.Println("result:", result)
}

func GroupAnagrams(strs []string) [][]string {
	// 创建一个map，用于将排序后的字符串映射到字母异位词组
	anagramsMap := make(map[string][]string)

	// 遍历输入字符串数组
	for _, str := range strs {
		// 将字符串转换为字符切片
		chars := []rune(str)
		// 对字符切片进行排序
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		// 将排序后的字符切片重新组合成字符串
		sortedStr := string(chars)
		// 将排序后的字符串作为键，将原始字符串加入对应的组
		anagramsMap[sortedStr] = append(anagramsMap[sortedStr], str)
	}

	fmt.Println("anagramsMap:", anagramsMap)
	// 将字母异位词组从map中提取出来
	var result [][]string
	for _, group := range anagramsMap {
		result = append(result, group)
	}

	return result
}
