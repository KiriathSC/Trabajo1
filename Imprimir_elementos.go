package main

import "fmt"

func imprimir(board [5][5]string){

	fmt.Println(board)
}

func main(){
	var board[5][5]string
	board[0][0]= "r"
	board[0][0]= "r"
	for column:= range board[1]{
		board[4][column] = "p"
	}
	imprimir(board)
}
