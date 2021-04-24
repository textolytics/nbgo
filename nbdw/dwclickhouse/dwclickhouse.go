package dwclickhouse

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ClickHouse/clickhouse-go"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ClickHouseDWClient() (connect *sql.DB) {
	connect, err := sql.Open("clickhouse", "tcp://ch.nb.lan:9000?username=default&password=y61327061&&debug=true&compress=true")
	checkErr(err)
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}
	return connect
}

func ClickHouseDWClientInsert(connect *sql.DB, table string, data string) {

	_, err := connect.Exec(`
		CREATE TABLE IF NOT EXISTS ` + table + ` (
			subzmqmsg String
		) engine=Memory
	`)

	checkErr(err)
	tx, err := connect.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("INSERT INTO " + table + " (subzmqmsg) VALUES (" + data + ")")
	checkErr(err)
	if _, err := stmt.Exec(
		data,
	); err != nil {
		log.Fatal(err)
	}
	checkErr(tx.Commit())
}

// {
// 	connect.Begin()
// 	stmt, _ := connect.Prepare(`SELECT count() FROM example`)

// 	rows, err := stmt.Query([]driver.Value{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	columns := rows.Columns()
// 	row := make([]driver.Value, 1)
// 	for rows.Next(row) == nil {
// 		for i, c := range columns {
// 			log.Print(c, " : ", row[i])
// 		}
// 	}

// 	if err := connect.Commit(); err != nil {
// 		log.Fatal(err)
// 	}
// }
// {
// 	connect.Begin()
// 	stmt, _ := connect.Prepare(`DROP TABLE example`)
// 	if _, err := stmt.Exec([]driver.Value{}); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := connect.Commit(); err != nil {
// 		log.Fatal(err)
// 	}
// }
// }

// func writeBatch(block *ClickHouseDwData.Block, n int) {
// 	block.Reserve()
// 	block.NumRows += uint64(n)

// 	for i := 0; i < n; i++ {
// 		block.WriteUInt8(0, uint8(10+i))
// 	}

// 	for i := 0; i < n; i++ {
// 		block.WriteDate(1, time.Now())
// 	}

// 	for i := 0; i < n; i++ {
// 		block.WriteArray(2, clickhouse.Array([]string{"A", "B", "C"}))
// 	}

// 	for i := 0; i < n; i++ {
// 		block.WriteArray(3, clickhouse.Array([]uint8{1, 2, 3, 4, 5}))
// 	}
// }
