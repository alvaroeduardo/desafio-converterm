package main

import "fmt"

func main() {
	fmt.Println("Criando um Programa em Go para a Conversão de Escalas Termométricas")

	ebulicaoKelvin := 373.0
	temperaturaCelsius := ebulicaoKelvin - 273.0

	fmt.Printf("A temperatura de embulição da água em K = %g, convertendo este valor para celsius é em C = %g", ebulicaoKelvin, temperaturaCelsius)
}
