package main

import "fmt"

func main(){
	var x, y float64
	fmt.Print("x,y: ")
	fmt.Scan(&x, &y)
	if x>=0 && y>=0{
		fmt.Println("I quadrante")
	} else if x<=0 && y>=0{
		fmt.Println("II quadrante")
	} else if x<=0 && y<=0{
		fmt.Println("III quadrante")
	} else if x>=0 && y<=0{
		fmt.Println("IV quadrante")
	}

}