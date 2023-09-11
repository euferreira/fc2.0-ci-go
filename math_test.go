package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Error("Resultado da soma é inválido: Resultado esperado: 30, Resultado obtido: ", total)
	}
	t.Log("Resultado da soma é válido: Resultado esperado: 30, Resultado obtido: ", total)
}
