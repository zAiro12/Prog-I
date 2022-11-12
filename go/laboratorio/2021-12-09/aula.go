package main

import (
	"fmt"
)

func main() {
	aula, _ := creaAula(7, 6)
	if occupa(aula, 4, 3) {
		fmt.Println("fatto")
	}
	if libera(aula, 4, 3) {
		fmt.Println("fatto")
	}
	if verificaDistanziamento(aula, 4, 3) {
		fmt.Println("posto libero")
	} else {
		fmt.Println("posto occupato")
	}
	stampaAula(aula)
}

func creaAula(col, riga int) ([][]bool, bool) {
	var aula [][]bool
	flag := false
	if riga >= 1 && col >= 1 {
		flag = true
	}
	if flag {
		for i := 0; i <= col; i++ {
			var rigaora []bool
			for j := 0; j <= riga; j++ {
				rigaora = append(rigaora, false)
			}
			aula = append(aula, rigaora)
		}
		return aula, flag
	} else {
		col = 5
		riga = 7
		for i := 0; i <= col; i++ {
			var rigaora []bool
			for j := 0; j <= riga; j++ {
				rigaora = append(rigaora, false)
			}
			aula = append(aula, rigaora)
		}
		return aula, flag

	}
}

func stampaAula(a [][]bool) {
	for col := 0; col < len(a); col++ {
		for riga := 0; riga < len(a[col]); riga++ {
			if a[col][riga] == true {
				fmt.Print("x ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
}

func occupa(a [][]bool, col, riga int) bool {
	if col > len(a) || riga > len(a[0]) || col < 0 || riga < 0 || a[col][riga] == true {
		return false
	} else {
		a[col][riga] = true
		return true
	}

}

func libera(a [][]bool, col, riga int) bool {
	if col > len(a) || riga > len(a[0]) || col < 0 || riga < 0 || a[col][riga] == false {
		return false
	} else {
		a[col][riga] = false
		return true
	}
}

func verificaDistanziamento(a [][]bool, riga, col int) bool {
	switch col {
	case 0: //CASO A B C
		switch riga {
		//A
		case 0:
			if a[col+1][riga] == false && a[col][riga+1] == false {
				fmt.Println("A")
				return true
			}

		//C
		case len(a[col]) - 1:
			if a[col+1][riga] == false && a[col][riga-1] == false {
				fmt.Println("C")
				return true
			}

		//B
		default:
			if a[col+1][riga] == false && a[col][riga-1] == false && a[col][riga+1] == false {
				fmt.Println("B")
				return true
			}
		}
	case len(a) - 1: //caso G H L
		switch riga {
		//G
		case 0:
			if a[col-1][riga] == false && a[col][riga+1] == false {
				fmt.Println("G")
				return true
			}

		//L
		case len(a[col]) - 1:
			if a[col-1][riga] == false && a[col][riga-1] == false {
				fmt.Println("L")
				return true
			}

		//H
		default:
			if a[col-1][riga] == false && a[col][riga-1] == false && a[col][riga+1] {
				fmt.Println("H")
				return true
			}
		}

	default: //caso D E F
		switch riga {
		//D
		case 0:
			if a[col-1][riga] == false && a[col][riga+1] && a[col+1][riga] {
				fmt.Println("D")
				return true
			}

		//F
		case len(a[col]) - 1:
			if a[col-1][riga] == false && a[col][riga-1] == false && a[col+1][riga] {
				fmt.Println("F")
				return true
			}

		//E
		default:
			if a[col-1][riga] == false && a[col][riga-1] == false && a[col][riga+1] && a[col+1][riga] {
				fmt.Println("E")
				return true
			}
		}
	}
	fmt.Println("fine")
	return false
}
