package main

import (
	"fmt"
)

// Definición de la clase base "Figura"
type Figura interface {
	Area() float64
}

// Definición de la clase derivada "Rectángulo"
type Rectangulo struct {
	Base   float64
	Altura float64
}

// Implementación del método Area para Rectángulo
func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

// Definición de la clase derivada "Triángulo"
type Triangulo struct {
	Base   float64
	Altura float64
}

// Implementación del método Area para Triángulo
func (t Triangulo) Area() float64 {
	return 0.5 * t.Base * t.Altura
}

func calcularArea(figura Figura) {
	area := figura.Area()
	fmt.Printf("El área es: %f\n", area)
}

func main() {
	rectangulo := Rectangulo{Base: 10, Altura: 5}
	triangulo := Triangulo{Base: 7, Altura: 3}

	calcularArea(rectangulo)
	calcularArea(triangulo)
}
