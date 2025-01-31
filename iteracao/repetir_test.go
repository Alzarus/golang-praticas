package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado: '%s', resultado: '%s'", esperado, repeticoes)
	}
}

var result string

func BenchmarkRepetir(b *testing.B) {
	var r string

	for i := 0; i < b.N; i++ {
		r = Repetir("a", 5)
	}

	result = r
}

func ExampleRepetir() {
	repeticoes := Repetir("p", 5)
	fmt.Println(repeticoes)
	//Output: ppppp
}
