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
	log := fmt.Sprintf("%s has performed for %d people.\n", b.Name(), b.PlayedFor)
	v.Log.Write([]byte(log))
	return nil
}

func (b *Band) Setup(v Venue) error {
	if b.IsSetup {
		log := fmt.Sprintf("%s has completed setup.\n", b.Name())
		v.Log.Write([]byte(log))
		return nil
	}
	b.IsSetup = true
	return nil
}

func (b *Band) Teardown(v Venue) error {
	if b.IsTeardown {
		log := fmt.Sprintf("%s has completed teardown.\n", b.Name())
		v.Log.Write([]byte(log))
		return nil
	}
	b.IsTeardown = true
	return nil
}

func main() {

	bb := &bytes.Buffer{}
	v := &Venue{Log: bb}

	err := v.Entertain(10, &Band{PlayedFor: 10, IsSetup: true})

	if err != nil {
		fmt.Errorf("Error %s", err)
	}

	fmt.Print(v.Log)

}
