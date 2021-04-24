package clickhouse

import (
	"database/sql/driver"
	"log"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go"
	ClickHouseDwData "github.com/ClickHouse/clickhouse-go/lib/data"
)

type clickhouseRepository struct {
	client clickhouse.Clickhouse
	table  string
}

//NewClickHouseDWClient test
func (c *clickhouseRepository) NewClickHouseDWClient() (client clickhouse.Clickhouse, err error) {
	client, err = clickhouse.OpenDirect("tcp://ch.nb.lan:9000?username=default&password=y61327061&&debug=true&compress=1")
	if err != nil {
		log.Fatal(err)
	}
	c.client = client
	return c.client, nil
}

func ClickHouseCreateTable(client clickhouse.Clickhouse, table string) (err error) {
	client.Begin()
	stmt, _ := client.Prepare(`
		CREATE TABLE IF NOT EXISTS ` + table + ` (
			data         String
		) engine=Memory
	`)
	if _, err := stmt.Exec([]driver.Value{}); err != nil {
		log.Fatal(err)
	}

	if err := client.Commit(); err != nil {
		log.Fatal(err)
	}
	return err
}

//Store Store
func (c *clickhouseRepository) Store(data string) (err error) {
	ClickHouseCreateTable(c.client, c.table)
	c.client.Begin()
	c.client.Prepare("INSERT INTO " + c.table + " (data) VALUES (" + data + ")")
	block, err := c.client.Block()
	if err != nil {
		log.Fatal(err)
	}
	blocks := []*ClickHouseDwData.Block{block, block.Copy()}
	var wg sync.WaitGroup
	wg.Add(len(blocks))
	for i := range blocks {
		b := blocks[i]
		go func() {
			defer wg.Done()
			writeBatch(b, 1)
			if err := c.client.WriteBlock(b); err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()
	if err := c.client.Commit(); err != nil {
		log.Fatal(err)
	}
	return err
}

//Find Find
func (c *clickhouseRepository) Find(selectQuery string) (rows driver.Rows, err error) {
	c.client.Begin()
	stmt, _ := c.client.Prepare(`SELECT count() FROM ` + c.table + ``)
	rows, err = stmt.Query([]driver.Value{})
	if err != nil {
		log.Fatal(err)
	}
	columns := rows.Columns()
	row := make([]driver.Value, 1)
	for rows.Next(row) == nil {
		for i, c := range columns {
			log.Print(c, " : ", row[i])
		}
	}
	if err := c.client.Commit(); err != nil {
		log.Fatal(err)
	}
	return rows, err

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
}

func writeBatch(block *ClickHouseDwData.Block, n int) {
	block.Reserve()
	block.NumRows += uint64(n)

	for i := 0; i < n; i++ {
		block.WriteUInt8(0, uint8(10+i))
	}

	for i := 0; i < n; i++ {
		block.WriteDate(1, time.Now())
	}

	for i := 0; i < n; i++ {
		block.WriteArray(2, clickhouse.Array([]string{"A", "B", "C"}))
	}

	for i := 0; i < n; i++ {
		block.WriteArray(3, clickhouse.Array([]uint8{1, 2, 3, 4, 5}))
	}
}
