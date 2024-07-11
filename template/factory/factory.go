package factory

import (
	"fmt"

	factoryZmq4 "github.com/pebbe/zmq4"
)

type (
	file struct {
		name    string
		content string
	}
	message struct {
		topic string
		body  string
	}

	query struct {
		input  string
		output string
	}

	ntfs struct {
		files map[string]file
	}

	ext4 struct {
		files map[string]file
	}

	parquet struct {
		files map[string]file
	}

	zmq struct {
		data map[string]string
	}

	clickhouse struct {
		database map[string]string
	}

	mongoDB struct {
		database map[string]string
	}

	sqlite struct {
		database map[string]string
	}

	KrakenDepth struct {
		Error  []interface{} `json:"error"`
		Result struct {
		} `json:"result"`
	}

	// Zmq4pbSubscriber XPUBXSUB
	Zmq4pbSubscriber struct {
		GetSubscribersMap map[string][]string
		GetSubscriber     *factoryZmq4.Socket
		SetSubscribersMap map[string][]string
		SetSubscriber     error
	}

	//FileSystem interface
	FileSystem interface {
		CreateFile(string)
		FindFile(string) file
	}
	//Database interface
	Database interface {
		GetData(string) string
		PutData(string, string)
	}

	//MessageBus interface
	MessageBus interface {
		ReceiveMessage(string) string
		SendMessage(string, string)
	}

	//ParquetFile interface
	ParquetFile interface {
		CreateFile(string) error
		FindFile(string) file
	}

	//Factory func(string) interface{}
	Factory func(string) interface{}
)

// func (zmq zmq) ReceiveMessage(topic string) string {
// 	if _, ok := zmq.agent[topic]; !ok {
// 		return ""
// 	}
// 	fmt.Println("ZMQ")
// 	return
// }

func (clh clickhouse) GetData(query string) string {
	if _, ok := clh.database[query]; !ok {
		return ""
	}

	fmt.Println("Clickhouse")
	return clh.database[query]
}

func (mdb mongoDB) GetData(query string) string {
	if _, ok := mdb.database[query]; !ok {
		return ""
	}

	fmt.Println("MongoDB")
	return mdb.database[query]
}

func (sql sqlite) GetData(query string) string {
	if _, ok := sql.database[query]; !ok {
		return ""
	}

	fmt.Println("Sqlite")
	return sql.database[query]
}

// func (zmq *zmq) SendMessage(header string, body string) {
// 	// zmq.agent[message] = data
// 	zmq.agent[header] = body

// }

func (mdb mongoDB) PutData(query string, data string) {
	mdb.database[query] = data
}

func (clh clickhouse) PutData(query string, data string) {
	clh.database[query] = data
}

func (sql sqlite) PutData(query string, data string) {
	sql.database[query] = data
}

func (ntfs ntfs) CreateFile(path string) {
	file := file{content: "NTFS file", name: path}
	ntfs.files[path] = file
	fmt.Println("NTFS")
}

func (parquet parquet) CreateFile(path string) {
	file := file{content: "Parquet file", name: path}
	parquet.files[path] = file
	fmt.Println("Parquet")
}

func (ext ext4) CreateFile(path string) {
	file := file{content: "EXT4 file", name: path}
	ext.files[path] = file
	fmt.Println("EXT4")
}

func (ntfs ntfs) FindFile(path string) file {
	if _, ok := ntfs.files[path]; !ok {
		return file{}
	}

	return ntfs.files[path]
}

func (ext ext4) FindFile(path string) file {
	if _, ok := ext.files[path]; !ok {
		return file{}
	}

	return ext.files[path]
}

func (parquet parquet) FindFile(path string) file {
	if _, ok := parquet.files[path]; !ok {
		return file{}
	}

	return parquet.files[path]
}

//FilesystemFactory (env string) interface{}
func FilesystemFactory(env string) interface{} {
	switch env {
	case "production":
		return ntfs{
			files: make(map[string]file),
		}
	case "development":
		return ext4{
			files: make(map[string]file),
		}
	default:
		return nil
	}
}

// //MessageBusFactory (env string) interface{}
// func MessageBusFactory(env string) interface{} {
// 	switch env {
// 	case "production":
// 		return zmq{
// 			agent: make(map[string]string),
// 		}
// 	case "development":
// 		return zmq{
// 			agent: make(map[string]string),
// 		}
// 	default:
// 		return nil
// 	}
// }

// //DatabaseFactory (env string) interface{}
// func DatabaseFactory(env string) interface{} {
// 	switch env {
// 	case "production":
// 		return mongoDB{
// 			database: make(map[string]string),
// 		}
// 	case "development":
// 		return sqlite{
// 			database: make(map[string]string),
// 		}
// 	default:
// 		return nil
// 	}
// }

// //AbstractFactory (fact string) Factory
// func AbstractFactory(fact string) Factory {
// 	switch fact {
// 	case "database":
// 		return DatabaseFactory
// 	case "filesystem":
// 		return FilesystemFactory
// 	case "messagebus":
// 		return MessageBusFactory
// 	default:
// 		return nil
// 	}
// }
