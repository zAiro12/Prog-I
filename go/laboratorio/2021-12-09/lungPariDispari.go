package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var r []rune
	var dopo []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		r = []rune(l)
		if len(r)%2 == 0 {
			fmt.Println(l)
		} else {
			dopo = append(dopo, l)
		}
	}
	for i := 0; i < len(dopo); i++ {
		fmt.Println(dopo[i])
	}
}
