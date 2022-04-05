package main

import( 
	"fmt"
	"strconv"
	"bufio"
	"os"
)
func main(){
	var nombre string
	nombre = "mike"
	var num string
	num = "50"
	var numero int
	numero = 10
	xz := 15
	var bandera bool
	bandera = true
	
	x_str := strconv.Itoa(xz)
	num_int,_ := strconv.Atoi(num)
	 
	fmt.Println("Hola Mundo", num_int, nombre, numero, xz, x_str, bandera)
	
	//codigo para poner datos por teclado normal
	var edad int
	fmt.Println("Coloca tu edad:")
	fmt.Scanf("%d\n",&edad)
	fmt.Println("Tu edad es:", edad)
	
	//codigo para poner datos por teclado con bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre:")
	nombre,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("hola",nombre)
	}
	
	x := 10
	y := 5
	
	if	x > 5{
		fmt.Println(x, "es mayor", y)
	}else{
		fmt.Println("error")
	}	
}
