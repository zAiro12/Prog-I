package main

import (
	"fmt"
)

func main() {
	var n int
	var l string
	fmt.Print("Scegli una lingua: ita, eng, esp, deu, fra: ")
	fmt.Scan(&l)

	switch l {
	case "ita":
		fmt.Print("Scegli un numero: ")
		fmt.Scan(&n)
		fmt.Println(ita(n))

	case "eng":
		fmt.Print("Choose a number: ")
		fmt.Scan(&n)
		fmt.Println(eng(n))

	case "esp":
		fmt.Print("Elige un número: ")
		fmt.Scan(&n)
		fmt.Println(esp(n))

	case "deu":
		fmt.Print("Wähle eine Nummer: ")
		fmt.Scan(&n)
		fmt.Println(deu(n))

	case "fra":
		fmt.Print("Choisir un numéro: ")
		fmt.Scan(&n)
		fmt.Println(fra(n))

	default:
		fmt.Println("Non hai inserito una lingua valida")
	}

}

func ita(n int) string {
	if n >= 0 && n <= 19 {
		switch n {
		case 10:
			return "dieci"
		case 11:
			return "undici"
		case 12:
			return "dodici"
		case 13:
			return "tredici"
		case 14:
			return "quattordici"
		case 15:
			return "quindici"
		case 16:
			return "sedici"
		case 17:
			return "diciasette"
		case 18:
			return "diciotto"
		case 19:
			return "diciannove"
		}

	}
	var res string
	var fin string = "a"
	switch n / 10 {
	case 2:
		res = "vent"
		fin = "i"
	case 3:
		res = "trent"
	case 4:
		res = "quarant"
	case 5:
		res = "cinquant"
	case 6:
		res = "sesant"
	case 7:
		res = "settant"
	case 8:
		res = "ottant"
	case 9:
		res = "novant"

	}
	if n%10 != 8 && n%10 != 1 {
		res += fin
	}

	switch n % 10 {
	case 1:
		res += "uno"
	case 2:
		res += "due"
	case 3:
		res += "tre"
	case 4:
		res += "quettro"
	case 5:
		res += "cinque"
	case 6:
		res += "sei"
	case 7:
		res += "sette"
	case 8:
		res += "otto"
	case 9:
		res += "nove"
	}
	return res
}

func eng(n int) string {
	if n >= 0 && n <= 12 {
		switch n {
		case 10:
			return "ten"
		case 11:
			return "eleven"
		case 12:
			return "twelwe"
		}

	}
	var res string
	switch n / 10 {
	case 1:
		res = "teen"
	case 2:
		res = "twenty"
	case 3:
		res = "thirty"
	case 4:
		res = "fourty"
	case 5:
		res = "fifty"
	case 6:
		res = "sixty"
	case 7:
		res = "seventy"
	case 8:
		res = "eithty"
	case 9:
		res = "ninety"

	}
	if n > 20 {
		switch n % 10 {
		case 1:
			res += "one"
		case 2:
			res += "two"
		case 3:
			res += "three"
		case 4:
			res += "four"
		case 5:
			res += "five"
		case 6:
			res += "six"
		case 7:
			res += "seven"
		case 8:
			res += "eigth"
		case 9:
			res += "nine"
		}
		return res
	} else {
		switch n % 10 {
		case 3:
			res = "thir" + res
		case 4:
			res = "four" + res
		case 5:
			res = "five" + res
		case 6:
			res = "six" + res
		case 7:
			res = "seven" + res
		case 8:
			res = "eigth" + res
		case 9:
			res = "nine" + res
		}
		return res
	}

}

func esp(n int) string {
	if n >= 0 && n <= 19 {
		switch n {
		case 10:
			return "diez"
		case 11:
			return "once"
		case 12:
			return "donce"
		case 13:
			return "trece"
		case 14:
			return "catorce"
		case 15:
			return "quince"
		case 16:
			return "dieciséis"
		case 17:
			return "diecisiete"
		case 18:
			return "dieciocho"
		case 19:
			return "diecinueve"
		}

	}
	var res string
	if n == 20 {
		res = "veinte"
	} else {
		switch n / 10 {
		case 2:
			res = "veinti"
		case 3:
			res = "trenta"
		case 4:
			res = "quaranta"
		case 5:
			res = "cinquanta"
		case 6:
			res = "sesanta"
		case 7:
			res = "settanta"
		case 8:
			res = "ottanta"
		case 9:
			res = "novanta"
		}
	}
	if n <= 29 {
		switch n % 10 {
		case 1:
			res += "uno"
		case 2:
			res += "dós"
		case 3:
			res += "trés"
		case 4:
			res += "cuatro"
		case 5:
			res += "cinco"
		case 6:
			res += "séis"
		case 7:
			res += "siete"
		case 8:
			res += "ocho"
		case 9:
			res += "nueve"
		}
		return res
	} else {
		switch n % 10 {
		case 1:
			res += " y uno"
		case 2:
			res += " y dós"
		case 3:
			res += " y trés"
		case 4:
			res += " y cuatro"
		case 5:
			res += " y cinco"
		case 6:
			res += " y séis"
		case 7:
			res += " y siete"
		case 8:
			res += " y ocho"
		case 9:
			res += " y nueve"
		}
		return res
	}
}

func deu(n int) string {
	if n >= 0 && n <= 12 {
		switch n {
		case 10:
			return "zehn"
		case 11:
			return "elf"
		case 12:
			return "zwölf"
		}

	}
	var res string
	switch n / 10 {
	case 1:
		res = "zehn"
	case 2:
		res = "zwanzig"
	case 3:
		res = "dreißig"
	case 4:
		res = "vierzig"
	case 5:
		res = "fünfzig"
	case 6:
		res = "sechzig"
	case 7:
		res = "siebzig"
	case 8:
		res = "achtzig"
	case 9:
		res = "neunzig"

	}

	switch n % 10 {
	case 1:
		res = "eins" + res
	case 2:
		res = "zwei" + res
	case 3:
		res = "drei" + res
	case 4:
		res = "vier" + res
	case 5:
		res = "fünf" + res
	case 6:
		res = "sechs" + res
	case 7:
		res = "sieben" + res
	case 8:
		res = "acht" + res
	case 9:
		res = "neun" + res
	}
	return res

}

func fra(n int) string {
	if n >= 0 && n <= 19 {
		switch n {
		case 10:
			return "dix"
		case 11:
			return "onze"
		case 12:
			return "douze"
		case 13:
			return "treize"
		case 14:
			return "quatorze"
		case 15:
			return "quinze"
		case 16:
			return "seize"
		case 17:
			return "dix-sept"
		case 18:
			return "dix-huit"
		case 19:
			return "dix-neuf"
		}
	}
	var res string
	switch n / 10 {
	case 2:
		res = "vingt"
	case 3:
		res = "trente"
	case 4:
		res = "quarante"
	case 5:
		res = "cinquante"
	case 6:
		res = "soixante"
	case 7:
		res = "soixante-dix"
	case 8:
		res = "quatre-vingts"
	case 9:
		res = "quatre-vingt-dix"
	}

	if n <= 0 {
		switch n % 10 {
		case 1:
			res = "un" + res
		case 2:
			res = "deux" + res
		case 3:
			res = "trois" + res
		case 4:
			res = "quatre" + res
		case 5:
			res = "cinq" + res
		case 6:
			res = "six" + res
		case 7:
			res = "sept" + res
		case 8:
			res = "huit" + res
		case 9:
			res = "neuf" + res
		}
		return res
	} else {
		switch n % 10 {
		case 1:
			res += " et un"
		case 2:
			res += "-deux"
		case 3:
			res += "-trois"
		case 4:
			res += "-quatre"
		case 5:
			res += "-cinq"
		case 6:
			res += "-six"
		case 7:
			res += "-sept"
		case 8:
			res += "-huit"
		case 9:
			res += "-neuf"
		}
		return res
	}
}
