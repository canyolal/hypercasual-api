package data

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// Establish a new DB and connection pool for testing
func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("HYPERCASUAL_TEST_DSN"))
	if err != nil {
		t.Fatal(err)
	}

	//Read the setup SQL script from file and execute the statements.
	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})
	return db
}
