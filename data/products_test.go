package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "maysara",
		Price: 14.5,
		SKU:   "maysara-loves-bigbaik",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
