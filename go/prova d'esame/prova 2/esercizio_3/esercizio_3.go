package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RegistroTelefonico map[string][]string

type Utenza struct {
	numero, cod string
}

func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()

	for _, utenza := range utenze {
		AggiungiUtenza(registro, utenza)
	}
	for telefono := range registro {
		if telefono[:3] == "338" {
			sim := Identifica(registro, telefono)
			fmt.Println("Il numero", telefono, "è associato alla sim", sim)
		}
	}
}
func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ";")
		utenze = append(utenze, Utenza{riga[0], riga[1]})
	}
	return
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	registro = make(RegistroTelefonico)
	return
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) RegistroTelefonico {
	registro[utenza.numero] = append(registro[utenza.numero], utenza.cod)
	return registro
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	lista := registro[telefono]
	lung := len(lista)

	if lung > 0 {
		codiceSIM = lista[lung-1]
	}
	return
}
