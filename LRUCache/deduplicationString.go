package main

import "fmt"

//无重复字符的最长子串：
//
//题目描述：给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
//示例：
//go
//复制代码
//输入：s = "abcabcbb"
//输出：3
//解释：答案是 "abc"，长度为 3。

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	var maxLength, start int

	for i := 0; i < len(s); i++ {
		fmt.Println("v:", s[i])
		fmt.Println("charMap:", charMap[s[i]])
		if index, ok := charMap[s[i]]; ok && index >= start {
			start = index + 1
		}
		charMap[s[i]] = i
		if maxLength < i-start+1 {
			maxLength = i - start + 1
		}
	}

	return maxLength
}

func main() {
	input := "abcabcbb"
	fmt.Println("Input:", input)
	output := lengthOfLongestSubstring(input)
	fmt.Println("Output:", output)
}
