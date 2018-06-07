package main

import (
	"github.com/Electra-project/electra-tools/src/wif"
)

func main() {
	privateKey := "QsmQyTSvpDrwo8Ma4KxKcGYq5Q5Gaoa6HadbNCGReyscffZk1D2F"

	wif.Decode(privateKey)
}
