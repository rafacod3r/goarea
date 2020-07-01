package goarea

import "math"

// PI É UMA proporção numerica definida pela relaçao entre
// o perimetro de uma circunferencia e seu diamentro
const PI = 3.1416

//Circ é responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Rect é responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é uma func visivel ! se iniciar por _ ou letra minuscula
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
