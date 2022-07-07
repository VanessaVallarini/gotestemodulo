//mesmo nome da pasta
package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre
//o perímetro de uma circunferência e seu diâmetro
const Pi = 3.1416

// Cir é responsável por calcular a área da circunferência
func Cir(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rec é responsável por calcular a área de um quadrado
func Rec(base, altura float64) float64 {
	return base * altura
}

//apenas para mostrar que por ser privada não vamos conseguir usar quando importarmos o pacote
func _TrianguloEq(base, altura float64) float64 {
	return (base + altura) / 2
}
