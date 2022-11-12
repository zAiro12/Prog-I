package main

import "fmt"

func main() {
	var consu, resa, dist, carb float64

	fmt.Print("distanza percorsa (in km): ")
	fmt.Scan(&dist)
	fmt.Print("quantità di carburante utilizzata (in l): ")
	fmt.Scan(&carb)

	consu = carb / dist
	resa = dist / carb

	fmt.Println("consumo medio:", consu, "l/km")
	fmt.Println("resa media:", resa, "km/l")

}
