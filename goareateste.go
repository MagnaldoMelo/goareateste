package goareateste

//Instalar na pasta GOPATH. Se certificar que está emgithub.com/MagnaldoMelo/golang_cod3r
//Esse é o padrão para o tipo de reuso abordado
//Com isso ele importa automaticamente, igual o fmt
import "math"

//Pi constante PI
const Pi = 3.1416

//Circ - função publica
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect - calcular área
func Rect(base, altura float64) float64 {
	return base * altura
}

//função que não é visível (_)
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
