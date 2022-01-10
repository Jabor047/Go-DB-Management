package s5_dbase

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB(fname string) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	file.Close()

	dbase, err := sql.Open("sqlite3", fname)
	if err != nil {
		return err
	}

	defer dbase.Close()

	if err = CreateUsersTables(dbase); err != nil {
		return err
	}

	if err = CreateItemsTable(dbase); err != nil {
		return err
	}

	return nil
}
