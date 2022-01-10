package s5_dbase

import (
	"database/sql"
	"fmt"
	"io"
)

func CreateItemsTable(dbase *sql.DB) error{
	itemStmt, err := dbase.Prepare(`
		CREATE TABLE items(
			"item_id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
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

func GetImage(db *sql.DB, item_id int, w io.Writer) error {
	row := db.QueryRow(`SELECT image FROM items WHERE item_id=?;`, item_id)
	
	var pic *sql.RawBytes
	if err := row.Scan(&pic); err != nil {
		return fmt.Errorf("could not get pic %d : %w", item_id, err)
	}

	w.Write(*pic)
	return nil
}

type ItemInfo struct {
	item_id int
	giver_id string
	name string
	description string
}

func SearchItem(db *sql.DB, info string)([]ItemInfo, error) {
	nm := fmt.Sprintf("%%%s%%", info)
	rows, err := db.Query(` SELECT item_id, giver_id, name, description From items WHERE description LIKE ? OR name LIKE ?`, nm, nm)
	
	res := []ItemInfo{}
	if err != nil {
		return res, fmt.Errorf("could not read items: %w", err)
	}

	for rows.Next() {
		var (
			item_id int
			giver_id string
			name string
			description string
		)
		if err = rows.Scan(&item_id, &giver_id, &name, &description); err != nil {
			return res, fmt.Errorf("could not read row : %w", err)
		}
		res = append(res, ItemInfo{item_id: item_id, giver_id: giver_id, name: name, description: description})
	}

	return res, 	nil
}

func SetReceiver(db *sql.DB, item_id int, receiver_id string) error {
	stmt, err := db.Prepare(`UPDATE items SET reciever_id = ? WHERE item_id = ?;`)
	if err != nil {
		return fmt.Errorf("statement bad %w", err)
	}
	_, err = stmt.Exec(receiver_id, item_id)
	if err != nil {
		return fmt.Errorf("could not set receiver: %w",  err)
	}

	return nil
}