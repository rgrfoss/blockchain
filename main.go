package main

import "fmt"

func main() {

	blockchain := CreateBlockchain(2)

	blockchain.addBlock("James", "Bob", 5)
	blockchain.addBlock("Sara", "Bob", 2)

	fmt.Println(blockchain.isValid())
}
