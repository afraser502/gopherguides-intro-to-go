package main

import (
	"testing"
)

func TestArray(t *testing.T) {

	exp := [4]string{"North", "South", "East", "West"}

	// Generate empty 'act' array
	act := [len(exp)]string{}

	// Copy 'exp' array to 'act' array.
	act = exp

	// Iterate through act slice and check for presence of each value in exp slice
	for i, v := range act {
		if v != exp[i] {
			t.Errorf("Value %s not found in array 'exp'.", v)
		}
	}

	// Perform a comparison of arrays
	if act != exp {
		t.Errorf("Returned value %s,  expect %s", act, exp)
	}

}

func TestSlice(t *testing.T) {

	exp := []string{"North", "South", "East", "West"}

	// Set empty 'act' slice
	act := make([]string, (len(exp)))

	// Copy exp slice to act slice
	copy(act, exp)

	// Iterate through act slice and check for presence of each value in exp slice
	for i, v := range act {
		if v != exp[i] {
			t.Errorf("Value %s not found in slice 'exp'.", v)
		}
	}

	// Assert the length is the same between act and exp slices
	if len(act) != len(exp) {
		t.Errorf("The length of 'act' slice is %d and 'exp' slice is %d", len(act), len(exp))
	}

}

func TestMap(t *testing.T) {

	exp := map[string]string{
		"North": "Up",
		"South": "Down",
		"East":  "Right",
		"West":  "Left",
	}

	// Generate the empty `act` map
	act := map[string]string{}

	// Copy the `exp` map to the `act` map
	for k, v := range exp {
		act[k] = v
	}

	// Range over the values in 'act'
	for key, _ := range act {

		// Check for the presence of each key in 'exp'.
		_, ok := exp[key]
		if !ok {
			t.Errorf("Key %s not found in map 'exp'.", key)
		}
	}

	// Test to ensure lengths of maps are the same
	if len(act) != len(exp) {
		t.Errorf("The length of 'act' map is %d and 'exp' map is %d", len(act), len(exp))
	}

}
