/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */
package main

// @lc code=start
func reverseVowels(s string) string {
	res := []rune(s)
	i, j := 0, len(s)-1
	for i < j {
		c1,c2 := rune(s[i]),rune(s[j])
		if isVowel(c1) && isVowel(c2) {
			res[i],res[j]=res[j],res[i]
			i++
			j--
		}
		if !isVowel(c1) {
			i++
		}
		if !isVowel(rune(c2)) {
			j--
		}
	}
	return string(res)
}

func isVowel(a rune) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, v := range vowels {
		if a == v {
			return true
		}
	}
	return false
}

// @lc code=end
