package week06

import (
	"fmt"
	"testing"
)

func Test_InvalidQty_String(t *testing.T) {

	var qty ErrInvalidQuantity = 1

	exp := fmt.Sprintf("quantity must be greater than 0, got %d", qty)
	act := qty.Error()

	if exp != act {
		t.Errorf("expected %s, got %s", exp, act)
	}

}

func Test_Product_Not_Built(t *testing.T) {

	var nop ErrProductNotBuilt = "IPad"

	exp := string(nop)
	act := nop.Error()

	if exp != act {
		t.Errorf("exp %s, got %s", exp, act)
	}

}

func Test_Invalid_Employee_String(t *testing.T) {

	var qty ErrInvalidEmployee = 1

	exp := fmt.Sprintf("invalid employee number: %d", qty)
	act := qty.Error()

	if exp != act {
		t.Errorf("expected %s, got %s", exp, act)
	}

}

func Test_Invalid_Employee_Count_String(t *testing.T) {

	var qty ErrInvalidEmployeeCount = 1

	exp := fmt.Sprintf("invalid employee count: %d", qty)
	act := qty.Error()

	if exp != act {
		t.Errorf("expected %s, got %s", exp, act)
	}

}
