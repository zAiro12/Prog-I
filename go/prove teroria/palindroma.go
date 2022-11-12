package main

func main() {

}

func paliIterASCII(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func paliIter(s string) bool {
	t := []rune(s)
	n := len(t)
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func paliRicASCII(s string) bool {
	if s == "" {
		return true
	}
	if s[0] != s[len(s)-1] { //uguale
		return false //uguale
	}
	return paliRicASCII(s[1 : len(s)-1]) //uguale
}

func paliRic(s string) bool {
	t := []rune(s)
	if len(t) <= 1 {
		return true
	}
	return t[0] == t[len(t)-1] && paliRic(string(t[1:len(t)-1])) //ugale a paliRicASCII 32:35
}
