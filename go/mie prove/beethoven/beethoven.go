package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Inserisci il nome di una sinfonia di beethoven e te la suonerò come farebbe lui: ")
	for scanner.Scan() {
		s += scanner.Text()
	}
	fmt.Println("AHAHAH, non te la posso suonare, beethoven è morto!")
}
