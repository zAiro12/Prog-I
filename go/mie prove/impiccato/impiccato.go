package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() //create a map for storing clear funcs
var errore int

func main() {
	var indovina string
	fmt.Print("Inserisci la parola da indovinare: ")
	fmt.Scan(&indovina)
	CallClear()
	criptatastring := cripta(indovina)
	criptata := []rune(criptatastring)

	fmt.Printf("La parola da indovinare è:%s (%d)\n", criptatastring, len(criptatastring)/2)
	daIndovinare := []rune(indovina)
	for {
		var lettera string
		fmt.Scan(&lettera)
		if len(lettera) < 2 {
			daIndovinare, criptata = check(lettera, daIndovinare, criptata)

		} else {
			if lettera == indovina {
				fmt.Printf("Bravo, la parola è: %s. Hai vinto!! \n", indovina)
				break
			} else {
				errore += 2
			}
		}
		fmt.Println(errore)
	}
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func cripta(s string) string {
	var out string
	for range s {
		out += " _"
	}
	return out
}

func check(lettera string, daIndovinare, criptata []rune) ([]rune, []rune) {
	for i := 0; i < len(daIndovinare); i++ {
		if lettera == string(daIndovinare[i]) {
			daIndovinare[i] = rune(lettera[0])
			criptata[i] = rune(lettera[0])
		} else {
			errore++
		}
	}
	return daIndovinare, criptata
}
