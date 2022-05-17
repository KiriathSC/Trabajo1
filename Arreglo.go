package main 

import "fmt"

func copia(semana[7]string){
	//funcion que permite pasar un areglo 

	for i,_:= range semana{
		semana[i]= "new" + semana[i]
	}
}
func main(){
	semana:= [...]string{
		"lunes"
		"martes"
		"miercoles"
		"jueves"
		"viernes"
		"sabado"
		"domingo"
	}
	copia(semana)
	fmt.Println(semana)
}
