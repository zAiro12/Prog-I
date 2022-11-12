package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserire una frase:")
	for scanner.Scan() {
		// premere "CTRL+D" per interrompere la Scan
		s = scanner.Text()
		fmt.Print("Inserire il numero dello shift: ")
		fmt.Scan(&n)
		s += s[:n] // Copio le prime n lettere in fondo alla stringa
		//shift di n posizione a sinistra
		for i := n; i < len(s); i++ {
			s[i-n] = s[i]
		}
		s = s[:len(s)-n] // tolgo le posizioni di troppo
		fmt.Println(s)
	}
}
