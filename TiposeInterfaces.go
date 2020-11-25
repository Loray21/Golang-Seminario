package tipoEinterfaces

import "fmt"

//definir interface

//go no tiene manera de saber si se creo un interface, en este caso si person implementa todos los metods
//de Printable queire decir que la implementa
type Printable interface {
	print()
}

//defino estructura
type person struct {
	name string
}

type figure struct {
	name string
}

func (p person) print() {
	fmt.Println("[person]", p.name)
}

func (f figure) print() {
	fmt.Println("[figure]", f.name)
}

/*//funcion remove el asterisco es un puntero que apunta al objeto y modifica el estado de est
func (p *person) clean() {
	p.name = ""

}
*/
func invokePrint(p Printable) {
	p.print()
}
func main() {
	//instancia de person
	p := person{name: "juan"}

	invokePrint(p)

}
