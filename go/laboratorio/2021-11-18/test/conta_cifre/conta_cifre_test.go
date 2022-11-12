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

func TestContaCifreBoolean(t *testing.T) {
	var numCifre [10]int

	st := "ciao65656565"
	if !contaCifre(st, &numCifre) {
		t.Fail()
	}
}

func TestContaCifre(t *testing.T) {
	var numCifre [10]int

	st := "ciao65656565alkdjaslj lkj lakjl dskja98989898"
	if !contaCifre(st, &numCifre) {
		t.Fail()
	}

	fmt.Println(numCifre)

	if numCifre[0] != 0 || numCifre[5] != 4 {
		t.Fail()
	}
}
