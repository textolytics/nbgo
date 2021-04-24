package clickhouse

import (
	"database/sql"
)

type clickhouseRepository struct {
	client *sql.DB
}

func newClickhouseClient(clickhouseURL string) (client clickhouseRepository, connect *sql.DB, err error) {
	client, connect, err = newClickhouseClient(clickhouseURL)
	connect, err = sql.Open("clickhouse", "tcp://ch.nb.lan:9000?username=default#&compress=true&debug=true")

	// checkErr(err)

	return client, connect, err
}

// //NewClickhouseRepository (clickhouseURL string)
// func NewClickhouseRepository(client *sql.DB, connect *sql.DB, clickhouseURL string) (client *err error) {
// 	repo := &clickhouseRepository{}
// 	// client, connect, err = newClickhouseClient(clickhouseURL)
// 	// // if err != nil {
// 	// // 	return nil, errors.Wrap(err, "repository.NewClickhouseRepository")
// 	// // }
// 	_, err = connect.Exec(`
// 		CREATE TABLE IF NOT EXISTS example (
// 			country_code FixedString(2),
// 			os_id        UInt8,
// 			browser_id   UInt8,
// 			categories   Array(Int16),
// 			action_day   Date,
// 			action_time  DateTime
// 		) engine=Memory
// 	`)
// 	checkErr(err)
// 	repo.client = client
// 	return err
// }

// var INSERTCH = "INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)"

// func (r *clickhouseRepository) Store(connect sql.DB, INSERTCH string) error {
// 	tx, err := connect.Begin()
// 	checkErr(err)
// 	stmt, err := tx.Prepare(INSERTCH)
// 	checkErr(err)

// 	for i := 0; i < 100; i++ {
// 		if _, err := stmt.Exec(
// 			"RU",
// 			10+i,
// 			100+i,
// 			[]int16{1, 2, 3},
// 			time.Now(),
// 			time.Now(),
// 		); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	checkErr(tx.Commit())
// 	return err
// }

// //SELECTCH sdsdfsdf
// var SELECTCH = "SELECT country_code, os_id, browser_id, categories, action_day, action_time FROM example"

// func (r *clickhouseRepository) Find(connect *sql.DB, SELECTCH string) (rows *sql.Rows, err error) {
// 	rows, err = connect.Query(SELECTCH)
// 	checkErr(err)
// 	for rows.Next() {
// 		var (
// 			country               string
// 			os, browser           uint8
// 			categories            []int16
// 			actionDay, actionTime time.Time
// 		)
// 		checkErr(rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime))
// 		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
// 	}
// 	return rows, err
// }

// func (r *clickhouseRepository) DROP(connect *sql.DB, input string) (err error) {

// 	if _, err := connect.Exec("DROP TABLE example"); err != nil {
// 		log.Fatal(err)
// 	}
// 	return err
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func newClickhouseClient(clickhouseURL string) (connect *sql.DB, err error) {
// 	connect, err = sql.Open("clickhouse", "tcp://ch.nb.lan:9000?username=&compress=true&debug=true")
// 	checkErr(err)
// 	if err := connect.Ping(); err != nil {
// 		if exception, ok := err.(*clickhouse.Exception); ok {
// 			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}
// 	return connect, err
// }

// //NewClickhouseRepository (clickhouseURL string)
// func NewClickhouseRepository(connect *sql.DB, clickhouseURL string) (repo *clickhouseRepository, err error) {
// 	repo = &clickhouseRepository{}
// 	client, err := newClickhouseClient(clickhouseURL)
// 	// if err != nil {
// 	// 	return nil, errors.Wrap(err, "repository.NewClickhouseRepository")
// 	// }
// 	_, err = connect.Exec(`
// 		CREATE TABLE IF NOT EXISTS example (
// 			country_code FixedString(2),
// 			os_id        UInt8,
// 			browser_id   UInt8,
// 			categories   Array(Int16),
// 			action_day   Date,
// 			action_time  DateTime
// 		) engine=Memory
// 	`)
// 	checkErr(err)
// 	repo.client = client
// 	return repo, err
// }

// func (r *redisRepository) Find(code string) (*shortener.Redirect, error) {
// 	redirect := &shortener.Redirect{}
// 	key := r.generateKey(code)
// 	data, err := r.client.HGetAll(r.client.Context(), key).Result()
// 	if err != nil {
// 		return nil, errors.Wrap(err, "repository.Redirect.Find")
// 	}
// 	if len(data) == 0 {
// 		return nil, errors.Wrap(shortener.ErrRedirectNotFound, "repository.Redirect.Find")
// 	}
// 	createdAt, err := strconv.ParseInt(data["created_at"], 10, 64)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "repository.Redirect.Find")
// 	}
// 	redirect.Code = data["code"]
// 	redirect.URL = data["url"]
// 	redirect.CreatedAt = createdAt
// 	return redirect, nil
// }

// func (r *redisRepository) Store(redirect *shortener.Redirect) error {
// 	key := r.generateKey(redirect.Code)
// 	data := map[string]interface{}{
// 		"code":       redirect.Code,
// 		"url":        redirect.URL,
// 		"created_at": redirect.CreatedAt,
// 	}
// 	_, err := r.client.HMSet(r.client.Context(), key, data).Result()
// 	if err != nil {
// 		return errors.Wrap(err, "repository.Redirect.Store")
// 	}
// 	return nil
// }

// import (
// 	"bufio"
// 	"crypto/tls"
// 	"database/sql/driver"
// 	"net"
// 	"sync/atomic"
// 	"time"
// )

// var tick int32

// type openStrategy int8

// type clickhouseRepository struct {
// 	client *connect
// }

// func (s openStrategy) String() string {
// 	switch s {
// 	case connOpenInOrder:
// 		return "in_order"
// 	case connOpenTimeRandom:
// 		return "time_random"
// 	}
// 	return "random"
// }

// const (
// 	connOpenRandom openStrategy = iota + 1
// 	connOpenInOrder
// 	connOpenTimeRandom
// )

// type connOptions struct {
// 	secure, skipVerify                     bool
// 	tlsConfig                              *tls.Config
// 	hosts                                  []string
// 	connTimeout, readTimeout, writeTimeout time.Duration
// 	noDelay                                bool
// 	openStrategy                           openStrategy
// 	logf                                   func(string, ...interface{})
// }

// type connect struct {
// 	net.Conn
// 	logf                  func(string, ...interface{})
// 	ident                 int
// 	buffer                *bufio.Reader
// 	closed                bool
// 	readTimeout           time.Duration
// 	writeTimeout          time.Duration
// 	lastReadDeadlineTime  time.Time
// 	lastWriteDeadlineTime time.Time
// }

// func dial(options connOptions) (*connect, error) {
// 	var (
// 		err error
// 		abs = func(v int) int {
// 			if v < 0 {
// 				return -1 * v
// 			}
// 			return v
// 		}
// 		conn  net.Conn
// 		ident = abs(int(atomic.AddInt32(&tick, 1)))
// 	)
// 	tlsConfig := options.tlsConfig
// 	if options.secure {
// 		if tlsConfig == nil {
// 			tlsConfig = &tls.Config{}
// 		}
// 		tlsConfig.InsecureSkipVerify = options.skipVerify
// 	}
// 	checkedHosts := make(map[int]struct{}, len(options.hosts))
// 	for i := range options.hosts {
// 		var num int
// 		switch options.openStrategy {
// 		case connOpenInOrder:
// 			num = i
// 		case connOpenRandom:
// 			num = (ident + i) % len(options.hosts)
// 		case connOpenTimeRandom:
// 			// select host based on milliseconds
// 			num = int((time.Now().UnixNano()/1000)%1000) % len(options.hosts)
// 			for _, ok := checkedHosts[num]; ok; _, ok = checkedHosts[num] {
// 				num = int(time.Now().UnixNano()) % len(options.hosts)
// 			}
// 			checkedHosts[num] = struct{}{}
// 		}
// 		switch {
// 		case options.secure:
// 			conn, err = tls.DialWithDialer(
// 				&net.Dialer{
// 					Timeout: options.connTimeout,
// 				},
// 				"tcp",
// 				options.hosts[num],
// 				tlsConfig,
// 			)
// 		default:
// 			conn, err = net.DialTimeout("tcp", options.hosts[num], options.connTimeout)
// 		}
// 		if err == nil {
// 			options.logf(
// 				"[dial] secure=%t, skip_verify=%t, strategy=%s, ident=%d, server=%d -> %s",
// 				options.secure,
// 				options.skipVerify,
// 				options.openStrategy,
// 				ident,
// 				num,
// 				conn.RemoteAddr(),
// 			)
// 			if tcp, ok := conn.(*net.TCPConn); ok {
// 				err = tcp.SetNoDelay(options.noDelay) // Disable or enable the Nagle Algorithm for this tcp socket
// 				if err != nil {
// 					return nil, err
// 				}
// 			}
// 			return &connect{
// 				Conn:         conn,
// 				logf:         options.logf,
// 				ident:        ident,
// 				buffer:       bufio.NewReader(conn),
// 				readTimeout:  options.readTimeout,
// 				writeTimeout: options.writeTimeout,
// 			}, nil
// 		} else {
// 			options.logf(
// 				"[dial err] secure=%t, skip_verify=%t, strategy=%s, ident=%d, addr=%s\n%#v",
// 				options.secure,
// 				options.skipVerify,
// 				options.openStrategy,
// 				ident,
// 				options.hosts[num],
// 				err,
// 			)
// 		}
// 	}
// 	return nil, err
// }

// func (conn *connect) Read(b []byte) (int, error) {
// 	var (
// 		n      int
// 		err    error
// 		total  int
// 		dstLen = len(b)
// 	)
// 	if currentTime := time.Now(); conn.readTimeout != 0 && currentTime.Sub(conn.lastReadDeadlineTime) > (conn.readTimeout>>2) {
// 		conn.SetReadDeadline(time.Now().Add(conn.readTimeout))
// 		conn.lastReadDeadlineTime = currentTime
// 	}
// 	for total < dstLen {
// 		if n, err = conn.buffer.Read(b[total:]); err != nil {
// 			conn.logf("[connect] read error: %v", err)
// 			conn.Close()
// 			return n, driver.ErrBadConn
// 		}
// 		total += n
// 	}
// 	return total, nil
// }

// func (conn *connect) Write(b []byte) (int, error) {
// 	var (
// 		n      int
// 		err    error
// 		total  int
// 		srcLen = len(b)
// 	)
// 	if currentTime := time.Now(); conn.writeTimeout != 0 && currentTime.Sub(conn.lastWriteDeadlineTime) > (conn.writeTimeout>>2) {
// 		conn.SetWriteDeadline(time.Now().Add(conn.writeTimeout))
// 		conn.lastWriteDeadlineTime = currentTime
// 	}
// 	for total < srcLen {
// 		if n, err = conn.Conn.Write(b[total:]); err != nil {
// 			conn.logf("[connect] write error: %v", err)
// 			conn.Close()
// 			return n, driver.ErrBadConn
// 		}
// 		total += n
// 	}
// 	return n, nil
// }

// func (conn *connect) Close() error {
// 	if !conn.closed {
// 		conn.closed = true
// 		return conn.Conn.Close()
// 	}
// 	return nil
// }

// type clickhouoseRepository struct {
// 	client   *clickhouose.Client
// 	database string
// 	timeout  time.Duration
// }

// func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		return nil, err
// 	}
// 	return client, nil
// }

// func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (shortener.RedirectRepository, error) {
// 	repo := &mongoRepository{
// 		timeout:  time.Duration(mongoTimeout) * time.Second,
// 		database: mongoDB,
// 	}
// 	client, err := newMongoClient(mongoURL, mongoTimeout)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "repository.NewMongoRepo")
// 	}
// 	repo.client = client
// 	return repo, nil
// }

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/ClickHouse/clickhouse-go"
// )

// type clickhouseRepository struct {
// 	client   *sql.Open
// 	database string
// }

// var clickhouseDatabase string = "clickhouse"
// var clickhouseURL string = "tcp://127.0.0.1:9000?debug=true"
// var clickhouseTable string = "testable"

// func newClickhouseClient(clickhouseDatabase string, clickhouseURL string) (*sql.Open, error) {
// 	connect, err := sql.Open(clickhouseDatabase, clickhouseURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := connect.Ping(); err != nil {
// 		if exception, ok := err.(*clickhouse.Exception); ok {
// 			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
// 		} else {
// 			fmt.Println(err)
// 		}
// 		return
// 	}
// 	return connect, nil
// }

// func newClickhouseRepository (clickhouseDatabase string, clickhouseURL string, clickhouseEngine string, clickhouseTable string) (err error) {

// 	_, err = connect.Exec(`
// 		CREATE TABLE IF NOT EXISTS example (
// 			country_code FixedString(2),
// 			os_id        UInt8,
// 			browser_id   UInt8,
// 			categories   Array(Int16),
// 			action_day   Date,
// 			action_time  DateTime
// 		) engine=Memory
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return err
// }

// 	var (
// 		tx, _   = connect.Begin()
// 		stmt, _ = tx.Prepare("INSERT INTO example (country_code, os_id, browser_id, categories, action_day, action_time) VALUES (?, ?, ?, ?, ?, ?)")
// 	)
// 	defer stmt.Close()
// 	for i := 0; i < 100; i++ {
// 		if _, err := stmt.Exec(
// 			"RU",
// 			10+i,
// 			100+i,
// 			clickhouse.Array([]int16{1, 2, 3}),
// 			time.Now(),
// 			time.Now(),
// 		); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	if err := tx.Commit(); err != nil {
// 		log.Fatal(err)
// 	}

// 	rows, err := connect.Query("SELECT country_code, os_id, browser_id, categories, action_day, action_time FROM example")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var (
// 			country               string
// 			os, browser           uint8
// 			categories            []int16
// 			actionDay, actionTime time.Time
// 		)
// 		if err := rows.Scan(&country, &os, &browser, &categories, &actionDay, &actionTime); err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Printf("country: %s, os: %d, browser: %d, categories: %v, action_day: %s, action_time: %s", country, os, browser, categories, actionDay, actionTime)
// 	}

// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	if _, err := connect.Exec("DROP TABLE example"); err != nil {
// 		log.Fatal(err)
// 	}
