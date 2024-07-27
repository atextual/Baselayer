package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"regexp"
	"testing"
)

func TestListDatabases(t *testing.T) {
	cxn, mock := MockConnection()
	databaseMockRows := sqlmock.NewRows([]string{"id", "name", "normalised_name", "driver", "username", "password", "database", "port"}).
		AddRow("1", "Test Fixture DB #1", "TEST FIXTURE DB #1", "mysql", "root", "root", "text_fixture_db_1", "3306")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM databases")).
		WithoutArgs().
		WillReturnRows(databaseMockRows)

	list, err := ListDatabases(cxn.Db)

	if err != nil {
		panic(err)
	}

	got := len(list)
	want := 1

	if got != want {
		t.Errorf("got %v items from mock db, wanted %v", got, want)
	}
}

func TestGetDatabase(t *testing.T) {
	cxn, mock := MockConnection()
	databaseMockRows := sqlmock.NewRows([]string{"id", "name", "normalised_name", "driver", "username", "password", "database", "port"}).
		AddRow("1", "Test Fixture DB #1", "TEST FIXTURE DB #1", "mysql", "root", "root", "text_fixture_db_1", "3306")

	mock.ExpectQuery("^SELECT \\* FROM databases WHERE id = \\?").
		WithArgs(1).
		WillReturnRows(databaseMockRows)

	db, err := GetDatabase(cxn.Db, 1)

	if err != nil {
		panic(err)
	}

	got := db.Id
	want := 1

	if got != want {
		t.Errorf("Got id %d from mock db, wanted id %d", got, want)
	}
}
