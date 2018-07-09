package stock

import (
	"fmt"
	"time"
)

// Calculate ...
type Calculate interface {
	ProcessStock()
}

// ProcessStock ...
func (s *stock) ProcessStock() {
	now := time.Now()
	year := now.Year()
	eps := s.eps
	fmt.Println("")

	for i := 0; i < 5; i++ {
		fmt.Print("Estimate EPS ", year, ": ")
		fmt.Printf("%.2f\n", eps)

		eps = eps * 1.15
		year++
	}
}
