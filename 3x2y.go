package main 

import "fmt"

func main () {
	x := 0
	y := 0
	z := 0
fmt.Println("Ingrese los valores")
fmt.Scan(&x, &y, &z)
resultado := (3*x + 2*y) / 2 * z
fmt.Println (resultado)
}
