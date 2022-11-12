package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var mem []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		mem = append(mem, string(l))
	}
	for i := 1; i < len(mem); i += 2 {
		fmt.Println(mem[i])

	}
	for i := 0; i < len(mem); i += 2 {
		fmt.Println(mem[i])
	}

}
