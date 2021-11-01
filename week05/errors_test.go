package demo

import (
	"fmt"
	"testing"
)

func Test_tablenotfound(t *testing.T) {
	t.Parallel()

	tn := "users"

	errTable := ErrTableNotFound{
		table: tn,
	}

	if errTable.Error() == "" {
		t.Fatalf("table %s not found", errTable.table)
	}

}

func Test_Table_Errors2(t *testing.T) {
	t.Parallel()

	s := &Store{}
	tn := "users"

	s.Insert(tn, Model{})

	errTable := ErrTableNotFound{
		table: tn,
	}

	if !IsErrTableNotFound(&errTable) {
		t.Fatalf("table not found %v", errTable.table)
	}

}

func Test_Table_Errors3(t *testing.T) {
	t.Parallel()

	errTable := ErrTableNotFound{
		table: "users",
	}

	alan := errTable.TableNotFound()
	val := fmt.Sprintf("table not found %s", errTable.table)
	if alan == val {
		t.Fatalf("table %s not found", errTable.table)
	}

}

func Test_Row_not_found(t *testing.T) {
	t.Parallel()

	c := Clauses{
		"age": 31,
	}

	valRows := errNoRows{
		clauses: c,
		table:   "users",
	}

	s := &Store{}

	s.Insert("users", Model{
		"age": 31,
	})

	exp := valRows.Clauses()

	if exp == nil {
		t.Fatalf("%v", exp)
	}

	res := valRows.Error()

	if res == "" {
		t.Fatalf("%s", res)
	}

	val, cl := valRows.RowNotFound()

	if val == "" {
		t.Fatalf("Row not found in table %s with clause %v", val, cl)
	}

}
