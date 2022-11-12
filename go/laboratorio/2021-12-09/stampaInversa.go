package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var mem []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		l := scanner.Text()
		mem = append(mem, string(l))
	}
	for i := len(mem); i > 0; i-- {
		fmt.Print(mem[i-1], " ")
	}
	fmt.Println()
}
