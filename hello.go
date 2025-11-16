package main
import "fmt"


func main() {
	var idade int
	var nome string
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idade)
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Printf("Hello!, %s, você tem %d anos!\n", nome, idade)
}
