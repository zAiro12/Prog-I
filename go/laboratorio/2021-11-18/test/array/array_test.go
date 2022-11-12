package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

//var array = [5]int{10, 20, 30, 40, 50} // domanda da fare a loro: perché se lo metto globale i test falliscono?

func TestTestaCoda(t *testing.T) {
	var array = [5]int{10, 20, 30, 40, 50}
	switchFirstLast(&array)
	if array[0] != 50 || array[4] != 10 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	var array = [5]int{10, 20, 30, 40, 50}
	reverse(&array)
	if array[0] != 50 || array[1] != 40 || array[2] != 30 || array[3] != 20 || array[4] != 10 {
		t.Fail()
	}
}
