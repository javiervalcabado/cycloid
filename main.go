package main

import (
	"fmt"

	"cycloid/karatechop"
	"cycloid/wordchains"
)

func main() {

	fmt.Println("Starting program...")

	startKarateChop()
	fmt.Println("**********")
	startWordChains()
}

func startKarateChop() {
	karatechop.Start()
}

func startWordChains() {
	wordchains.Start("lead", "gold")
}
