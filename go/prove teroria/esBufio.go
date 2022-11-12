package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var mem string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		mem = l
	}
	fmt.Println(mem)
}
