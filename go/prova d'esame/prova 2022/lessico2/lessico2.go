package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	PrintMenu()
	scanner := bufio.NewScanner(os.Stdin)
	dict := make(map[string][]int)
	c := 0
	for scanner.Scan() {
		s := scanner.Text()
		switch s[0] {
		case '+':
			c++
			fmt.Println("alimento il dizionario")
			AddToDict(s, c, dict)
		case '?':
			fmt.Println("consulto il dizionario")
			fmt.Println("parola:", s[2:])
			fmt.Println("righe", dict[s[2:]])
		case 'm':
			fmt.Println("lunghezza min e max")
			fmt.Println(Stats(dict))
		case 'p':
			fmt.Println("stampo il dizionario ordinato")
			PrintDict(dict)
		default:
			fmt.Println("opzione non valida")
		}
	}
}

func PrintMenu() {
	fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")
}

func PrintDict(m map[string][]int) {
	var chiavi []string
	for k := range m {
		chiavi = append(chiavi, k)
	}
	sort.Strings(chiavi)
	for i := 0; i < len(chiavi); i++ {
		fmt.Println(chiavi[i], ":", m[chiavi[i]])
	}
}

func AddToDict(line string, n int, dict map[string][]int) {
	split := strings.Split(line[2:], " ")
	for i := 0; i < len(split); i++ {
		dict[split[i]] = append(dict[split[i]], n)
	}
}

func Stats(dict map[string][]int) (int, int) {
	var chiavi []string
	for k := range dict {
		chiavi = append(chiavi, k)
	}
	sort.Slice(chiavi, func(i int, j int) bool { return len(chiavi[i]) < len(chiavi[j]) })
	return len(chiavi[0]), len(chiavi[len(chiavi)-1])
}
