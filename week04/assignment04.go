package main

import (
	"bytes"
	"fmt"
)

var bb bytes.Buffer

type Entertainer interface {
	Name() string
	Perform(v Venue) error
	Entertain(audience int, acts []Entertainer) error
}

type Setuper interface {
	Setup(v Venue) error
}

type Teardowner interface {
	Teardown(v Venue) error
}

// type Venue struct {
// 	Audience int
// 	Log      io.Writer
// }

type Entertain struct {
	Firstname string
	Act       string
	Set       bool
	Tear      bool
}

func (e Entertain) Name() string {
	return e.Firstname
}

func (e Entertain) Perform(v Venue) error {
	log := fmt.Sprintf("%s has performed for %d people.\n", e.Name(), v.Audience)
	bb.Write([]byte(log))

	return nil
}

func (e Entertain) Entertain(audience int, acts []Entertainer) error {
	e.Setup(Venue{})
	e.Teardown(Venue{})
	return nil
}

func (e Entertain) Setup(v Venue) error {
	if e.Set {
		log := fmt.Sprintf("%s has completed setup\n", e.Firstname)
		bb.Write([]byte(log))
	}
	return nil
}

func (e Entertain) Teardown(v Venue) error {
	if e.Tear {
		log := fmt.Sprintf("%s has completed teardown\n", e.Firstname)
		bb.Write([]byte(log))
	}
	return nil
}

func EntertainerName(e Entertainer) {
	e.Name()
}

func main() {

	d := Venue{Audience: 10, Log: &bb}

	entList := []*Entertain{
		{Firstname: "Bob", Act: "pianist", Set: true, Tear: true},
		{Firstname: "Jay", Act: "poet"},
		{Firstname: "John", Act: "comedian"},
		{Firstname: "Amy", Act: "rockstar", Set: true, Tear: true},
	}

	for _, v := range entList {
		v.Perform(d)
		v.Entertain(d.Audience, []Entertainer{})
	}

	fmt.Print(d.Log)

}
