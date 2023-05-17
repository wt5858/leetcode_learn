package main

import "fmt"

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	res := countCharacters(words, chars)
	fmt.Println(res)
}

func countCharacters(words []string, chars string) int {
	var charCount [26]uint8
	for _, v := range chars {
		charCount[v-'a']++
	}

	res := 0
	var flag bool
	for _, word := range words {
		flag = true
		var wordCount [27]uint8
		for _, w := range word {
			wordCount[w-'a']++
			if wordCount[w-'a'] > charCount[w-'a'] {
				flag = false
				break
			}
		}
		if flag {
			res += len(word)
		}
	}
	return res
}
