package estDeDatos

import "fmt"

func main() {

	//definimo arreglo
	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	//recorrer arreglo

	for i, v := range arr {
		fmt.Println(i, v)
	}

	//slices son iguakl que los arraylist

	//var l []int
	//para dar un valor fijio, si llega a 100 lo expande otra vez + 100
	l := make([]int, 100)
	//agrego elementos
	l = append(l, 10)
	l = append(l, 1)
	l = append(l, 25)

	for i, l := range l {
		fmt.Println(i, l)
	}

	//mapas igual que el hashmap de java

	m := make(map[int]string)
	m[0] = "a"
	m[1] = ""

	//recorrer mapa

	for k, v := range m {
		fmt.Println(k, v)
	}

}
