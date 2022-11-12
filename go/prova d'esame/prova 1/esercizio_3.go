package main

import (
	"bufio"
	"os"
	"strings"
)

type Studente struct {
	nome, cognome string
}

func main() {
	var studenti []Studente
	//capienza := os.Args[1]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		l := scanner.Text()
		studenti = append(studenti, creaStudente(l))
	}
}

func creaStudente(id string) Studente {
	s := strings.Split(id, ".")
	var in Studente
	in.nome = s[0]
	cognome := s[1]
	var i int
	for i := 0; i < len(cognome); i++ {
		if cognome[i] >= '0' && cognome[i] <= '9' {
			break
		}
	}
	in.cognome = cognome[:i]
	return in
}

func parseStudenti(s string){

}