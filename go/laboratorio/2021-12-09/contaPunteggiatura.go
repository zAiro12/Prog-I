package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mappa := make(map[string]int)
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Split(bufio.ScanRunes)
	for myScanner.Scan() {
		l := myScanner.Text()
		if l[0] == '.' || l[0] == ',' || l[0] == ';' || l[0] == ':' || l[0] == '!' || l[0] == '?' || l[0] == '"' || l[0] == '\'' {
			mappa[l]++
		}

	}
	for k, v := range mappa {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
}
