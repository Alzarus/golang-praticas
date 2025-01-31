package iteracao

func Repetir(caractere string, qtdRepeticoes int) string {
	var repeticoes string
	for i := 0; i < qtdRepeticoes; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
