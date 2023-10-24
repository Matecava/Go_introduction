// “En Easyred una fábrica de computadoras se planea ofrecer a los clientes un descuento que dependerá del número de computadoras que compren. Si las computadoras son menos de cinco se les dará un 10% de descuento sobre el total de la compra; si el número de computadoras es mayor o igual a cinco pero menos de diez se le otorga un 20% de descuento; y si son 10 o más se les da un 40% de descuento. El precio de cada computadora es de $1.100.000”

package main

import "fmt"

func main() {
	var n int
	computadora := 1100000
	fmt.Print("Computadoras compradas: ")
	fmt.Scanln(&n)
	if n < 5 {
		precio_inicial := n * computadora
		descuento := precio_inicial / 10
		precio_final := precio_inicial - descuento
		fmt.Println("Tienes un descuento del 10%, total de la compra: ", precio_final)
	} else if n >= 5 && n < 10 {
		precio_inicial := n * computadora
		descuento := precio_inicial / 100 * 20
		precio_final := precio_inicial - descuento
		fmt.Println("Tienes un descuento del 20%, total de la compra: ", precio_final)
	} else {
		precio_inicial := n * computadora
		descuento := precio_inicial / 100 * 40
		precio_final := precio_inicial - descuento
		fmt.Println("Tienes un descuento del 40%, total de la compra: ", precio_final)
	}
}
