package ola_test

import (
	"golang-praticas/ola"
	"testing"
)

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t testing.TB, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("Diz olá para as pessoas", func(t *testing.T) {
		resultado := ola.Ola("Chris", "")
		esperado := "Olá, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := ola.Ola("", "")
		esperado := "Olá, Mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em espanhol", func(t *testing.T) {
		resultado := ola.Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Em frances", func(t *testing.T) {
		resultado := ola.Ola("Majur", "frances")
		esperado := "Bonjour, Majur"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
