package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSoma(t *testing.T) {
	total := Soma(10, 2)

	if total != 12 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 12)
	}
}

func TestMain(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != "12\n" {
		t.Errorf("Expected %s, got %s", "this is value: 12", out)
	}
}
