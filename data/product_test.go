package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "up",
		Price: 1.0,
		SKU:   "adc-adc-adc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
