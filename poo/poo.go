package main

import "fmt"


type  Animal struct {
	cantidad_patas int
	color_pelo string
	sonido string
}

//NewAnimal Contructor
func NewAnimal(cantidad_patas int, color_pelo string, sonido string) *Animal {
	return &Animal{cantidad_patas: cantidad_patas, color_pelo: color_pelo, sonido: sonido}
}

type Gato struct {
	//HERENCIA
	Animal
	vidas int
}

//NewGato Contructor
func NewGato(cantidad_patas int, color_pelo string, sonido string, vidas int) *Gato {
	return &Gato{Animal: *NewAnimal(cantidad_patas, color_pelo, sonido), vidas: vidas}
}

// Implementacion de la interfaz ruido en la clase gato
func (this *Gato) hacerRuido() string {
	return this.sonido
}

type Perro struct {
	//HERENCIA
	Animal
	oficio string
}

// NewPerro  Contructor/**
func NewPerro(cantidad_patas int, color_pelo string, sonido string, oficio string) *Perro {
	return &Perro{Animal: *NewAnimal(cantidad_patas, color_pelo, sonido), oficio: oficio}
}
// Implementacion de la interfaz ruido en la clase perro
func (this *Perro) hacerRuido() string {
	return this.sonido
}
// Interfaz ruido
type Ruido interface {
	hacerRuido() string
}

func sonidoAnimal(r Ruido) {
	fmt.Println(r.hacerRuido())
}


func main() {
 fmt.Println("Hola")
 nuevoGato := NewGato(3, "negro", "miauuuu", 7)
 nuevoPerro := NewPerro(4, "dorado", "guauuuuuu", "lazarillo")
 fmt.Println(*nuevoPerro, *nuevoGato)
 sonidoAnimal(nuevoPerro)
 sonidoAnimal(nuevoGato)
 
}
