package week06

import (
	"testing"
)

func Test_Product_BuiltBy(t *testing.T) {

	qty := 0

	p := Product{Quantity: qty}

	exp := p.BuiltBy()
	act := Employee(p.Quantity)

	if exp != Employee(p.Quantity) {
		t.Errorf("expected %b, got %b", exp, act)
	}

}

func Test_Invalid_Product(t *testing.T) {

	qty := 0

	p := Product{Quantity: qty}

	err := p.IsValid()

	if err == nil {
		t.Errorf("exp > 0, got %b", p.IsValid())
	}

}

func Test_IsBuilt_Error(t *testing.T) {

	qty := 0

	p := Product{Quantity: qty}

	p.Build(1)
	err := p.IsBuilt()

	if err == nil {
		t.Errorf("exp %b", err)
	}

}
