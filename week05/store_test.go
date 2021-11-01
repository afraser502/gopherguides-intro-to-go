package demo

import (
	"errors"
	"testing"
)

func Test_db_data(t *testing.T) {
	t.Parallel()

	s := &Store{}

	s.Insert("users", Model{
		"name": "John",
		"age":  30,
	})
}

func Test_Selectstore(t *testing.T) {
	t.Parallel()

	s := Store{}
	tn := ","

	s.Insert(tn, Model{
		"name": "John",
		"age":  30,
	})

	exp, err := s.Select(tn, Clauses{})

	if err != nil {
		t.Fatalf("Expected %s, got %s", exp, err)
	}

	if len(exp) != 1 {
		t.Fatalf("Expected 1 record")
	}
}

func Test_Selectstore_error(t *testing.T) {
	t.Parallel()

	s := Store{}

	tn := "alan"
	new := "brian"

	s.Insert(tn, Model{})

	exp, err := s.Select(new, Clauses{})

	if err == nil {
		t.Errorf("Expected %s, got %s", exp, err)
	}
}

func Test_select_result(t *testing.T) {
	t.Parallel()
	s := Store{}

	tn := "alan"

	s.Insert(tn, Model{
		"name": "John",
		"age":  30,
	})

	exp, err := s.Select(tn, Clauses{"age": 30})
	if err != nil {
		t.Fatalf("expected %v, got %v", exp, err)
	}

}

func Test_select_result_returned_rows(t *testing.T) {
	t.Parallel()
	s := Store{}

	tn := "alan"

	s.Insert(tn, Model{
		"name": "John",
		"age":  30,
	})

	exp, err := s.Select(tn, Clauses{"age": 31})
	if err == nil {
		t.Fatalf("expected %v, got %v", exp, err)
	}

}

func Test_no_clauses(t *testing.T) {
	t.Parallel()
	s := Store{}

	tn := "alan"

	s.Insert(tn, Model{
		"name": "John",
		"age":  30,
	})

	exp, err := s.Select(tn, Clauses{})
	if err != nil {
		t.Fatalf("expected %v, got %v", exp, err)
	}

}

func Test_TableDrivenTests_Anatomy(t *testing.T) {
	t.Parallel()

	s := Store{}
	tngood := "users"
	tnbad := "random"

	table := []struct {
		name string
		tng  string
		tnb  string
	}{
		{name: "Select-data", tng: tngood},
		{name: "Select-data-error", tng: tngood, tnb: tnbad},
	}

	for _, tt := range table {
		s.Insert(tngood, Model{})
		exp, err := s.Select(tt.tng, Clauses{})
		if err != nil {
			t.Fatalf("Expected %s, got %s", exp, err)
		}
	}
}

func Test_Store_All_Errors(t *testing.T) {
	t.Parallel()

	tn := "users"

	// noData := &Store{}
	withData := &Store{data: data{}}
	withUsers := &Store{data: data{"users": Models{}}}

	table := []struct {
		store *Store
		exp   error
	}{
		// {store: noData, exp: ErrNoData(tn)},
		{store: withData, exp: ErrTableNotFound{}},
		{store: withUsers, exp: nil},
	}

	for _, tt := range table {
		_, err := tt.store.All(tn)

		ok := errors.Is(err, tt.exp)

		if !ok {
			t.Fatalf("expected error %v, got %v", tt.exp, err)
		}
	}
}

func Test_Length(t *testing.T) {
	t.Parallel()

	s := Store{}

	tn := "users"

	s.Insert(tn, Model{})

	exp, err := s.Len(tn)
	if err != nil {
		t.Fatalf("Expected %#v, got %#v", exp, err)
	}

}

func Test_Length_error(t *testing.T) {
	t.Parallel()

	s := Store{}

	tnexpected := "users"
	tnfail := "payment"

	s.Insert(tnexpected, Model{})

	exp, err := s.Len(tnfail)
	if err == nil {
		t.Fatalf("Expected %#v, got %#v", exp, err)
	}

}
