package main

import (
	"fmt"
	"io"
)

type Venue struct {
	Audience int
	Log      io.Writer
}

func (v *Venue) Entertain(audience int, acts ...Entertainer) error {
	v.Audience = audience
	for _, act := range acts {
		if err := v.Play(act); err != nil {
			return err
		}
	}
	return nil
}

func (v Venue) Play(act Entertainer) error {

	name := act.Name()
	if s, ok := act.(Setuper); ok {
		if err := s.Setup(v); err != nil {
			return fmt.Errorf("%s: %w", name, err)
		}
		fmt.Fprintf(v.Log, "%s has completed setup\n", name)
	}
	return nil
}
