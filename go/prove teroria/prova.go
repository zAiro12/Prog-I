/*
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(randomString(10))
}

func randomString(l int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < l; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	line, err := in.ReadString('\n')
	fmt.Print(line)
	if err != nil {
		fmt.Println(err)
	}
}
*/

/* /* 0 0 1 1 2 2 3 4 4 5 5 6 7 8 8 9

package main

import "fmt"

func main() {
	// var cmax, a, b int
	var r, r1 rune
	var s string
	for {
		fmt.Scan(&r)
		// Controllo che il nuovo numero non sia minore del precedente o troppo maggiore
		if (r < r1) || (r > r1+1) {
			break
		}
		s += string(r)
		r1 = r
	}
	fmt.Println(s)
}
*/

/* package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("ciao, spero che io venga canellato")
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("chi sa ha funzionato...")
}
*/

package main

import "fmt"

func main() {
	var v []float64
	const N = 6
	const L = 0.01
	for i:=0; i<N; i++{
		var in float64
		fmt.Scan(&in)
		v = append(v, in)
	}

	for i:=0; i<len(v); i++{
		if v[i]>L{
			fmt.Print(v[i], " ")
		}
	}
	fmt.Print()
}
