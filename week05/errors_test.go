package demo

import (
	"fmt"
	"testing"
)

func Test_Table_Errors2(t *testing.T) {
	t.Parallel()

	s := &Store{}
	tn := "users"

	s.Insert(tn, Model{})

	errTable := ErrTableNotFound{
		table: tn,
	}

	if IsErrTableNotFound(&errTable) {
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

// func Test_Row_not_found(t *testing.T) {
// 	t.Parallel()

// 	// c := Clauses{
// 	// 	"name":   "Age test",
// 	// 	"age":    31,
// 	// 	"expect": `"age": = '31'`,
// 	// }

// 	// c2 := Clauses{
// 	// 	"age":  31,
// 	// 	"name": "John",
// 	// }

// 	valRows := errNoRows{
// 		clauses: map[string]interface{}{},
// 		// clauses: c,
// 		table: "users",
// 	}

// 	if valRows.clauses.String()

// }
