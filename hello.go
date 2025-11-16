package main
import "fmt"


func main() {
	var idade int
	var nome string
	var dinheiro float64
	fmt.Println("Digite sua idade: ")
	fmt.Scan(&idade)
	fmt.Println("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Println("Digite seu saldo bancário: ")
	fmt.Scan(&dinheiro)
	fmt.Printf("Hello!, %s, você tem %d anos!\nPossui %.2f na conta\n", nome, idade, dinheiro)
}
