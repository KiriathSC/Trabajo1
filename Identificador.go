package main 

import "fmt"

func main(){
	manzana := 85
	uva := &manzana
	coco := &uva
	
	if *uva == manzana{
	    fmt.Println("la direccion de memoria es manzana")
    }
    if **coco == *uva{
	    fmt.Println("la direccion de memoria es la misma a la anterior")
    } 
}
