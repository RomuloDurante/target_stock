package main

import (
	"github.com/RomuloDurante/target_stock/stock"
)

func main() {

	var calc stock.Calculate

	// we need pass the adress value to interface or create an object pointer to it
	// -> acao = stock.CreateStock("Ross", 86.32, 3.58, 3.01, 22.42, 0.64, 17.88) cria o objeto
	// -> calc = acao aponta o endereço da memoria para a interface poder usa-lo

	calc = stock.CreateStock("Ross", 86.32, 3.58, 3.01, 22.42, 0.64, 17.88)
	// fmt.Println(calc) // pointer to stock object
	calc.ProcessStock(15)
}
