package main

import "fmt"

func diaSemana(numero int) string {
	switch numero{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func idade (numero int) string {
	switch {
	case numero >= 18:
		return "Maior de idade"
	case numero < 18:
		return "Menor de idade"
	case numero >= 60:
		return "Idoso"
	default:
		return "Erro"
	}
}

func main() {
	dia := diaSemana(3)
	fmt.Println(dia)

	age := idade(10)
	fmt.Println(age)
}