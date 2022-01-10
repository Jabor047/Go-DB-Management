package s5_dbase

import (
	"database/sql"
	"fmt"
)

func CreateItemsTable(dbase *sql.DB) error{
	itemStmt, err := dbase.Prepare(`
		CREATE TABLE items(
			"giver_id" text NOT NULL,
			"receiver_id" text,
			"name" text,
			"description" text,
			"image" blob,
			FOREIGN KEY (giver_id) REFERENCES users (user_name),
			FOREIGN KEY (receiver_id) REFERENCES users (user_name)
		)
	`)
	if err != nil {
		return err
	}

	_, err = itemStmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func AddItem(db *sql.DB, giver_id, item_name, desc string, image []byte) error {
	
	if len(image) > 100_000 {
		return fmt.Errorf("image too big for %s", item_name)
	}

	stmt, err := db.Prepare(`INSERT INTO items (giver_id, name, description, image) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("could not prepare item creator %s: %w", item_name, err)
	}

	_, err = stmt.Exec(giver_id, item_name, desc, image)
	if err != nil {
		return fmt.Errorf("could not add %s to Database : %w", item_name, err)
	}

	return nil
}