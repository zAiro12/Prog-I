/* package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func main() {
	var cantina []BottigliaVino
	file := os.Args[1]
	o, _ := os.Open(file)
	defer o.Close()
	for {
		var s string
		_, err := fmt.Fscan(o, &s)
		if err != nil {
			break
		}
		split := strings.Split(s, ",")
		anno, _ := strconv.Atoi(split[1])
		grado, _ := strconv.ParseFloat(split[2], 32)
		quantita, _ := strconv.Atoi(split[3])
		bot, ok := CreaBottiglia(split[0], anno, float32(grado), quantita)
		if ok {
			AggiungiBottiglia(bot, &cantina)
		}
	}
	for _, b := range cantina {
		fmt.Println(b.String())
	}

}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool){

}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	var nuova BottigliaVino
	if anno > 0 && grado > 0 && quantita > 0 && nome != "" {
		nuova.nome = nome
		nuova.anno = anno
		nuova.grado = grado
		nuova.quantita = quantita
		return nuova, true
	} else {
		return nuova, false
	}
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	split := strings.Split(bot, ",")
	anno, _ := strconv.Atoi(split[1])
	grado, _ := strconv.ParseFloat(split[2], 32)
	quantita, _ := strconv.Atoi(split[3])
	var nuova BottigliaVino
	if anno > 0 && grado > 0 && quantita > 0 && split[0] != "" {
		nuova.nome = split[0]
		nuova.anno = anno
		nuova.grado = float32(grado)
		nuova.quantita = quantita
		*cantina = append(*cantina, nuova)
	} else {
		*cantina = append(*cantina, nuova)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	c := 0
	for i := 0; i < len(cantina); i++ {
		if cantina[i].nome == nome {
			c++
		}
	}
	return c
}

func (bot BottigliaVino) String() string {
	out := fmt.Sprintf("%s,%d,%.1f,%d", bot.nome, bot.anno, bot.grado, bot.quantita)
	/* pos := strings.Index(out, ".0,")
	if pos != -1 {
		out = out[:pos] + "," + out[pos+3:]
	}
	return out
}



package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

type  BottigliaVino struct {
	nome string
	anno int
	grado float32
	quantita int
}

func (b BottigliaVino) String() string {
	out := fmt.Sprintf("%s,%d,%.1f,%d", b.nome, b.anno, b.grado, b.quantita)
	pos := strings.Index(out, ".0,")

	if pos != -1 {
		out = out[:pos] + "," + out[pos+3:]
	}

	return out
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || grado <= 0 || quantita <= 0 {
		return BottigliaVino{}, false
		} else {
			return BottigliaVino{nome, anno, grado, quantita}, true
		}
	}

	func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
		line := strings.Split(riga, ",")
		if len(line) != 4 {
			return BottigliaVino{},false
		}
		anno, _ := strconv.Atoi(line[1])
		grd, _ := strconv.ParseFloat(line[2], 32)
		grado := float32(grd)
		qnt, _ := strconv.Atoi(line[3])
		return CreaBottiglia(line[0], anno, grado, qnt)
	}

	func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
		*cantina = append(*cantina, bot)
	}

	func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
		bottiglia, ok := CreaBottigliaDaRiga(bot)
		if ok {
			AggiungiBottiglia(bottiglia, cantina)
		}
	}

	func ContaPerTipo(nome string, cantina []BottigliaVino) int {
		count := 0
		for _, tipo := range cantina {
			if tipo.nome == nome {
				count++
			}
		}

		return count
	}

	func main() {
		cantina := []BottigliaVino{}

		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("file non trovato")
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			if scanner.Text() != "" {
				AggiungiBottigliaDaRiga(scanner.Text(), &cantina)
			}
		}

		for _, b := range cantina {
			fmt.Println(b.String())
		}
	}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func main() {
	var cantina []BottigliaVino
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() != "" {
			AggiungiBottigliaDaRiga(scanner.Text(), &cantina)
		}
	}
	for _, b := range cantina {
		fmt.Println(b.String())
	}
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || grado <= 0 || quantita <= 0 {
		return BottigliaVino{}, false
	} else {
		return BottigliaVino{nome, anno, grado, quantita}, true
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	if len(riga) == 4 {
		split := strings.Split(riga, ",")
		anno, _ := strconv.Atoi(split[1])
		grado, _ := strconv.ParseFloat(split[2], 32)
		quantita, _ := strconv.Atoi(split[3])
		return CreaBottiglia(split[0], anno, float32(grado), quantita)
	} else {
		return BottigliaVino{}, false
	}
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	aggiunti, err := CreaBottigliaDaRiga(bot)
	if err {
		*cantina = append(*cantina, aggiunti)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	c := 0
	for i := 0; i < len(cantina); i++ {
		if cantina[i].nome == nome {
			c++
		}
	}
	return c
}

func (bot BottigliaVino) String() string {
	out := fmt.Sprintf("%s,%d,%.1f,%d", bot.nome, bot.anno, bot.grado, bot.quantita)
	pos := strings.Index(out, ".0,")
	if pos != -1 {
		out = out[:pos] + "," + out[pos+3:]
	}
	return out
}
