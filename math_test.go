package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 2)

	if total != 12 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 12)
	}
}
