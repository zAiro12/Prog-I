package main

import (
	"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

/** in realtà serve solo a testare esistenza struct e nonZero()
 */
func TestZero(t *testing.T) {
	var timer Clock

	if timer.hour != 0 {
		t.Fail()
	}
}

func TestZeroDown(t *testing.T) {
	var timer Clock
	changeHour(&timer)
	fmt.Println(timer)

	if timer.hour != -1 {
		t.Fail()
	}
}

/*
func TestZeroPositive(t *testing.T) {
	var timer Clock
	//timer.min = 2
	changeHour(&timer)
	fmt.Println(timer)

	if nonZeroPositive(&timer) {
		t.Fail()
	}
}
*/

func TestNormal(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 2
	timer.sec = 20
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 3 && timer.min != 2 && timer.sec != 19 {
		t.Fail()
	}
}

func TestScattaMin(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 2
	timer.sec = 0
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 3 && timer.min != 1 && timer.sec != 59 {
		t.Fail()
	}
}

func TestScattaOra(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 0
	timer.sec = 0
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 2 && timer.min != 59 && timer.sec != 59 {
		t.Fail()
	}
}
