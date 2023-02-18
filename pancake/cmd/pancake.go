package main

import "pancake/services"

func main() {
	bg := &services.BancoGeneralTxt{}
	bg.ParseTxtFile("./sources/banco-general/bg-sample.txt")
}
