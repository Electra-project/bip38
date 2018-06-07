package main

import (
	"github.com/Electra-project/bip38/src/wif"
)

func main() {
	privateKey := "QsmQyTSvpDrwo8Ma4KxKcGYq5Q5Gaoa6HadbNCGReyscffZk1D2F"

	wif.Decode(privateKey)
}
