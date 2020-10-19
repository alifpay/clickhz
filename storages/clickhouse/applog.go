package clickhouse

import (
	"log"
)

//AppLog batch insert of app's logs
func (d *db) AppLog(data [][]interface{}) {
	tx, err := d.conn.Begin()
	if err != nil {
		log.Println("tx, err := p.conn.Begin()", err)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO applogs(errType, caller, message, dt) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("tx.Prepare", err)
		return
	}

	for _, r := range data {
		_, err = stmt.Exec(r...)
		if err != nil {
			log.Println("stmt.Exec(record...):", err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Println("err = tx.Commit():", err)
		return
	}
	return
}
