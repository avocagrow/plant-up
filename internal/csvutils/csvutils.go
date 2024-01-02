package csvutils

import (
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type CoffeeBean struct {
	Name     string  `csv:"name"`
	Roast    string  `csv:"roast"`
	Price    float64 `csv:"price"`
	Quantity int     `csv:"quantity"` // quantity is measures in Kilograms
}

// LoadCSVForBeans loads a CSV file into a slice of CoffeeBeans
func LoadCSVForBeans() ([]*CoffeeBean, error) {
	p := filepath.Join("..", "..", "beans.csv")
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}

	beans := make([]*CoffeeBean, 5)
	defer f.Close()
	if err := gocsv.UnmarshalFile(f, beans); err != nil {
		return nil, err
	}

	return beans, nil
}
