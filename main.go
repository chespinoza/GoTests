package main

import (
	"fmt"
	"github.com/chespinoza/GoTests"
	"github.com/spf13/nitro"
)

func main() {
	timer := nitro.Initalize()
	timer.Step("stringConcatenate()")
	fmt.Println(stringConcatenate("Hola ", "Mundo"))
	timer.Step("concat()")
	fmt.Println(concat("Hola ", "Mundo"))
	timer.Step("conc()")
	fmt.Println(conc("Hola", " ", "Mundo"))

}
