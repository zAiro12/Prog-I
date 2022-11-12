package main

import (
	"fmt"
	"sort"
)

func main() {
	var stringa string
	fmt.Scan(&stringa)
	if isBalanced(stringa) {
		fmt.Println("bilanciata")
	} else {
		fmt.Println("non bilanciata")
	}
	mappa := sub(stringa)
	var chiavi []string
	for k := range mappa {
		chiavi = append(chiavi, k)
	}
	sort.Slice(chiavi, func(i, j int) bool { return len(chiavi[i]) < len(chiavi[j]) })

	for _, v := range chiavi {
		fmt.Printf("%s (%d) \n", v, mappa[v])

	}
}

func isBalanced(stringa string) bool {
	c := 0
	for i := 0; i < len(stringa); i++ {
		if stringa[i] == '(' {
			c++
		} else {
			c--
		}
		if c < 0 {
			return false
		}
	}
	if c == 0 {
		return true
	} else {
		return false
	}
}

func sub(stringa string) (m map[string]int) {
	mappa := make(map[string]int)
	for i := 0; i < len(stringa); i++ {
		for j := i + 1; j <= len(stringa); j++ {
			s := stringa[i:j]
			if isBalanced(s) {
				mappa[s]++
			}
		}
	}
	return mappa
}
