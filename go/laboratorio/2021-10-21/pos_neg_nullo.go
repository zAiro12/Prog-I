package main

import "fmt"

func main(){
	var x int
	fmt.Print("un int: ")
	fmt.Scan(&x)
	if x>0{
		fmt.Println("positivo")
	}else if x<0{
		fmt.Println("negativo")
	}else{
		fmt.Println("nullo")	
	}
}
