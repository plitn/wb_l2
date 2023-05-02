package dev04

import (
	"fmt"
	"strings"
)

// работает для англ слов
func GroupAnagrams(words *[]string) *map[string][]string {
	// создаем мапу, которая считает кол-во каждой буквы в слове
	// все слова у которых такое же количество букв попадут в мапу под этот же ключ
	// потом превращаем эту мапу в словарь
	hmap := make(map[[26]int][]string)
	for _, word := range *words {
		word = strings.ToLower(word)
		sum := [26]int{}
		for i := 0; i < len(word); i++ {
			sum[word[i]-'a']++
		}
		hmap[sum] = append(hmap[sum], word)
	}
	ans := make(map[string][]string)
	for _, v := range hmap {
		key := v[0]
		for _, word := range v {
			ans[key] = append(ans[key], word)
		}
	}
	return &ans
}

func main() {
	data := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(GroupAnagrams(&data))
}
