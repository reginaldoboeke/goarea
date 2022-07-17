package goarea

import "math"

// Valor de Pi
const Pi = 3.1416

//  Calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect eh responsavel por calcular a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Nao eh visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
