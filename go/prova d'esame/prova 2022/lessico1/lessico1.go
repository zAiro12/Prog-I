package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)")
	//for {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		// var s string
		// _, err := fmt.Scanln(&s)
		// if err != nil {
		// 	break
		// }
		switch s[0] {
		case '+':
			fmt.Println("alimento il dizionario")
		case '?':
			fmt.Println("consulto il dizionario")
		case 'm':
			fmt.Println("lunghezza min e max")
		case 'p':
			fmt.Println("stampo il dizionario ordinato")
		default:
			fmt.Println("opzione non valida")
		}
	}
	//}
}
