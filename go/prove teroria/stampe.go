package main

import "fmt"

func main() {
	var n int
	//var ast string
	fmt.Scan(&n)
	//fatto meglio ma simo dice che non va bene
	/*for i := 1; i <= n; i++ {
		ast = ast + "*"
		fmt.Println(ast)
	}*/
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
