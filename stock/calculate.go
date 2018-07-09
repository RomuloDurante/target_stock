package stock

import (
	"fmt"
	"time"
)

// Calculate ...
type Calculate interface {
	ProcessStock(float64)
}

// ProcessStock ...
func (s *stock) ProcessStock(p float64) {
	percent := (p / 100) + 1
	now := time.Now()
	year := now.Year()
	eps := s.eps
	pl := pl(s)
	total := 0.

	show(s)

	fmt.Println("You will use 15% to calculate")
	fmt.Println("-----------------------------")

	for i := 0; i < 5; i++ {
		eps = eps * percent
		year++
		total += eps

		fmt.Print("Estimate EPS ", year, ": ")
		fmt.Printf("%.2f\n", eps)
	}

	fmt.Println("-----------------------------")
	finalPrice := pl*eps + total*s.payout
	fmt.Print("Final Price R$-> ")
	fmt.Printf("%.2f\n", finalPrice)
	fmt.Println("-----------------------------")
	for i := 5; i > 0; i-- {
		finalPrice = finalPrice / percent
		fmt.Print("Estimate Price to return ", p, "% ", year, ": R$ ")
		fmt.Printf("%.2f\n", finalPrice/percent)

		year--

	}

}

//show stock
func show(s *stock) {
	fmt.Printf("%-20s%s\n", "Name:", s.name)
	fmt.Printf("%-20s%.2f\n", "Price:", s.price)
	fmt.Printf("%-20s%.2f\n", "Eps:", s.eps)
	fmt.Printf("%-20s%.2f\n", "Dividend:", s.dividend)
	fmt.Printf("%-20s%.2f\n", "Eps:", s.payout)
	fmt.Println("-----------------------------")
}

// Calc pl ...
func pl(s *stock) float64 {
	if s.peRatio <= 20 {
		return 12
	}
	return 17

}
