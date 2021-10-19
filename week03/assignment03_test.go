package main

import (
	"fmt"
	"log"
	"testing"
)

func TestCritique(t *testing.T) {
	movieList := []*Movie{
		{Length: 100, Name: "Pulp Fiction", play: 2, rating: 90},
		{Length: 140, Name: "True Romance"},
		{Length: 106, Name: "Reservoir Dogs", play: 4, rating: 80},
	}

	th := Theatre{viewers: 20}

	err := th.Critique(movieList, func(m *Movie) (float32, error) {
		return 0.0, fmt.Errorf("oops")
	})

	if err == nil {
		log.Fatalf("expected nil instead")
	}

}
