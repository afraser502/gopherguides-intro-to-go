package main

import (
	"fmt"
	"log"
)

type Movie struct {
	Length  int     // Running length of movie
	Name    string  // name of movie
	rating  float32 // Movie rating
	play    int     // Number of times movie has been played
	viewers int     // Total number of viewers
}

//Function to add movie rating
func (m *Movie) Rate(r float32) error {

	if m.play == 0 {
		return fmt.Errorf("can't review a movie without watching it first")
	} else {
		m.rating += r
		return fmt.Errorf("Rating accepted")
	}

}

// Function to add number of viewers to a movie + increment plays
func (m *Movie) Play(v int) {
	m.play++
	m.viewers += v
}

// Function to return number of viewers
func (m Movie) Viewers() int {
	return m.viewers
}

// Function to return number of plays
func (m Movie) Plays() int {
	return m.play
}

// Function to return movie rating
func (m Movie) Rating() float64 {
	return float64(m.rating) / float64(m.play)
}

// Function to return Movie details
func (m Movie) String() string {
	// Use %% to escape the % sign to show percentage in the output
	return fmt.Sprintf("%v (%vm) %.1f%%", m.Name, m.Length, m.Rating())
}

type CritiqueFn func(m *Movie) (float32, error)

type Theatre struct {
	viewers int
}

//Function to add movie rating from theatre viewers
func (t *Theatre) Play(v int, m ...*Movie) error {

	if len(m) == 0 {
		return fmt.Errorf("no movies to play")
	}

	for _, mov := range m {
		mov.Play(v)
	}

	return nil
}

// Function to return critique of movie
func (t Theatre) Critique(m []*Movie, fn CritiqueFn) error {

	for _, mov := range m {
		mov.Play(1)

		f, err := fn(mov)
		if err != nil {
			return err
		}

		err = mov.Rate(f)
		if err != nil {
			return err
		}

	}

	return nil
}

func main() {

	// Generate struct data
	movieList := []*Movie{
		{Length: 100, Name: "Pulp Fiction", play: 2, rating: 90},
		{Length: 140, Name: "True Romance"},
		{Length: 106, Name: "Reservoir Dogs", play: 4, rating: 80},
	}

	q := Movie{Length: 100, Name: "Pulp Fiction", play: 2, rating: 90}
	r := Movie{Length: 140, Name: "True Romance"}
	s := Movie{Length: 106, Name: "Reservoir Dogs", play: 4, rating: 80}

	// Generate pointer variables
	u := &q
	v := &r
	w := &s

	// Add rating of 10 to movie u
	fmt.Println(u.Rate(10))

	// Add 20 viewers
	u.Play(20)
	v.Play(10)
	w.Play(10)

	th := Theatre{viewers: 20}

	err := th.Critique(movieList, func(m *Movie) (float32, error) {
		return 0.0, fmt.Errorf("oops")
	})

	if err != nil {
		log.Fatalf("oops")
	}

	// Show number of viewers, Plays and name
	fmt.Printf("Viewers = %v, Plays = %v, Movie = %s\n", u.Viewers(), u.Plays(), u.Name)
	fmt.Printf("Viewers = %v, Plays = %v, Movie = %s\n", v.Viewers(), v.Plays(), v.Name)
	fmt.Printf("Viewers = %v, Plays = %v, Movie = %s\n", w.Viewers(), w.Plays(), w.Name)

	// Show ratings
	fmt.Println(u)
	fmt.Println(v)
	fmt.Println(w)

	// Theatre components
	// th := Theatre{viewers: 20}
	fmt.Println("Viewers = ", u.Viewers())
	fmt.Println("Plays = ", u.Plays())

	fmt.Println("Viewers = ", v.Viewers())
	fmt.Println("Plays = ", v.Plays())

	fmt.Println("Viewers = ", w.Viewers())
	fmt.Println("Plays = ", w.Plays())

}
