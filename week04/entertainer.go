package main

import (
	"bytes"
	"fmt"
)

var _ Entertainer = &Band{}
var _ Setuper = &Band{}
var _ Teardowner = &Band{}

type Band struct {
	IsSetup    bool
	IsTeardown bool
	PlayedFor  int
}

func (b *Band) Name() string {
	return "My great band"
}

func (b *Band) Perform(v Venue) error {
	b.PlayedFor = v.Audience
	return nil
}

func (b *Band) Setup(v Venue) error {
	if b.IsSetup {
		return fmt.Errorf("%s has completed setup.\n", b.Name())
	}
	b.IsSetup = true
	return nil
}

func (b *Band) Teardown(v Venue) error {
	if b.IsTeardown {
		return fmt.Errorf("%s has completed teardown.\n", b.Name())
	}
	b.IsTeardown = true
	return nil
}

func main() {

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}

	err := v.Entertain(1, &Band{PlayedFor: 10, IsSetup: true})

	if err != nil {
		fmt.Errorf("Error %s", err)
	}

	fmt.Print()

}
