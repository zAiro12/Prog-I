package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var c int
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Split(bufio.ScanRunes)
	for myScanner.Scan() {
		l := myScanner.Text()
		for _, d := range l {
			if d == 'à' || d == 'è' || d == 'ì' || d == 'ò' || d == 'ù' || d == 'é' {
				c++
			}
		}
	}
	fmt.Println(c)
}
