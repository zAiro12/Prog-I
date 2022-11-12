package main

import "fmt"

func main(){
    var n int
	fmt.Print("inserisci anno: ")
	fmt.Scan(&n)
	if n%4==0{
		fmt.Println("bisestile")
	}else{
		fmt.Println("non bisestile")
	}
}
