package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v"

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

//go test -cover para ver quanto do teste está sendo coberto
// go test -coverprofile=resultado.out &&&& go tool cover -func=resultado.out
// go tool cover -html=resultado.out GERAR UM HTML
