package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var daCercare, file string
	if len(os.Args) < 3 {
		fmt.Println("args missing")
		return
	} else {
		daCercare = os.Args[1]
		file = os.Args[2]
	}
	if file == "file_uno" {
		file += ".txt"
	} else if file == "file_due" {
		file += ".txt"
	} else {
		return
	}
	k, err := os.Open(file)
	if err != nil {
		fmt.Println("file not found")
		return
	}
	defer k.Close()
	fileScanner := bufio.NewScanner(k)
	var c int
	flag := false
	for fileScanner.Scan() {
		l := fileScanner.Text()
		if strings.Contains(l, daCercare) {
			fmt.Println("found at", c)
			flag = true
		}
		c++
	}
	

}
