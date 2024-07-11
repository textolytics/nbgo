package parquet

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "979902a52f0bc5447df723a3fb94c9b1-02d41003d7bd5d5acd1fa3e6b63a4b99", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
// 	instrs  []string
// )

// type NbOandaTick struct {
// 	instrs string
// 	tick   string
// }

// type I interface {
// 	M()
// }

// func main() {
// 	flag.Parse()
// 	if *token == "" {
// 		panic("An Oanda authorization token is required")
// 	}

// 	if *account == 0 {
// 		panic("An Oanda account is required")
// 	}

// 	client, err := oanda.NewFxPracticeClient(*token)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// List available account

// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(instruments)

// 	for i := range instruments {
// 		fmt.Println(i)
// 		instrs = append(instrs, i)
// 	}

// 	// Create and run a NewPriceServer server.
// 	priceServer, err := client.NewPriceServer(instrs...)
// 	if err != nil {
// 		panic(err)
// 	}

// 	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
// 		if err != nil {
// 			// fmt.Println("Received err:", err)
// 			panic(err)
// 		}
// fmt.Println("Received tick:", instrs, tick)
// fmt.Printf("Received instrs type : %T tick type : %T ", instrs, tick)
// tickParquet := oanda.PriceTick{
// 	Time:   tick.Time,
// 	Bid:    tick.Bid,
// 	Ask:    tick.Ask,
// 	Status: tick.Status,
// }
// fmt.Println("INST:%d TICK - %d ", instr, tickParquet)

// writeParquetTEST(&tickParquet)

// priceServer.Stop()
// 	})
// }

//-----------------------------------Parquet--writeNested----------------------

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name    string               `parquet:"name=name, type=UTF8"`
// 	Age     int32                `parquet:"name=age, type=INT32"`
// 	Weight  *int32               `parquet:"name=weight, type=INT32"`
// 	Classes *map[string][]*Class `parquet:"name=classes, type=MAP, keytype=UTF8"`
// }

// type Class struct {
// 	Name     string   `parquet:"name=name, type=UTF8"`
// 	Id       *int32   `parquet:"name=id, type=INT32"`
// 	Required []string `parquet:"name=required, type=LIST, valuetype=UTF8"`
// 	Ignored  string
// }

// func (c Class) String() string {
// 	id := "nil"
// 	if c.Id != nil {
// 		id = fmt.Sprintf("%d", *c.Id)
// 	}
// 	res := fmt.Sprintf("{Name:%s, Id:%v, Required:%s}", c.Name, id, fmt.Sprint(c.Required))
// 	return res
// }

// func (s Student) String() string {
// 	weight := "nil"
// 	if s.Weight != nil {
// 		weight = fmt.Sprintf("%d", *s.Weight)
// 	}

// 	cs := "{"
// 	for key, classes := range *s.Classes {
// 		s := string(key) + ":["
// 		for _, class := range classes {
// 			s += (*class).String() + ","
// 		}
// 		s += "]"
// 		cs += s
// 	}
// 	cs += "}"
// 	res := fmt.Sprintf("{Name:%s, Age:%d, Weight:%s, Classes:%s}", s.Name, s.Age, weight, cs)
// 	return res
// }

// func writeNested() {
// 	var err error
// 	math01ID := int32(1)
// 	math01 := Class{
// 		Name:     "Math1",
// 		Id:       &math01ID,
// 		Required: make([]string, 0),
// 	}

// 	math02ID := int32(2)
// 	math02 := Class{
// 		Name:     "Math2",
// 		Id:       &math02ID,
// 		Required: make([]string, 0),
// 	}
// 	math02.Required = append(math02.Required, "Math01")

// 	physics := Class{
// 		Name:     "Physics",
// 		Id:       nil,
// 		Required: make([]string, 0),
// 	}
// 	physics.Required = append(physics.Required, "Math01", "Math02")

// 	weight01 := int32(60)
// 	stu01Class := make(map[string][]*Class)
// 	stu01Class["Science1"] = make([]*Class, 0)
// 	stu01Class["Science1"] = append(stu01Class["Science"], &math01, &math02)
// 	stu01Class["Science2"] = make([]*Class, 0)
// 	stu01Class["Science2"] = append(stu01Class["Science"], &math01, &math02)
// 	stu01 := Student{
// 		Name:    "zxt",
// 		Age:     18,
// 		Weight:  &weight01,
// 		Classes: &stu01Class,
// 	}

// 	stu02Class := make(map[string][]*Class)
// 	stu02Class["Science"] = make([]*Class, 0)
// 	stu02Class["Science"] = append(stu02Class["Science"], &physics)
// 	stu02 := Student{
// 		Name:    "tong",
// 		Age:     29,
// 		Weight:  nil,
// 		Classes: &stu02Class,
// 	}

// 	stus := make([]Student, 0)
// 	stus = append(stus, stu01, stu02)

// 	//write nested
// 	fw, err := local.NewLocalFileWriter("nested.parquet")
// 	if err != nil {
// 		log.Println("Can't create file", err)
// 		return
// 	}
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	for _, stu := range stus {
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 			return
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 	}
// 	fw.Close()
// 	log.Println("Write Finished")

// 	//read nested
// 	fr, err := local.NewLocalFileReader("nested.parquet")
// 	if err != nil {
// 		log.Println("Can't open file", err)
// 		return
// 	}
// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num := int(pr.GetNumRows())
// 	for i := 0; i < num; i++ {
// 		stus := make([]Student, 1)
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}
// 	pr.ReadStop()
// 	fr.Close()
// }

// func main() {
// 	writeNested()
// }

//-----------------------------------Parquet--CSV-------------------------------------------

// package main
// import (
// 	"fmt"
// 	"log"
// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go/writer"
// )
// func main() {
// 	var err error
// 	md := []string{
// 		"name=Name, type=UTF8, encoding=PLAIN_DICTIONARY",
// 		"name=Age, type=INT32",
// 		"name=Id, type=INT64",
// 		"name=Weight, type=FLOAT",
// 		"name=Sex, type=BOOLEAN",
// 	}
// 	//write
// 	fw, err := local.NewLocalFileWriter("csv.parquet")
// 	if err != nil {
// 		log.Println("Can't open file", err)
// 		return
// 	}
// 	pw, err := writer.NewCSVWriter(md, fw, 4)
// 	if err != nil {
// 		log.Println("Can't create csv writer", err)
// 		return
// 	}
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		data := []string{
// 			fmt.Sprintf("%s_%d", "Student Name", i),
// 			fmt.Sprintf("%d", 20+i%5),
// 			fmt.Sprintf("%d", i),
// 			fmt.Sprintf("%f", 50.0+float32(i)*0.1),
// 			fmt.Sprintf("%t", i%2 == 0),
// 		}
// 		rec := make([]*string, len(data))
// 		for j := 0; j < len(data); j++ {
// 			rec[j] = &data[j]
// 		}
// 		if err = pw.WriteString(rec); err != nil {
// 			log.Println("WriteString error", err)
// 		}
// 		data2 := []interface{}{
// 			"Student Name",
// 			int32(20 + i%5),
// 			int64(i),
// 			float32(50.0 + float32(i)*0.1),
// 			i%2 == 0,
// 		}
// 		if err = pw.Write(data2); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()
// }

//-----------------------------------GOANDA---------------------------------------------

// package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
// 	instrs  []string
// )

// func main() {
// 	flag.Parse()
// 	if *token == "" {
// 		panic("An Oanda authorization token is required")
// 	}

// 	if *account == 0 {
// 		panic("An Oanda account is required")
// 	}

// 	client, err := oanda.NewFxPracticeClient(*token)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// List available account

// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(instruments)

// 	for i := range instruments {
// 		fmt.Println(i)
// 		instrs = append(instrs, i)
// 	}

// 	// Create and run a NewPriceServer server.
// 	priceServer, err := client.NewPriceServer(instrs...)
// 	if err != nil {
// 		panic(err)
// 	}

// 	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
// 		if err != nil {
// 			fmt.Println("Received err:", err)
// 			panic(err)
// 		}
// 		fmt.Println("Received tick:", instrs, tick)
// 		fmt.Printf("Received instrs type : %T tick type : %T ", instrs, tick.Bid)
// tickParquet := tick{
// 	Instrument: instrs,
// 	Time:       tick.Time,
// 	Bid:        tick.Bid,
// 	Ask:        tick.Ask,
// 	Status:     tick.Status,
// }
// writeParquetTEST(&tickParquet)

// priceServer.Stop()
// })
// }

// type parquetRecord struct {
// 	instrs []string
// 	tick struct
// }

// func writeParquetTEST(t *tickRecord) {
// 	var err error
// 	fw, err := local.NewLocalFileWriter("flat.parquet")
// 	if err != nil {
// 		log.Println("Can't create local file", err)
// 		return
// 	}

// 	//write
// 	pw, err := writer.NewParquetWriter(fw, new(tickRecord), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}

// 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// 	pw.CompressionType = parquet.CompressionCodec_SNAPPY
// 	num := 100
// 	for i := 0; i < num; i++ {
// 		stu :=tickRecord{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			Id:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 			Day:    int32(time.Now().Unix() / 3600 / 24),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.Write(t); err != nil {
// 		log.Println("Write error", err)
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 		return
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()

// 	///read
// 	fr, err := local.NewLocalFileReader("flat.parquet")
// 	if err != nil {
// 		log.Println("Can't open file")
// 		return
// 	}

// 	pr, err := reader.NewParquetReader(fr, new(tickRecord), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
// 	for i := 0; i < num/10; i++ {
// 		if i%2 == 0 {
// 			pr.SkipRows(10) //skip 10 rows
// 			continue
// 		}
// 		stus := make([]tickRecord, 10) //read 10 rows
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}

// 	pr.ReadStop()
// 	fr.Close()
// }

// Buy one unit of EUR/USD with a trailing stop of 10 pips.
// tradeInfo, err := client.NewTrade(oanda.Buy, 1, "eur_usd", oanda.TrailingStop(10.0))
// if err != nil {
// 	panic(err)
// }
// fmt.Println("tradeInfo:", tradeInfo)

//--------Create and run a NewPricePoller server.

// pricePoller, err := client.NewPricePoller(time.Now(), instrs...)

// if err != nil {
// 	panic(err)
// }
// pricesPoll, err := pricePoller.Poll()
// if err != nil {
// 	panic(err)
// }
// fmt.Println("pricesPoll:", pricesPoll)

// fmt.Printf("pricesPollFORMAT:%T", pricesPoll)

// 	// Close the previously opened trade.
// 	tradeCloseInfo, err := client.CloseTrade(tradeInfo.TradeId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("tradeCloseInfo:", tradeCloseInfo)
//-----------------------------------------------------------

// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	account := os.Getenv("OANDA_ACCOUNT")
// 	token := os.Getenv("OANDA_API_KEY")
// 	// currencies := os.Getenv("OANDA_CURRENCIES")

// 	c := client.New(account, token, "EUR_USD")

// 	c.Run(func(t *client.Tick) {

// 		fmt.Printf("%#v\n", t)
// 		fmt.Println(t)

// 	})
// }

// // type OandaHeartbitJSON struct {
// // 	Time time.Time `json:"time"`
// // 	Type string    `json:"type"`
// // }

// // type OandaTickJSON struct {
// // 	Asks []struct {
// // 		Liquidity int    `json:"askLiquidity"`
// // 		Price     string `json:"askPrice"`
// // 	} `json:"asks"`
// // 	Bids []struct {
// // 		Liquidity int    `json:"bidLiquidity"`
// // 		Price     string `json:"bidPrice"`
// // 	} `json:"bids"`
// // 	CloseoutAsk string    `json:"closeoutAsk"`
// // 	CloseoutBid string    `json:"closeoutBid"`
// // 	Instrument  string    `json:"instrument"`
// // 	Status      string    `json:"status"`
// // 	Time        string 	  `json:"time"`
// // 	Type        string    `json:"type"`
// // }
// type Tick struct {
// 	Asks        []Quote `json:"asks"`
// 	Bids        []Quote `json:"bids"`
// 	CloseoutAsk string  `json:"closeoutAsk"`
// 	CloseoutBid string  `json:"closeoutBid"`
// 	Instrument  string  `json:"instrument"`
// 	Status      string  `json:"status"`
// 	Time        string  `json:"time,omitempty"`
// 	Type        string  `json:"type"`

// 	// used to avoid parsing the Time multiple times
// 	parsedTime time.Time
// }

// func (t *Tick) IsJapanese() bool {
// 	return strings.Contains(t.Instrument, "JPY")
// }

// func (t *Tick) IsHeartbeat() bool {
// 	return "HEARTBEAT" == t.Type
// }

// func (t *Tick) IsTradeable() bool {
// 	return "tradeable" == t.Status
// }

// func (t *Tick) Symbol() string {
// 	return strings.Replace(t.Instrument, "_", "", 1)
// }

// func (t *Tick) parseTime() time.Time {
// 	if !t.parsedTime.IsZero() {
// 		return t.parsedTime
// 	}
// 	fmt.Println(t)
// 	parsedTime, err := time.Parse(time.RFC3339Nano, t.Time)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	t.parsedTime = parsedTime

// 	return t.parsedTime
// }
// func (t *Tick) UnixTimestamp() int64 {
// 	return t.parseTime().Unix()
// }

// func (t *Tick) Nanoseconds() int64 {
// 	return int64(t.parseTime().Nanosecond())
// }

// func (t *Tick) BestAsk() float64 {
// 	if 0 == len(t.Asks) {
// 		return 0.0
// 	}

// 	var best float64

// 	// best ask is the lowest
// 	for _, ask := range t.Asks {
// 		val := ask.PriceAsFloat()
// 		if val < best {
// 			best = val
// 		}
// 	}

// 	return best
// }

// func (t *Tick) BestBid() float64 {
// 	if 0 == len(t.Bids) {
// 		return 0.0
// 	}

// 	var best float64

// 	// best bid is the highest
// 	for _, bid := range t.Bids {
// 		val := bid.PriceAsFloat()
// 		if val > best {
// 			best = val
// 		}
// 	}

// 	return best
// }

// type Quote struct {
// 	Liquidity int64  `json:"liquidity"`
// 	Price     string `json:"price"`
// }

// func (q *Quote) PriceAsFloat() float64 {
// 	val, err := strconv.ParseFloat(q.Price, 64)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return val
// }

// type Client struct {
// 	account    string
// 	token      string
// 	currencies string
// }

// func New(account, token, currencies string) *Client {
// 	return &Client{
// 		account:    account,
// 		token:      token,
// 		currencies: currencies,
// 	}
// }

// func (c *Client) url() string {
// 	return fmt.Sprintf("api-fxpractice.oanda.com", c.account, c.currencies)
// }

// func (c *Client) Run(f func(*Tick)) {
// 	req, err := http.NewRequest("GET", c.url(), nil)
// 	if err != nil {
// 		log.Fatalln("http.NewRequest:", err)
// 		return
// 	}

// 	// set our bearer token
// 	req.Header.Set("Authorization", "Bearer "+c.token)

// 	// just use the DefaultClient, no need to be fancy here
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Fatalln("http.Get:", err)
// 		return
// 	}

// 	tick := &Tick{}

// 	reader := bufio.NewReader(resp.Body)
// 	for {
// 		line, err := reader.ReadBytes('\n')
// 		if err != nil {
// 			// technically, we should never get io.EOF here
// 			log.Fatalln("reader.ReadBytes:", err)
// 			return
// 		}

// 		if err := json.Unmarshal(line, tick); err != nil {
// 			log.Fatalln("json.Unmarshal:", err)
// 			return
// 		}

// 		// skip a few kinds of ticks here:
// 		//   - the heartbeat which is sent every 5 seconds
// 		//   - the "last prices" sent when initially connecting to the API
// 		if tick.IsTradeable() {
// 			f(tick)
// 		}
// 	}
// }

// func main() {
// 	// account := "3914094"
// 	// token := "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869"
// 	// currencies := os.Getenv("OANDA_CURRENCIES")
// 	// fmt.Printf(account)

// 	c := client.New("101-004-3748257-002", "2b557bdd4fa3dee56c8a159ece012a48-5f5a29d25cb2e7ea1aaeba98f4bbca40", "EURUSD")
// 	c.Run(func(t *client.Tick) {
// 		fmt.Println(t)
// 		// this function fires every time a tick is received
// 	})
// }

//-----------------

// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	fmt.Fprint(os.Stdout, "unix_timestamp,nanoseconds,symbol,best_bid,best_ask\n")

// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		tick := &client.Tick{}

// 		err := json.Unmarshal(scanner.Bytes(), tick)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		// don't assume heartbeats were removed
// 		if !tick.IsTradeable() {
// 			continue
// 		}

// 		var pipFormat string

// 		// output the correct number of pips
// 		if tick.IsJapanese() {
// 			pipFormat = "%.3f"
// 		} else {
// 			pipFormat = "%.5f"
// 		}

// 		format := "%d,%d,%s," + pipFormat + "," + pipFormat + "\n"

// 		fmt.Fprintf(os.Stdout, format,
// 			tick.UnixTimestamp(),
// 			tick.Nanoseconds(),
// 			tick.Symbol(),
// 			tick.BestBid(),
// 			tick.BestAsk(),
// 		)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

//---------------------------------------------------------------

// package main

// import (
// 	"fmt"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	// account := "3914094"
// 	// token := "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869"
// 	// currencies := os.Getenv("OANDA_CURRENCIES")
// 	// fmt.Printf(account)

// 	c := client.New("101-004-3748257-002", "2b557bdd4fa3dee56c8a159ece012a48-5f5a29d25cb2e7ea1aaeba98f4bbca40", "EURUSD")
// 	c.Run(func(t *client.Tick) {
// 		fmt.Println(t)
// 		// this function fires every time a tick is received
// 	})
// }

//---------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/awoldes/goanda"
// 	// "github.com/davecgh/go-spew/spew"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	key := os.Getenv("OANDA_API_KEY")
// 	accountID := os.Getenv("OANDA_ACCOUNT_ID")
// 	oanda := goanda.NewConnection(accountID, key, false)
// 	fmt.Printf("%T\n", oanda)
// 	ilist := getIstrumentsListStringCommaDelimiter(oanda)
// 	fmt.Println(ilist)
// 	// history := oanda.GetCandles("EUR_USD", "1", "M5")
// 	// quotes := oanda.GetInstrumentPrice(ilist)
// 	// fmt.Println(quotes)
// 	// spew.Dump(quotes)
// }

// func getIstrumentsListStringCommaDelimiter(oanda *goanda.OandaConnection) (data string) {
// 	a := oanda.GetAccounts()
// 	fmt.Println(a)
// 	for _, v := range a.Accounts {
// 		ai := oanda.GetAccount(v.ID)
// 		if ai.Account.Currency == "USD" {
// 			ua := ai.Account.ID
// 			fmt.Println(ua)
// 			il := oanda.GetAccountInstruments(ua)
// 			for _, v := range il.Instruments {
// 				data += v.Name + ","
// 			}
// 		}
// 	}
// 	return data[:(len(data) - 1)]
// }

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "979902a52f0bc5447df723a3fb94c9b1-02d41003d7bd5d5acd1fa3e6b63a4b99", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
// 	instrs  []string
// )

// type NbOandaTick struct {
// 	instrs string
// 	tick   string
// }

// type I interface {
// 	M()
// }

// func main() {
// 	flag.Parse()
// 	if *token == "" {
// 		panic("An Oanda authorization token is required")
// 	}

// 	if *account == 0 {
// 		panic("An Oanda account is required")
// 	}

// 	client, err := oanda.NewFxPracticeClient(*token)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// List available account

// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(instruments)

// 	for i := range instruments {
// 		fmt.Println(i)
// 		instrs = append(instrs, i)
// 	}

// 	// Create and run a NewPriceServer server.
// 	priceServer, err := client.NewPriceServer(instrs...)
// 	if err != nil {
// 		panic(err)
// 	}

// 	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
// 		if err != nil {
// 			// fmt.Println("Received err:", err)
// 			panic(err)
// 		}
// 		// fmt.Println("Received tick:", instrs, tick)
// 		// fmt.Printf("Received instrs type : %T tick type : %T ", instrs, tick)
// 		// tickParquet := oanda.PriceTick{
// 		// 	Time:   tick.Time,
// 		// 	Bid:    tick.Bid,
// 		// 	Ask:    tick.Ask,
// 		// 	Status: tick.Status,
// 		// }
// 		// fmt.Println("INST:%d TICK - %d ", instr, tickParquet)

// 		// writeParquetTEST(&tickParquet)

// 		// priceServer.Stop()
// 	})
// }

//-----------------------------------Parquet--writeNested----------------------

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name    string               `parquet:"name=name, type=UTF8"`
// 	Age     int32                `parquet:"name=age, type=INT32"`
// 	Weight  *int32               `parquet:"name=weight, type=INT32"`
// 	Classes *map[string][]*Class `parquet:"name=classes, type=MAP, keytype=UTF8"`
// }

// type Class struct {
// 	Name     string   `parquet:"name=name, type=UTF8"`
// 	Id       *int32   `parquet:"name=id, type=INT32"`
// 	Required []string `parquet:"name=required, type=LIST, valuetype=UTF8"`
// 	Ignored  string
// }

// func (c Class) String() string {
// 	id := "nil"
// 	if c.Id != nil {
// 		id = fmt.Sprintf("%d", *c.Id)
// 	}
// 	res := fmt.Sprintf("{Name:%s, Id:%v, Required:%s}", c.Name, id, fmt.Sprint(c.Required))
// 	return res
// }

// func (s Student) String() string {
// 	weight := "nil"
// 	if s.Weight != nil {
// 		weight = fmt.Sprintf("%d", *s.Weight)
// 	}

// 	cs := "{"
// 	for key, classes := range *s.Classes {
// 		s := string(key) + ":["
// 		for _, class := range classes {
// 			s += (*class).String() + ","
// 		}
// 		s += "]"
// 		cs += s
// 	}
// 	cs += "}"
// 	res := fmt.Sprintf("{Name:%s, Age:%d, Weight:%s, Classes:%s}", s.Name, s.Age, weight, cs)
// 	return res
// }

// func writeNested() {
// 	var err error
// 	math01ID := int32(1)
// 	math01 := Class{
// 		Name:     "Math1",
// 		Id:       &math01ID,
// 		Required: make([]string, 0),
// 	}

// 	math02ID := int32(2)
// 	math02 := Class{
// 		Name:     "Math2",
// 		Id:       &math02ID,
// 		Required: make([]string, 0),
// 	}
// 	math02.Required = append(math02.Required, "Math01")

// 	physics := Class{
// 		Name:     "Physics",
// 		Id:       nil,
// 		Required: make([]string, 0),
// 	}
// 	physics.Required = append(physics.Required, "Math01", "Math02")

// 	weight01 := int32(60)
// 	stu01Class := make(map[string][]*Class)
// 	stu01Class["Science1"] = make([]*Class, 0)
// 	stu01Class["Science1"] = append(stu01Class["Science"], &math01, &math02)
// 	stu01Class["Science2"] = make([]*Class, 0)
// 	stu01Class["Science2"] = append(stu01Class["Science"], &math01, &math02)
// 	stu01 := Student{
// 		Name:    "zxt",
// 		Age:     18,
// 		Weight:  &weight01,
// 		Classes: &stu01Class,
// 	}

// 	stu02Class := make(map[string][]*Class)
// 	stu02Class["Science"] = make([]*Class, 0)
// 	stu02Class["Science"] = append(stu02Class["Science"], &physics)
// 	stu02 := Student{
// 		Name:    "tong",
// 		Age:     29,
// 		Weight:  nil,
// 		Classes: &stu02Class,
// 	}

// 	stus := make([]Student, 0)
// 	stus = append(stus, stu01, stu02)

// 	//write nested
// 	fw, err := local.NewLocalFileWriter("nested.parquet")
// 	if err != nil {
// 		log.Println("Can't create file", err)
// 		return
// 	}
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	for _, stu := range stus {
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 			return
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 	}
// 	fw.Close()
// 	log.Println("Write Finished")

// 	//read nested
// 	fr, err := local.NewLocalFileReader("nested.parquet")
// 	if err != nil {
// 		log.Println("Can't open file", err)
// 		return
// 	}
// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num := int(pr.GetNumRows())
// 	for i := 0; i < num; i++ {
// 		stus := make([]Student, 1)
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}
// 	pr.ReadStop()
// 	fr.Close()
// }

// func main() {
// 	writeNested()
// }

//-----------------------------------Parquet--CSV-------------------------------------------

// package main
// import (
// 	"fmt"
// 	"log"
// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go/writer"
// )
// func main() {
// 	var err error
// 	md := []string{
// 		"name=Name, type=UTF8, encoding=PLAIN_DICTIONARY",
// 		"name=Age, type=INT32",
// 		"name=Id, type=INT64",
// 		"name=Weight, type=FLOAT",
// 		"name=Sex, type=BOOLEAN",
// 	}
// 	//write
// 	fw, err := local.NewLocalFileWriter("csv.parquet")
// 	if err != nil {
// 		log.Println("Can't open file", err)
// 		return
// 	}
// 	pw, err := writer.NewCSVWriter(md, fw, 4)
// 	if err != nil {
// 		log.Println("Can't create csv writer", err)
// 		return
// 	}
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		data := []string{
// 			fmt.Sprintf("%s_%d", "Student Name", i),
// 			fmt.Sprintf("%d", 20+i%5),
// 			fmt.Sprintf("%d", i),
// 			fmt.Sprintf("%f", 50.0+float32(i)*0.1),
// 			fmt.Sprintf("%t", i%2 == 0),
// 		}
// 		rec := make([]*string, len(data))
// 		for j := 0; j < len(data); j++ {
// 			rec[j] = &data[j]
// 		}
// 		if err = pw.WriteString(rec); err != nil {
// 			log.Println("WriteString error", err)
// 		}
// 		data2 := []interface{}{
// 			"Student Name",
// 			int32(20 + i%5),
// 			int64(i),
// 			float32(50.0 + float32(i)*0.1),
// 			i%2 == 0,
// 		}
// 		if err = pw.Write(data2); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()
// }

//-----------------------------------GOANDA---------------------------------------------

// package main

// import (
// 	"flag"
// 	"fmt"

// 	"github.com/santegoeds/oanda"
// )

// var (
// 	token   = flag.String("token", "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869", "Oanda authorization token.")
// 	account = flag.Int64("account", 3914094, "Oanda account.")
// 	instrs  []string
// )

// func main() {
// 	flag.Parse()
// 	if *token == "" {
// 		panic("An Oanda authorization token is required")
// 	}

// 	if *account == 0 {
// 		panic("An Oanda account is required")
// 	}

// 	client, err := oanda.NewFxPracticeClient(*token)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// List available account

// 	client.SelectAccount(oanda.Id(*account))

// 	// List available instruments
// 	instruments, err := client.Instruments(nil, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// fmt.Println(instruments)

// 	for i := range instruments {
// 		fmt.Println(i)
// 		instrs = append(instrs, i)
// 	}

// 	// Create and run a NewPriceServer server.
// 	priceServer, err := client.NewPriceServer(instrs...)
// 	if err != nil {
// 		panic(err)
// 	}

// 	priceServer.ConnectAndHandle(func(instrs string, tick oanda.PriceTick) {
// 		if err != nil {
// 			fmt.Println("Received err:", err)
// 			panic(err)
// 		}
// 		fmt.Println("Received tick:", instrs, tick)
// 		fmt.Printf("Received instrs type : %T tick type : %T ", instrs, tick.Bid)
// 		// tickParquet := tick{
// 		// 	Instrument: instrs,
// 		// 	Time:       tick.Time,
// 		// 	Bid:        tick.Bid,
// 		// 	Ask:        tick.Ask,
// 		// 	Status:     tick.Status,
// 		// }
// 		// writeParquetTEST(&tickParquet)

// 		// priceServer.Stop()
// 	})
// }

// type parquetRecord struct {
// 	instrs []string
// 	tick struct
// }

// func writeParquetTEST(t *tickRecord) {
// 	var err error
// 	fw, err := local.NewLocalFileWriter("flat.parquet")
// 	if err != nil {
// 		log.Println("Can't create local file", err)
// 		return
// 	}

// 	//write
// 	pw, err := writer.NewParquetWriter(fw, new(tickRecord), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}

// 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// 	pw.CompressionType = parquet.CompressionCodec_SNAPPY
// 	num := 100
// 	for i := 0; i < num; i++ {
// 		stu :=tickRecord{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			Id:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 			Day:    int32(time.Now().Unix() / 3600 / 24),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.Write(t); err != nil {
// 		log.Println("Write error", err)
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 		return
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()

// 	///read
// 	fr, err := local.NewLocalFileReader("flat.parquet")
// 	if err != nil {
// 		log.Println("Can't open file")
// 		return
// 	}

// 	pr, err := reader.NewParquetReader(fr, new(tickRecord), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
// 	for i := 0; i < num/10; i++ {
// 		if i%2 == 0 {
// 			pr.SkipRows(10) //skip 10 rows
// 			continue
// 		}
// 		stus := make([]tickRecord, 10) //read 10 rows
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}

// 	pr.ReadStop()
// 	fr.Close()
// }

// Buy one unit of EUR/USD with a trailing stop of 10 pips.
// tradeInfo, err := client.NewTrade(oanda.Buy, 1, "eur_usd", oanda.TrailingStop(10.0))
// if err != nil {
// 	panic(err)
// }
// fmt.Println("tradeInfo:", tradeInfo)

//--------Create and run a NewPricePoller server.

// pricePoller, err := client.NewPricePoller(time.Now(), instrs...)

// if err != nil {
// 	panic(err)
// }
// pricesPoll, err := pricePoller.Poll()
// if err != nil {
// 	panic(err)
// }
// fmt.Println("pricesPoll:", pricesPoll)

// fmt.Printf("pricesPollFORMAT:%T", pricesPoll)

// 	// Close the previously opened trade.
// 	tradeCloseInfo, err := client.CloseTrade(tradeInfo.TradeId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("tradeCloseInfo:", tradeCloseInfo)
//-----------------------------------------------------------

// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	account := os.Getenv("OANDA_ACCOUNT")
// 	token := os.Getenv("OANDA_API_KEY")
// 	// currencies := os.Getenv("OANDA_CURRENCIES")

// 	c := client.New(account, token, "EUR_USD")

// 	c.Run(func(t *client.Tick) {

// 		fmt.Printf("%#v\n", t)
// 		fmt.Println(t)

// 	})
// }

// // type OandaHeartbitJSON struct {
// // 	Time time.Time `json:"time"`
// // 	Type string    `json:"type"`
// // }

// // type OandaTickJSON struct {
// // 	Asks []struct {
// // 		Liquidity int    `json:"askLiquidity"`
// // 		Price     string `json:"askPrice"`
// // 	} `json:"asks"`
// // 	Bids []struct {
// // 		Liquidity int    `json:"bidLiquidity"`
// // 		Price     string `json:"bidPrice"`
// // 	} `json:"bids"`
// // 	CloseoutAsk string    `json:"closeoutAsk"`
// // 	CloseoutBid string    `json:"closeoutBid"`
// // 	Instrument  string    `json:"instrument"`
// // 	Status      string    `json:"status"`
// // 	Time        string 	  `json:"time"`
// // 	Type        string    `json:"type"`
// // }
// type Tick struct {
// 	Asks        []Quote `json:"asks"`
// 	Bids        []Quote `json:"bids"`
// 	CloseoutAsk string  `json:"closeoutAsk"`
// 	CloseoutBid string  `json:"closeoutBid"`
// 	Instrument  string  `json:"instrument"`
// 	Status      string  `json:"status"`
// 	Time        string  `json:"time,omitempty"`
// 	Type        string  `json:"type"`

// 	// used to avoid parsing the Time multiple times
// 	parsedTime time.Time
// }

// func (t *Tick) IsJapanese() bool {
// 	return strings.Contains(t.Instrument, "JPY")
// }

// func (t *Tick) IsHeartbeat() bool {
// 	return "HEARTBEAT" == t.Type
// }

// func (t *Tick) IsTradeable() bool {
// 	return "tradeable" == t.Status
// }

// func (t *Tick) Symbol() string {
// 	return strings.Replace(t.Instrument, "_", "", 1)
// }

// func (t *Tick) parseTime() time.Time {
// 	if !t.parsedTime.IsZero() {
// 		return t.parsedTime
// 	}
// 	fmt.Println(t)
// 	parsedTime, err := time.Parse(time.RFC3339Nano, t.Time)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	t.parsedTime = parsedTime

// 	return t.parsedTime
// }
// func (t *Tick) UnixTimestamp() int64 {
// 	return t.parseTime().Unix()
// }

// func (t *Tick) Nanoseconds() int64 {
// 	return int64(t.parseTime().Nanosecond())
// }

// func (t *Tick) BestAsk() float64 {
// 	if 0 == len(t.Asks) {
// 		return 0.0
// 	}

// 	var best float64

// 	// best ask is the lowest
// 	for _, ask := range t.Asks {
// 		val := ask.PriceAsFloat()
// 		if val < best {
// 			best = val
// 		}
// 	}

// 	return best
// }

// func (t *Tick) BestBid() float64 {
// 	if 0 == len(t.Bids) {
// 		return 0.0
// 	}

// 	var best float64

// 	// best bid is the highest
// 	for _, bid := range t.Bids {
// 		val := bid.PriceAsFloat()
// 		if val > best {
// 			best = val
// 		}
// 	}

// 	return best
// }

// type Quote struct {
// 	Liquidity int64  `json:"liquidity"`
// 	Price     string `json:"price"`
// }

// func (q *Quote) PriceAsFloat() float64 {
// 	val, err := strconv.ParseFloat(q.Price, 64)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return val
// }

// type Client struct {
// 	account    string
// 	token      string
// 	currencies string
// }

// func New(account, token, currencies string) *Client {
// 	return &Client{
// 		account:    account,
// 		token:      token,
// 		currencies: currencies,
// 	}
// }

// func (c *Client) url() string {
// 	return fmt.Sprintf("api-fxpractice.oanda.com", c.account, c.currencies)
// }

// func (c *Client) Run(f func(*Tick)) {
// 	req, err := http.NewRequest("GET", c.url(), nil)
// 	if err != nil {
// 		log.Fatalln("http.NewRequest:", err)
// 		return
// 	}

// 	// set our bearer token
// 	req.Header.Set("Authorization", "Bearer "+c.token)

// 	// just use the DefaultClient, no need to be fancy here
// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Fatalln("http.Get:", err)
// 		return
// 	}

// 	tick := &Tick{}

// 	reader := bufio.NewReader(resp.Body)
// 	for {
// 		line, err := reader.ReadBytes('\n')
// 		if err != nil {
// 			// technically, we should never get io.EOF here
// 			log.Fatalln("reader.ReadBytes:", err)
// 			return
// 		}

// 		if err := json.Unmarshal(line, tick); err != nil {
// 			log.Fatalln("json.Unmarshal:", err)
// 			return
// 		}

// 		// skip a few kinds of ticks here:
// 		//   - the heartbeat which is sent every 5 seconds
// 		//   - the "last prices" sent when initially connecting to the API
// 		if tick.IsTradeable() {
// 			f(tick)
// 		}
// 	}
// }

// func main() {
// 	// account := "3914094"
// 	// token := "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869"
// 	// currencies := os.Getenv("OANDA_CURRENCIES")
// 	// fmt.Printf(account)

// 	c := client.New("101-004-3748257-002", "2b557bdd4fa3dee56c8a159ece012a48-5f5a29d25cb2e7ea1aaeba98f4bbca40", "EURUSD")
// 	c.Run(func(t *client.Tick) {
// 		fmt.Println(t)
// 		// this function fires every time a tick is received
// 	})
// }

//-----------------

// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	fmt.Fprint(os.Stdout, "unix_timestamp,nanoseconds,symbol,best_bid,best_ask\n")

// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		tick := &client.Tick{}

// 		err := json.Unmarshal(scanner.Bytes(), tick)
// 		if err != nil {
// 			log.Fatalln(err)
// 		}

// 		// don't assume heartbeats were removed
// 		if !tick.IsTradeable() {
// 			continue
// 		}

// 		var pipFormat string

// 		// output the correct number of pips
// 		if tick.IsJapanese() {
// 			pipFormat = "%.3f"
// 		} else {
// 			pipFormat = "%.5f"
// 		}

// 		format := "%d,%d,%s," + pipFormat + "," + pipFormat + "\n"

// 		fmt.Fprintf(os.Stdout, format,
// 			tick.UnixTimestamp(),
// 			tick.Nanoseconds(),
// 			tick.Symbol(),
// 			tick.BestBid(),
// 			tick.BestAsk(),
// 		)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatalln(err)
// 	}
// }

//---------------------------------------------------------------

// package main

// import (
// 	"fmt"

// 	"github.com/dseevr/go-oanda-streaming-api/client"
// )

// func main() {
// 	// account := "3914094"
// 	// token := "5b2e1521432ad31ef69270b682394010-4df302be03bbefb18ad70e457f3db869"
// 	// currencies := os.Getenv("OANDA_CURRENCIES")
// 	// fmt.Printf(account)

// 	c := client.New("101-004-3748257-002", "2b557bdd4fa3dee56c8a159ece012a48-5f5a29d25cb2e7ea1aaeba98f4bbca40", "EURUSD")
// 	c.Run(func(t *client.Tick) {
// 		fmt.Println(t)
// 		// this function fires every time a tick is received
// 	})
// }

//---------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/awoldes/goanda"
// 	// "github.com/davecgh/go-spew/spew"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	key := os.Getenv("OANDA_API_KEY")
// 	accountID := os.Getenv("OANDA_ACCOUNT_ID")
// 	oanda := goanda.NewConnection(accountID, key, false)
// 	fmt.Printf("%T\n", oanda)
// 	ilist := getIstrumentsListStringCommaDelimiter(oanda)
// 	fmt.Println(ilist)
// 	// history := oanda.GetCandles("EUR_USD", "1", "M5")
// 	// quotes := oanda.GetInstrumentPrice(ilist)
// 	// fmt.Println(quotes)
// 	// spew.Dump(quotes)
// }

// func getIstrumentsListStringCommaDelimiter(oanda *goanda.OandaConnection) (data string) {
// 	a := oanda.GetAccounts()
// 	fmt.Println(a)
// 	for _, v := range a.Accounts {
// 		ai := oanda.GetAccount(v.ID)
// 		if ai.Account.Currency == "USD" {
// 			ua := ai.Account.ID
// 			fmt.Println(ua)
// 			il := oanda.GetAccountInstruments(ua)
// 			for _, v := range il.Instruments {
// 				data += v.Name + ","
// 			}
// 		}
// 	}
// 	return data[:(len(data) - 1)]
// }
