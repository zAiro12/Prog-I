package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scanln(&s)
	flag := true
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			flag = false
			break
		}
	}
	if flag {
		mappa := Sottostringhe(s)
		var chiavi []string
		for k := range mappa {
			chiavi = append(chiavi, k)
		}
		sort.Strings(chiavi)
		fmt.Println("output:")
		for i := range chiavi {
			chiave := chiavi[i]
			fmt.Println(chiave, mappa[chiave])
		}
	}
}

func Sottostringhe(s string) map[string]int {
	mappa := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] > s[j-1] {
				substring := s[i : j+1]
				mappa[substring]++
			} else {
				break
			}
		}
	}
	return mappa
}
