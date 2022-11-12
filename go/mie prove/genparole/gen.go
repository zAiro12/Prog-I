package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var controllo []string
	for {
		var in string
		_, err := fmt.Scan(&in)
		if err != nil {
			break
		}
		controllo = append(controllo, in)
	}
	file := leggi()
	for i := 0; i < len(file); i++ {
		for _, v := range controllo {
			flag := true
			if file[i] == v {
				flag = false
			}
			if flag {
				scrivi(v)
			}
		}
	}
}

func scrivi(in string) {
	file, err := os.OpenFile(
		"parole.txt",
		os.O_APPEND|os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Println("file not found")
		return
	}
	defer file.Close()
	w := []byte(in)
	w = append(w, '\n')
	_, err = file.Write(w)
}

func leggi() []string {
	file, err := os.Open("parole.txt")
	if err != nil {
		panic("Error while opening the file")
	}
	defer file.Close()
	words := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic("Error while reading")
	}
	return words
}
