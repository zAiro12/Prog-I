package main

import (
	"fmt"
	"os"
	"os/exec"
)

var table [9][9]int

func main() {
	var x, y, n int
	stampa()
	fmt.Print("Inserisci numero e posizione in questo formato [N(numero) x(riga) y(colonna)]. Se si vuole stoppare il programma inserire 0: ")
	for {
		fmt.Scanf("%d %d %d", &n, &x, &y)
		if n == 0 {
			break
		}
		if n < 1 || n > 9 {
			fmt.Println("input non valido! Numero inserito sbagliato")
			continue
		}
		if x < 1 || x > 9 || y < 1 || y > 9 {
			fmt.Println("input non valido! Coordinata sbagliata")
			continue
		}
		table[x-1][y-1] = n
		cmd := exec.Command("clear") // Mac
		cmd.Stdout = os.Stdout
		cmd.Run()
		stampa()
		fmt.Print("Input: ")
	}
	risolvi()
}

func stampa() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				fmt.Print("  ")
			} else {
				fmt.Print(table[i][j], " ")
			}
			if j == 2 || j == 5 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if i == 2 || i == 5 {
			fmt.Println("----------------------")
		}
	}
}

func risolvi() {
	for i := 0; i < 9; i++ {
		for j := 0; j < j; j++ {
			candidato(i, j)
		}
	}
}

func candidato(x, y int) {

}

func controlla(x,y int) bool {
	for i := 0; i < 9; i++ {
		if table[x][i] == table[x][y] && table[i][y] == table[x][y] {
			return true
		}	
	}
	return false
}
