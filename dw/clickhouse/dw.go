package main

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gookit/config"
)

// env, err := GetNativeTestEnvironment()
// if err != nil {
// 	return err
// }
// add Decoder and Encoder
// config.WithOptions(config.ParseEnv)

// add Decoder and Encoder
// config.AddDriver(toml.Driver)
// err := config.LoadFiles("/home/ml/nabla/go/src/gitlab/nbgo/conf/db.toml")
// if err != nil {
// 	panic(err)
// }

var Host, _ = config.String("db.clickhouse.Host")
var Port, _ = config.String("db.clickhouse.Port")
var Database, _ = config.String("db.clickhouse.Database")
var Username, _ = config.String("db.clickhouse.Username")
var Password, _ = config.String("db.clickhouse.Password")

func ClickHouseConnect() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: Database,
			Username: Username,
			Password: Password,
		},
	})
	if err != nil {
		return err
	}
	v, err := conn.ServerVersion()
	fmt.Println(v)
	if err != nil {
		return err
	}
	return nil
}

func ClientInfo() error {
	conn, err := clickhouse.Open(&clickhouse.Options{
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "my-app", Version: "0.1"},
			},
		},
	})
	if err != nil {
		return err
	}
	return conn.Exec(context.TODO(), "SELECT 1")
}

// type clickhouseDB struct {
// 	client *clickhouse.Conn
// }

// func newCLickhoseClient() {
// 	conn, err := clickhouse.Open(&clickhouse.Options{
// 		Addr: []string{"127.0.0.1:9000"},
// 		Auth: clickhouse.Auth{
// 			Database: "default",
// 			Username: "default",
// 			Password: "",
// 		},
// 		DialContext: func(ctx context.Context, addr string) (net.Conn, error) {
// 			dialCount++
// 			var d net.Dialer
// 			return d.DialContext(ctx, "tcp", addr)
// 		},
// 		Debug: true,
// 		Debugf: func(format string, v ...any) {
// 			fmt.Printf(format+"\n", v...)
// 		},
// 		Settings: clickhouse.Settings{
// 			"max_execution_time": 60,
// 		},
// 		Compression: &clickhouse.Compression{
// 			Method: clickhouse.CompressionLZ4,
// 		},
// 		DialTimeout:          time.Second * 30,
// 		MaxOpenConns:         5,
// 		MaxIdleConns:         5,
// 		ConnMaxLifetime:      time.Duration(10) * time.Minute,
// 		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
// 		BlockBufferSize:      10,
// 		MaxCompressionBuffer: 10240,
// 		ClientInfo: clickhouse.ClientInfo{ // optional, please see Client info section in the README.md
// 			Products: []struct {
// 				Name    string
// 				Version string
// 			}{
// 				{Name: "my-app", Version: "0.1"},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return conn.Ping(context.Background())
// }

// // // CREATE TABLE test_on_fly_mutations (id UInt64, v String)
// // ENGINE = MergeTree ORDER BY id;

// // -- Disable background materialization of mutations to showcase
// // -- default behavior when lightweight updates are not enabled
// // SYSTEM STOP MERGES test_on_fly_mutations;
// // SET mutations_sync = 0;

// // -- Insert some rows in our new table
// // INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

// // -- Update the values of the rows
// // ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
// // ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';

// // CREATE DATABASE IF NOT EXISTS level2

// // CREATE TABLE level1.gateio
// // (
// //     instrument_id UInt32,
// //     message String,
// //     timestamp DateTime,
// //     metric Float32
// // )
// // ENGINE = MergeTree()
// // PRIMARY KEY (instrument_id, timestamp)

// // CREATE TABLE test_on_fly_mutations (id UInt64, v String)
// // ENGINE = MergeTree ORDER BY id;

// // CREATE TABLE test_on_fly_mutations (id UInt64, v String)
// // ENGINE = MergeTree ORDER BY id;

// // -- Disable background materialization of mutations to showcase
// // -- default behavior when lightweight updates are not enabled
// // SYSTEM STOP MERGES test_on_fly_mutations;
// // SET mutations_sync = 0;

// // -- Insert some rows in our new table
// // INSERT INTO test_on_fly_mutations VALUES (1, 'a'), (2, 'b'), (3, 'c');

// // -- Update the values of the rows
// // ALTER TABLE test_on_fly_mutations UPDATE v = 'd' WHERE id = 1;
// // ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'd';
// // ALTER TABLE test_on_fly_mutations UPDATE v = 'e' WHERE id = 2;
// // ALTER TABLE test_on_fly_mutations DELETE WHERE v = 'e';

// // Let's check the result of the updates via a SELECT query:

// // -- Explicitly disable lightweight updates
// // SET apply_mutations_on_fly = 0;

// // SELECT id, v FROM test_on_fly_mutations ORDER BY id;

// // CREATE TABLE IF NOT EXIST myFirstReplacingMT
// // (
// //     `key` Int64,
// //     `someCol` String,
// //     `eventTime` DateTime
// // )
// // ENGINE = ReplacingMergeTree
// // ORDER BY key;

// // INSERT INTO myFirstReplacingMT Values (1, 'first', '2020-01-01 01:01:01');
// // INSERT INTO myFirstReplacingMT Values (1, 'second', '2020-01-01 00:00:00');

// // SELECT * FROM myFirstReplacingMT FINAL;
