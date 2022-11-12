package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		appoggio := string(s[i])
		for j := i + 1; j < len(s); j++ {
			appoggio += string(s[j])
			if s[j] == s[i] && len(appoggio) > 2 {
				m[appoggio]++
			}
		}

	}

	var chiavi []string
	for key := range m {
		chiavi = append(chiavi, key)
	}
	//sort.Strings(chiavi)
	for volte := 0; volte < len(chiavi); volte++ {
		for i := 0; i < len(chiavi)-1; i++ {
			if len(chiavi[i+1]) < len(chiavi[i]) {
				chiavi[i+1], chiavi[i] = chiavi[i], chiavi[i+1]
			}
		}
	}
	for i := len(chiavi) - 1; i >= 0; i-- {
		fmt.Println(chiavi[i], "-> Occorrenze:", m[chiavi[i]])
	}
}
