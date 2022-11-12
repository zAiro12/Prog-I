/* package main

import "fmt"

func main() {
	var x, i, j int
	fmt.Print("dimensione v: ")
	fmt.Scan(&x)
	g := 2*x - 3
	for i = x; i > 0; i-- { //<--- numero di righe
		for j = 0; j < x-i; j++ { //<--- numero di spazi inizio righe
			fmt.Print(" ")
		}
		fmt.Print("*")              //<--- back slash
		for k := g; k > 0; k -= 2 { //<--- spazi dentro la V
			fmt.Print(" ")
		}
		fmt.Print("*")
		fmt.Println()
	}
} */

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Inserire il numero di righe: ")
	fmt.Scan(&n)
	k := 2*n - 3 //funzione che calcola il numero degli spazi nella prima riga
	for i := 0; ; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		for j := 0; j < k; j++ {
			fmt.Print(" ")
		}
		if k < 0 {
			fmt.Println()
			break
		}
		fmt.Print("*\n")
		k -= 2
	}
}
