package main

import (
	"image"
	"image/png"
	"math"
	"net/http"
	"os"

	"github.com/holizz/terrapin"
)

func main() {
	http.HandleFunc("/pippo", pagePippo)
	http.HandleFunc("/pluto", gg)
	http.HandleFunc("/forme", kochImage)
	http.ListenAndServe(":3000", nil)
}

func pagePippo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<!doctype html>
		<title>Pagina pippo </title>
		<h1 style="color: red">Pippo</h1>
		<h3>ciao</h3>
		<p><img scr="/pluto"></p>
	`))
}
func gg(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("/Users/zairo/Dropbox/UNIMI/programmazione/primo anno/prove teroria/server/palkia.png")
	if err != nil {
		return
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	if err != nil {
		panic("errore")
	}
	png.Encode(w, image)
}
func kochImage(w http.ResponseWriter, r *http.Request) {
	i := image.NewRGBA(image.Rect(0, 0, 3000, 3000))
	t := terrapin.NewTerrapin(i, terrapin.Position{500.0, 2000.0})

	/* for i := 0; i < 8; i++ {
		t.Forward(50)
		t.Right(math.Pi - 135*2*math.Pi/360)
	} */
	kochCurve(t, 20.0, 7)
	png.Encode(w, i)
}
func kochCurve(t *terrapin.Terrapin, lung float64, liv int) {
	if liv == 0 {
		t.Forward(lung)
		return
	}
	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3)
	kochCurve(t, lung, liv-1)
	t.Right(math.Pi - math.Pi/3)
	kochCurve(t, lung, liv-1)
	t.Left(math.Pi / 3)
	kochCurve(t, lung, liv-1)
}
