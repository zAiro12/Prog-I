package main

import "fmt"

func main() {
	var sec int

	fmt.Print("numero di secondi: ")
	fmt.Scan(&sec)

	fmt.Print("g:h:m:s = ", sec/(60*60*24), ":", sec%(60*60*24)/(60*60), ":", sec%(60*60*24)%(60*60)/60, ":", sec%(60*60*24)%(60*60)%(60), "\n")
}
