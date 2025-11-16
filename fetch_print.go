package main

import "fmt"


func fetchUserdata() (string, int) {
	var n string
	var i int
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&i)
	fmt.Println("Digite seu Nome: ")
	fmt.Scan(&n)
	return n, i
}

func printMessage(n string, i int) {
	fmt.Printf("Olá, %s, tudo bem? Você tem %d anos!\n", n, i)
}

