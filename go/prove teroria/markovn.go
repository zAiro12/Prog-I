package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type lm map[string][]rune

func readAll(k int) lm {
	var result lm
	var curr rune
	var prev string

	result = make(lm)
	prev = ""
	for i := 0; i < k; i++ {
		var c rune
		fmt.Scanf("%c", &c)
		prev += string(c)
	}
	for {
		_, err := fmt.Scanf("%c", &curr)
		if err != nil {
			break
		}
		result[prev] = append(result[prev], curr)
		prev = string([]rune(prev)[1:]) + string(curr)
	}
	return result
}

func printMap(m lm) {
	for k, v := range m {
		fmt.Printf("%s: ", k)
		for _, r := range v {
			fmt.Printf("'%c' ", r)
		}
		fmt.Println()
	}
}

func getRandomKey(m lm) string {
	var k string
	n := rand.Intn(len(m))
	for k, _ = range m {
		if n == 0 {
			return k
		}
		n--
	}
	return k
}

func getRandomValue(m lm, prev string) string {
	_, exists := m[prev]
	if exists {
		return string(m[prev][rand.Intn(len(m[prev]))])
	} else {
		return getRandomKey(m)
	}
}

func main() {
	k, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[2])
	rand.Seed(time.Now().UnixNano())
	m := readAll(k)
	//printMap(m)

	last := getRandomKey(m)
	fmt.Printf("%s", last)
	for i := 0; i < n; i++ {
		lastAsSlice := []rune(last)
		t := getRandomValue(m, string(lastAsSlice[len(lastAsSlice)-k:]))
		fmt.Printf("%s", t)
		last += t
	}
	fmt.Println()
}
