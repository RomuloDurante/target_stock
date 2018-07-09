package stock

// Stock struct
type stock struct {
	name        string
	price       float64
	eps         float64
	tenYearsEps float64
	epsgrowth   float64
	peRatio     float64
	dividend    float64
	payout      float64
}

// CreateStock constructor
func CreateStock(name string, price, eps, tenYearsEps, peRatio, dividend, payout float64) *stock {
	return &stock{
		name:        name,
		price:       price,
		eps:         eps,
		tenYearsEps: tenYearsEps,
		epsgrowth:   (eps*100)/tenYearsEps - 100,
		peRatio:     peRatio,
		dividend:    dividend,
		payout:      payout,
	}
}
