package clickhouse

import (
	"database/sql"

	"github.com/alifpay/clickhz/storages"
)

type db struct {
	conn *sql.DB
}

//Connect clickhouse database
func Connect(dsn string) (storages.ClickDB, error) {

	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return &db{conn: conn}, nil
}

func (d *db) Close() {
	d.conn.Close()
}
