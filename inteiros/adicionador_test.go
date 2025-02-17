package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	resultado := Adiciona(2, 2)
	esperado := 4

	if resultado != esperado {
		t.Errorf("esperado: '%d', resultado: '%d'", esperado, resultado)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
