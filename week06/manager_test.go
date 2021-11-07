package week06

import (
	"fmt"
	"testing"
)

func Test_Manager_No_Employees(t *testing.T) {

	m := NewManager()

	// Pass 0 employees to the manager
	err := m.Start(0)

	// Check to ensure Employee count error fires
	if err == nil {
		// fmt.Printf("employee count invalid - expected > 0 - got %d", err)
		t.Fatalf("employee count invalid - expected > 0 - got %d", err)
	}

}

func Test_Manager_Stopped(t *testing.T) {

	m := NewManager()

	err := m.Start(1)

	if err != nil {
		t.Fatal(err)
	}

	buildCount := 1

	m.stopped = true
	exp := "manager is stopped"

	for i := 0; i < buildCount; i++ {
		go func() {
			err := m.Assign(&Product{Quantity: 2})

			act := err.Error()

			if err != nil {
				fmt.Printf("expected %s, got %s", exp, act)
				m.Errors() <- err
			}

		}() // go

	}

}

func Test_Manager_Workers(t *testing.T) {
	t.Parallel()

	m := NewManager()

	table := []struct {
		name    string
		workers int
	}{
		{name: "Single Worker", workers: 1},
		{name: "Multiple Workers", workers: 10},
	}

	for _, tt := range table {

		t.Run(tt.name, func(t *testing.T) {
			err := m.Start(tt.workers)
			if err != nil {
				t.Fatal(err)
			}

			buildCount := 5

			for i := 0; i < buildCount; i++ {
				go func() {
					err := m.Assign(&Product{Quantity: 2})
					if err != nil {
						m.Errors() <- err
					}
				}() // go
			}
		})
	}
}

func Test_Complete_Products(t *testing.T) {

	t.Parallel()

	m := NewManager()
	var e Employee = 1
	p := &Product{Quantity: 1}
	p.Build(1)

	go func() {
		err := m.Complete(e, p)

		if err != nil {
			t.Errorf("err %d", err)
		}
	}()

}

func Test_Manager_Stop_Channel(t *testing.T) {

	t.Parallel()
	m := NewManager()
	m.Stop()

	// Confirm channel is closed
	close, ok := <-m.quit

	if ok {
		t.Errorf("closed channel %s", close)
	}
}
