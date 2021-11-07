package week06

import "testing"

func Test_Completed_Product(t *testing.T) {

	t.Parallel()

	c := CompletedProduct{Product: Product{Quantity: 1}, Employee: 1}

	err := c.IsValid()

	if err == nil {
		t.Errorf("err %d", err)
	}

}
