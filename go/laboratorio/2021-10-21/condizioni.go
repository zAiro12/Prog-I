package main

import "fmt"

func main() {
	var n int
	//fmt.Scan(&n)

	fmt.Print("int unguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n == 10)

	fmt.Print("int diverso da 10: ")
	fmt.Scan(&n)
	fmt.Println(n != 10)

	fmt.Print("int diverso da 10 e da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 && n != 20)

	fmt.Print("int diverso da 10 o da 20: ")
	fmt.Scan(&n)
	fmt.Println(n != 10 || n != 20)

	fmt.Print("maggiore o uguale a 10: ")
	fmt.Scan(&n)
	fmt.Println(n >= 10)

	fmt.Print("compreso tra 10 e da 20, nell'intervallo [10,20]: ")
	fmt.Scan(&n)
	fmt.Println(n >= 10 && n <= 20)

	fmt.Print("int compreso tra 10 e da 20, nell'intervallo (10,20): ")
	fmt.Scan(&n)
	fmt.Println(n > 10 && n < 20)

	fmt.Print("int compreso tra 10 e da 20, nell'intervallo [10,20): ")
	fmt.Scan(&n)
	fmt.Println(n >= 10 && n < 20)

	fmt.Print("esterno all'intervallo [10,20]: ")
	fmt.Scan(&n)
	fmt.Println(n <= 10 || n >= 20)

	fmt.Print("tra 10 e 20 (nell'intervallo [10,20]) o tra 30 e 40 ([30,40]): ")
	fmt.Scan(&n)
	fmt.Println((n >= 10 && n <= 20) || (n >= 30 && n <= 40))

	fmt.Print("multiplo di 4 ma non di 100: ")
	fmt.Scan(&n)
	fmt.Println(((n % 4) == 0) && ((n % 100) != 0))

	fmt.Print("dispari e compreso tra 0 e 100 ([0,100]): ")
	fmt.Scan(&n)
	fmt.Println(((n % 2) != 0) && (n >= 0 && n <= 100))

	var x float64
	const k = 1e-6
	fmt.Print("float64 'vicino' a 10.0 (non più lontano di 10^-6) (1e-6): ")
	fmt.Scan(&x)
	fmt.Println((x > 10.0-k) && (x < 10.0+k))

}
