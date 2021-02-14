package hdfs

// import (
// 	"log"

// 	"github.com/colinmarc/hdfs"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
	// "github.com/xitongsys/parquet-go-source/hdfs"
	// "github.com/xitongsys/parquet-go/reader"
	// "github.com/xitongsys/parquet-go/writer"
// )

// //HdfsRepository "github.com/colinmarc/hdfs"
// type HdfsRepository struct {
// 	client   *hdfs.Client
// 	database string
// 	timeout  time.Duration
// }

// Parquete struct
// type Parquete struct {
// }

// // Student struct
// type Student struct {
// 	Name   string  `parquet:"name=name, type=UTF8"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	Id     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// }

// func (parquet *parquet) Write() {
// 	// client, _ := hdfs.New("namenode:8020")
// 	var err error
// 	//write
// 	fw, err := hdfs.NewHdfsFileWriter([]string{"spark.nb.lan:9000"}, "hadoop", "/old.parquet")
// 	if err != nil {
// 		log.Println("Can't create hdfs file", err)
// 		return
// 	}

// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	num := 100
// 	for i := 0; i < num; i++ {
// 		stu := Student{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			Id:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop err", err)
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()

// 	///read
// 	fr, err := hdfs.NewHdfsFileReader([]string{"spark.nb.lan:9000"}, "hadoop", "/old.parquet")
// 	if err != nil {
// 		log.Println("Can't open hdfs file", err)
// 		return
// 	}
// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
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

// func parquetAppend() {

// }

// func NewHDFSClient() () {

// 	client, _ := hdfs.New("hdfs://spark.nb.lan:9000")

// }

// 	file, _ := client.Open("/mobydick.txt")

// 	buf := make([]byte, 59)
// 	file.ReadAt(buf, 48847)

// 	fmt.Println(string(buf))

// Bool              bool    `parquet:"name=bool, type=BOOLEAN"`
// Int32             int32   `parquet:"name=int32, type=INT32"`
// Int64             int64   `parquet:"name=int64, type=INT64"`
// Int96             string  `parquet:"name=int96, type=INT96"`
// Float             float32 `parquet:"name=float, type=FLOAT"`
// Double            float64 `parquet:"name=double, type=DOUBLE"`
// ByteArray         string  `parquet:"name=bytearray, type=BYTE_ARRAY"`
// FixedLenByteArray string  `parquet:"name=FixedLenByteArray, type=FIXED_LEN_BYTE_ARRAY, length=10"`

// Utf8             string `parquet:"name=utf8, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
// Int_8            int32   `parquet:"name=int_8, type=INT32, convertedtype=INT32, convertedtype=INT_8"`
// Int_16           int32  `parquet:"name=int_16, type=INT32, convertedtype=INT_16"`
// Int_32           int32  `parquet:"name=int_32, type=INT32, convertedtype=INT_32"`
// Int_64           int64  `parquet:"name=int_64, type=INT64, convertedtype=INT_64"`
// Uint_8           int32  `parquet:"name=uint_8, type=INT32, convertedtype=UINT_8"`
// Uint_16          int32 `parquet:"name=uint_16, type=INT32, convertedtype=UINT_16"`
// Uint_32          int32 `parquet:"name=uint_32, type=INT32, convertedtype=UINT_32"`
// Uint_64          int64 `parquet:"name=uint_64, type=INT64, convertedtype=UINT_64"`
// Date             int32  `parquet:"name=date, type=INT32, convertedtype=DATE"`
// Date2            int32  `parquet:"name=date2, type=INT32, convertedtype=DATE, logicaltype=DATE"`
// TimeMillis       int32  `parquet:"name=timemillis, type=INT32, convertedtype=TIME_MILLIS"`
// TimeMillis2      int32  `parquet:"name=timemillis2, type=INT32, logicaltype=TIME, logicaltype.isadjustedtoutc=true, logicaltype.unit=MILLIS"`
// TimeMicros       int64  `parquet:"name=timemicros, type=INT64, convertedtype=TIME_MICROS"`
// TimeMicros2      int64  `parquet:"name=timemicros2, type=INT64, logicaltype=TIME, logicaltype.isadjustedtoutc=false, logicaltype.unit=MICROS"`
// TimestampMillis  int64  `parquet:"name=timestampmillis, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
// TimestampMillis2 int64  `parquet:"name=timestampmillis2, type=INT64, logicaltype=TIMESTAMP, logicaltype.isadjustedtoutc=true, logicaltype.unit=MILLIS"`
// TimestampMicros  int64  `parquet:"name=timestampmicros, type=INT64, convertedtype=TIMESTAMP_MICROS"`
// TimestampMicros2 int64  `parquet:"name=timestampmicros2, type=INT64, logicaltype=TIMESTAMP, logicaltype.isadjustedtoutc=false, logicaltype.unit=MICROS"`
// Interval         string `parquet:"name=interval, type=BYTE_ARRAY, convertedtype=INTERVAL"`

// Decimal1 int32  `parquet:"name=decimal1, type=INT32, convertedtype=DECIMAL, scale=2, precision=9"`
// Decimal2 int64  `parquet:"name=decimal2, type=INT64, convertedtype=DECIMAL, scale=2, precision=18"`
// Decimal3 string `parquet:"name=decimal3, type=FIXED_LEN_BYTE_ARRAY, convertedtype=DECIMAL, scale=2, precision=10, length=12"`
// Decimal4 string `parquet:"name=decimal4, type=BYTE_ARRAY, convertedtype=DECIMAL, scale=2, precision=20"`

// Decimal5 int32 `parquet:"name=decimal5, type=INT32, logicaltype=DECIMAL, logicaltype.precision=10, logicaltype.scale=2"`

// Map      map[string]int32 `parquet:"name=map, type=MAP, convertedtype=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=INT32"`
// List     []string         `parquet:"name=list, type=MAP, convertedtype=LIST, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
// Repeated []int32          `parquet:"name=repeated, type=INT32, repetitiontype=REPEATED"`

//------------------------

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"time"
// "github.com/colinmarc/hdfs"
// 	"github.com/xitongsys/parquet-go-source/writerfile"
// 	"github.com/xitongsys/parquet-go/parquet"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name   string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	Id     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// 	Day    int32   `parquet:"name=day, type=DATE"`
// }

// func main() {
// 	var err error
// 	buf := new(bytes.Buffer)
// 	fw := writerfile.NewWriterFile(buf)
// 	fmt.Printf("%T\n", fw)
// 	//write
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	// writerfile.WriterFile("flat.parquet")
// 	file, _ := hdfs.3.-*0("/mobydick.txt")

// 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// 	pw.CompressionType = parquet.CompressionCodec_SNAPPY
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		stu := Student{
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
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 		return
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()
// 	log.Println(buf)

// }

// // //-----------WRITE-hdfs--------------------

// package main

// import (
// 	"log"

// 	// "github.com/colinmarc/hdfs"
// 	"github.com/xitongsys/parquet-go-source/hdfs"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name   string  `parquet:"name=name, type=UTF8"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	Id     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// }

// func main() {
// 	// client, _ := hdfs.New("namenode:8020")
// 	var err error
// 	//write
// 	fw, err := hdfs.NewHdfsFileWriter([]string{"spark.nb.lan:9000"}, "hadoop", "/old.parquet")
// 	if err != nil {
// 		log.Println("Can't create hdfs file", err)
// 		return
// 	}

// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	num := 100
// 	for i := 0; i < num; i++ {
// 		stu := Student{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			Id:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop err", err)
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()

// 	///read
// 	fr, err := hdfs.NewHdfsFileReader([]string{"spark.nb.lan:9000"}, "hadoop", "/old.parquet")
// 	if err != nil {
// 		log.Println("Can't open hdfs file", err)
// 		return
// 	}
// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
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

// func parquetAppend() {

// }

//-----------GSS=WRINTE---------------------
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/xitongsys/parquet-go-source/gcs"
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
// 	ctx := context.Background()
// 	projectId := "Please change this to your own gCloud project Id"
// 	bucketName := "Your bucket name"
// 	fileName := "gcs_example/csv.parquet"

// 	//write
// 	fw, err := gcs.NewGcsFileWriter(ctx, projectId, bucketName, fileName)
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
// 			[]byte("Student Name"),
// 			int32(20 + i%5),
// 			int64(i),
// 			float32(50.0 + float32(i)*0.1),
// 			bool(i%2 == 0),
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

// //--------------HDFS-FLAT===================
// package main

// import (
// 	"log"

// 	"github.com/xitongsys/parquet-go-source/hdfs"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name   string  `parquet:"name=name, type=UTF8"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	Id     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// }

// func main() {
// 	var err error
// 	//write
// 	fw, err := hdfs.NewHdfsFileWriter([]string{"localhost:9000"}, "root", "/flat.parquet")
// 	if err != nil {
// 		log.Println("Can't create hdfs file", err)
// 		return
// 	}
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}

// 	num := 10
// 	for i := 0; i < num; i++ {
// 		stu := Student{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			Id:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop err", err)
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()

// 	///read
// 	fr, err := hdfs.NewHdfsFileReader([]string{"localhost:9000"}, "", "/flat.parquet")
// 	if err != nil {
// 		log.Println("Can't open hdfs file", err)
// 		return
// 	}
// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
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

// //---------------S3--------------------------==============

// package main

// import (
// 	"context"
// 	"log"

// 	"github.com/xitongsys/parquet-go-source/s3"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type student struct {
// 	Name   string  `parquet:"name=name, type=UTF8"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	ID     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// }

// // s3Example provides a sample write and read using the S3 Parquet File
// func s3Example() {
// 	ctx := context.Background()
// 	bucket := "my-bucket"
// 	key := "test/foobar.parquet"
// 	num := 100

// 	// create new S3 file writer
// 	fw, err := s3.NewS3FileWriter(ctx, bucket, key, nil)
// 	if err != nil {
// 		log.Println("Can't open file", err)
// 		return
// 	}
// 	// create new parquet file writer
// 	pw, err := writer.NewParquetWriter(fw, new(student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	// write 100 student records to the parquet file
// 	for i := 0; i < num; i++ {
// 		stu := student{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			ID:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 		}
// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
// 	}
// 	// write parquet file footer
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop err", err)
// 	}

// 	err = fw.Close()
// 	if err != nil {
// 		log.Println("Error closing S3 file writer")
// 	}
// 	log.Println("Write Finished")

// 	// read the written parquet file
// 	// create new S3 file reader
// 	fr, err := s3.NewS3FileReader(ctx, bucket, key)
// 	if err != nil {
// 		log.Println("Can't open file")
// 		return
// 	}

// 	// create new parquet file reader
// 	pr, err := reader.NewParquetReader(fr, new(student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}

// 	// read the student rows and print
// 	num = int(pr.GetNumRows())
// 	for i := 0; i < num/10; i++ {
// 		if i%2 == 0 {
// 			pr.SkipRows(10) //skip 10 rows
// 			continue
// 		}
// 		stus := make([]student, 10) //read 10 rows
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}

// 	// close the parquet file
// 	pr.ReadStop()
// 	err = fr.Close()
// 	if err != nil {
// 		log.Println("Error closing S3 file reader")
// 	}
// 	log.Println("Read Finished")
// }

//---------MEMFS-----------------------
// package main

// import (
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go-source/mem"
// 	"github.com/xitongsys/parquet-go/parquet"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name   string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
// 	Age    int32   `parquet:"name=age, type=INT32"`
// 	Id     int64   `parquet:"name=id, type=INT64"`
// 	Weight float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex    bool    `parquet:"name=sex, type=BOOLEAN"`
// 	Day    int32   `parquet:"name=day, type=DATE"`
// }

// func main() {
// 	// create in-memory ParquetFile with Closer Function
// 	// NOTE: closer function can be nil, no action will be
// 	// run when the writer is closed.
// 	fw, err := mem.NewMemFileWriter("flat.parquet.snappy", func(name string, r io.Reader) error {
// 		dat, err := ioutil.ReadAll(r)
// 		if err != nil {
// 			log.Printf("error reading data: %v", err)
// 			os.Exit(1)
// 		}

// 		// write file to disk
// 		if err := ioutil.WriteFile(name, dat, 0644); err != nil {
// 			log.Printf("error writing result file: %v", err)
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		log.Println("Can't create local file", err)
// 		return
// 	}
// 	//write
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}
// 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// 	pw.CompressionType = parquet.CompressionCodec_SNAPPY
// 	num := 10
// 	for i := 0; i < num; i++ {
// 		stu := Student{
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
// 	if err = pw.WriteStop(); err != nil {
// 		log.Println("WriteStop error", err)
// 		return
// 	}
// 	log.Println("Write Finished")
// 	fw.Close()
// 	// os.Exit(1)

// 	///read
// 	fr, err := local.NewLocalFileReader("flat.parquet.snappy")
// 	if err != nil {
// 		log.Println("Can't open file")
// 		return
// 	}

// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet reader", err)
// 		return
// 	}
// 	num = int(pr.GetNumRows())
// 	for i := 0; i < num; i++ {
// 		stus := make([]Student, 1)
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}
// 	pr.ReadStop()
// 	fr.Close()

// 	// NOTE: you can access the underlying MemFs using ParquetFile.GetMemFileFs()
// 	// EXAMPLE: this will delete the file we created from the in-memory file system
// 	if err := mem.GetMemFileFs().Remove("flat.parquet.snappy"); err != nil {
// 		log.Printf("error removing file from memfs: %v", err)
// 		os.Exit(1)
// 	}

// }
// }

///////////////////////////////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/xitongsys/parquet-go-source/local"
// 	"github.com/xitongsys/parquet-go/parquet"
// 	"github.com/xitongsys/parquet-go/reader"
// 	"github.com/xitongsys/parquet-go/writer"
// )

// type Student struct {
// 	Name    string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
// 	Age     int32   `parquet:"name=age, type=INT32, encoding=DELTA_BINARY_PACKED"`
// 	ID      int64   `parquet:"name=id, type=INT64, encoding=DELTA_BINARY_PACKED "`
// 	Weight  float32 `parquet:"name=weight, type=FLOAT"`
// 	Sex     bool    `parquet:"name=sex, type=BOOLEAN"`
// 	Day     int32   `parquet:"name=day, type=DATE, encoding=DELTA_BINARY_PACKED"`
// 	Ignored int32   //without parquet tag and won't write
// }

// type oandaTick struct {
// 	Ticker    string  `parquet:"name=ticker, type=UTF8, encoding=PLAIN_DICTIONARY"`
// 	Timestamp int64   `parquet:"name=timestamp, type=TIMESTAMP_MICROS, encoding=DELTA_BINARY_PACKED"`
// 	Status    string  `parquet:"name=status, type=UTF8"`
// 	Bid       float32 `parquet:"name=bid, type=FLOAT"`
// 	Ask       float32 `parquet:"name=ask, type=FLOAT"`
// }

// type parquetFile interface {
// 	io.Seeker
// 	io.Reader
// 	io.Writer
// 	io.Closer
// 	Open(name string) (parquetFile, error)
// 	Create(name string) (parquetFile, error)
// }

// // func (self *PqFile) CreateOrAppend(name string) (ParquetFile, error) {
// // 	//file, err := os.Create(name)
// // 	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// // 	myFile := new(PqFile)
// // 	myFile.File = file
// // 	return myFile, err
// // }

// func main() {
// 	var err error
// 	fw, err := local.NewLocalFileWriter("flat.parquet")
// 	if err != nil {
// 		log.Println("Can't create local file", err)
// 		return
// 	}

// 	f, err := os.OpenFile("flat.parquet", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("%T\n", *f)
// 	fmt.Printf("%T\n", fw)
// 	fmt.Println(*f)

// 	//write
// 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// 	if err != nil {
// 		log.Println("Can't create parquet writer", err)
// 		return
// 	}

// 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// 	pw.CompressionType = parquet.CompressionCodec_GZIP
// 	num := 100
// 	for i := 0; i < num; i++ {
// 		stu := Student{
// 			Name:   "StudentName",
// 			Age:    int32(20 + i%5),
// 			ID:     int64(i),
// 			Weight: float32(50.0 + float32(i)*0.1),
// 			Sex:    bool(i%2 == 0),
// 			Day:    int32(time.Now().Unix() / 3600 / 24),
// 		}

// 		if err = pw.Write(stu); err != nil {
// 			log.Println("Write error", err)
// 		}
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

// 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
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
// 		stus := make([]Student, 10) //read 10 rows
// 		if err = pr.Read(&stus); err != nil {
// 			log.Println("Read error", err)
// 		}
// 		log.Println(stus)
// 	}

// 	pr.ReadStop()
// 	fr.Close()

// }

////---------------------TEST---------------------------------
// // package main

// // import (
// // 	"log"
// // 	"time"

// // 	"github.com/xitongsys/parquet-go-source/local"
// // 	"github.com/xitongsys/parquet-go/parquet"
// // 	"github.com/xitongsys/parquet-go/reader"
// // 	"github.com/xitongsys/parquet-go/writer"
// // )

// // type Student struct {
// // 	Name    string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
// // 	Age     int32   `parquet:"name=age, type=INT32, encoding=DELTA_BINARY_PACKED"`
// // 	ID      int64   `parquet:"name=id, type=INT64, encoding=DELTA_BINARY_PACKED "`
// // 	Weight  float32 `parquet:"name=weight, type=FLOAT"`
// // 	Sex     bool    `parquet:"name=sex, type=BOOLEAN"`
// // 	Day     int32   `parquet:"name=day, type=DATE, encoding=DELTA_BINARY_PACKED"`
// // 	Ignored int32   //without parquet tag and won't write
// // }

// // type oandaTick struct {
// // 	Ticker    string  `parquet:"name=ticker, type=UTF8, encoding=PLAIN_DICTIONARY"`
// // 	Timestamp int64   `parquet:"name=timestamp, type=TIMESTAMP_MICROS, encoding=DELTA_BINARY_PACKED"`
// // 	Status    string  `parquet:"name=status, type=UTF8"`
// // 	Bid       float32 `parquet:"name=bid, type=FLOAT"`
// // 	Ask       float32 `parquet:"name=ask, type=FLOAT"`
// // }

// // func main() {
// // 	var err error
// // 	// fw, err := local.NewLocalFileWriter("flat.parquet")
// // 	// if err != nil {
// // 	// 	log.Println("Can't create local file", err)
// // 	// 	return
// // 	// }

// // 	fr, err := local.NewLocalFileReader("flat.parquet")
// // 	if err != nil {
// // 		log.Println("Can't open file")
// // 		return
// // 	}

// // 	//write
// // 	pw, err := writer.NewParquetWriter(fr, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet writer", err)
// // 		return
// // 	}

// // 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// // 	pw.CompressionType = parquet.CompressionCodec_GZIP
// // 	num := 10000000
// // 	for i := 0; i < num; i++ {
// // 		stu := Student{
// // 			Name:   "StudentName",
// // 			Age:    int32(20 + i%5),
// // 			ID:     int64(i),
// // 			Weight: float32(50.0 + float32(i)*0.1),
// // 			Sex:    bool(i%2 == 0),
// // 			Day:    int32(time.Now().Unix() / 3600 / 24),
// // 		}
// // 		if err = pw.Write(stu); err != nil {
// // 			log.Println("Write error", err)
// // 		}
// // 	}

// // 	if err = pw.WriteStop(); err != nil {
// // 		log.Println("WriteStop error", err)
// // 		return
// // 	}
// // 	log.Println("Write Finished")
// // 	fr.Close()

// // 	///read
// // 	fv, err := local.NewLocalFileReader("flat.parquet")
// // 	if err != nil {
// // 		log.Println("Can't open file")
// // 		return
// // 	}

// // 	pr, err := reader.NewParquetReader(fv, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet reader", err)
// // 		return
// // 	}
// // 	num = int(pr.GetNumRows())
// // 	for i := 0; i < num/10; i++ {
// // 		if i%2 == 0 {
// // 			pr.SkipRows(10) //skip 10 rows
// // 			continue
// // 		}
// // 		stus := make([]Student, 10) //read 10 rows
// // 		if err = pr.Read(&stus); err != nil {
// // 			log.Println("Read error", err)
// // 		}
// // 		log.Println(stus)
// // 	}

// // 	pr.ReadStop()
// // 	fv.Close()

// // }

// // //-------------------NESTED------------------------
// // package main

// // import (
// // 	"fmt"
// // 	"log"

// // 	"github.com/xitongsys/parquet-go-source/local"
// // 	"github.com/xitongsys/parquet-go/reader"
// // 	"github.com/xitongsys/parquet-go/writer"
// // )

// // type oandaTick struct {
// // 	Ticker    string  `parquet:"name=ticker, type=UTF8, encoding=PLAIN_DICTIONARY"`
// // 	Timestamp int64   `parquet:"name=timestamp, type=TIMESTAMP_MICROS, encoding=DELTA_BINARY_PACKED"`
// // 	Status    string  `parquet:"name=status, type=UTF8"`
// // 	Bid       float32 `parquet:"name=bid, type=FLOAT"`
// // 	Ask       float32 `parquet:"name=ask, type=FLOAT"`
// // }

// // type Student struct {
// // 	Name    string               `parquet:"name=name, type=UTF8"`
// // 	Age     int32                `parquet:"name=age, type=INT32"`
// // 	Weight  *int32               `parquet:"name=weight, type=INT32"`
// // 	Classes *map[string][]*Class `parquet:"name=classes, type=MAP, keytype=UTF8"`
// // }

// // type Class struct {
// // 	Name     string   `parquet:"name=name, type=UTF8"`
// // 	Id       *int32   `parquet:"name=id, type=INT32"`
// // 	Required []string `parquet:"name=required, type=LIST, valuetype=UTF8"`
// // 	Ignored  string
// // }

// // func (c Class) String() string {
// // 	id := "nil"
// // 	if c.Id != nil {
// // 		id = fmt.Sprintf("%d", *c.Id)
// // 	}
// // 	res := fmt.Sprintf("{Name:%s, Id:%v, Required:%s}", c.Name, id, fmt.Sprint(c.Required))
// // 	return res
// // }

// // func (s Student) String() string {
// // 	weight := "nil"
// // 	if s.Weight != nil {
// // 		weight = fmt.Sprintf("%d", *s.Weight)
// // 	}

// // 	cs := "{"
// // 	for key, classes := range *s.Classes {
// // 		s := string(key) + ":["
// // 		for _, class := range classes {
// // 			s += (*class).String() + ","
// // 		}
// // 		s += "]"
// // 		cs += s
// // 	}
// // 	cs += "}"
// // 	res := fmt.Sprintf("{Name:%s, Age:%d, Weight:%s, Classes:%s}", s.Name, s.Age, weight, cs)
// // 	return res
// // }

// // func writeNested() {
// // 	var err error
// // 	math01ID := int32(1)
// // 	math01 := Class{
// // 		Name:     "Math1",
// // 		Id:       &math01ID,
// // 		Required: make([]string, 0),
// // 	}

// // 	math02ID := int32(2)
// // 	math02 := Class{
// // 		Name:     "Math2",
// // 		Id:       &math02ID,
// // 		Required: make([]string, 0),
// // 	}
// // 	math02.Required = append(math02.Required, "Math01")

// // 	physics := Class{
// // 		Name:     "Physics",
// // 		Id:       nil,
// // 		Required: make([]string, 0),
// // 	}
// // 	physics.Required = append(physics.Required, "Math01", "Math02")

// // 	weight01 := int32(60)
// // 	stu01Class := make(map[string][]*Class)
// // 	stu01Class["Science1"] = make([]*Class, 0)
// // 	stu01Class["Science1"] = append(stu01Class["Science"], &math01, &math02)
// // 	stu01Class["Science2"] = make([]*Class, 0)
// // 	stu01Class["Science2"] = append(stu01Class["Science"], &math01, &math02)
// // 	stu01 := Student{
// // 		Name:    "zxt",
// // 		Age:     18,
// // 		Weight:  &weight01,
// // 		Classes: &stu01Class,
// // 	}

// // 	stu02Class := make(map[string][]*Class)
// // 	stu02Class["Science"] = make([]*Class, 0)
// // 	stu02Class["Science"] = append(stu02Class["Science"], &physics)
// // 	stu02 := Student{
// // 		Name:    "tong",
// // 		Age:     29,
// // 		Weight:  nil,
// // 		Classes: &stu02Class,
// // 	}

// // 	stus := make([]Student, 0)
// // 	stus = append(stus, stu01, stu02)

// // 	//write nested
// // 	fw, err := local.NewLocalFileWriter("nested.parquet")
// // 	if err != nil {
// // 		log.Println("Can't create file", err)
// // 		return
// // 	}
// // 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet writer", err)
// // 		return
// // 	}
// // 	for _, stu := range stus {
// // 		if err = pw.Write(stu); err != nil {
// // 			log.Println("Write error", err)
// // 			return
// // 		}
// // 	}
// // 	if err = pw.WriteStop(); err != nil {
// // 		log.Println("WriteStop error", err)
// // 	}
// // 	fw.Close()
// // 	log.Println("Write Finished")

// // 	//read nested
// // 	fr, err := local.NewLocalFileReader("nested.parquet")
// // 	if err != nil {
// // 		log.Println("Can't open file", err)
// // 		return
// // 	}
// // 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet reader", err)
// // 		return
// // 	}
// // 	num := int(pr.GetNumRows())
// // 	for i := 0; i < num; i++ {
// // 		stus := make([]Student, 1)
// // 		if err = pr.Read(&stus); err != nil {
// // 			log.Println("Read error", err)
// // 		}
// // 		log.Println(stus)
// // 	}
// // 	pr.ReadStop()
// // 	fr.Close()
// // }

// // func main() {
// // 	writeNested()
// // }

// //----------------------------FLAT----------------------------------

// // package main

// // import (
// // 	"log"
// // 	"time"

// // 	"github.com/xitongsys/parquet-go-source/local"
// // 	"github.com/xitongsys/parquet-go/parquet"
// // 	"github.com/xitongsys/parquet-go/reader"
// // 	"github.com/xitongsys/parquet-go/writer"
// // )

// // type Student struct {
// // 	Name    string  `parquet:"name=name, type=UTF8, encoding=PLAIN_DICTIONARY"`
// // 	Age     int32   `parquet:"name=age, type=INT32, encoding=DELTA_BINARY_PACKED"`
// // 	ID      int64   `parquet:"name=id, type=INT64, encoding=DELTA_BINARY_PACKED "`
// // 	Weight  float32 `parquet:"name=weight, type=FLOAT"`
// // 	Sex     bool    `parquet:"name=sex, type=BOOLEAN"`
// // 	Day     int32   `parquet:"name=day, type=DATE, encoding=DELTA_BINARY_PACKED"`
// // 	Ignored int32   //without parquet tag and won't write
// // }

// // type oandaTick struct {
// // 	Ticker    string  `parquet:"name=ticker, type=UTF8, encoding=PLAIN_DICTIONARY"`
// // 	Timestamp int64   `parquet:"name=timestamp, type=TIMESTAMP_MICROS, encoding=DELTA_BINARY_PACKED"`
// // 	Status    string  `parquet:"name=status, type=UTF8"`
// // 	Bid       float32 `parquet:"name=bid, type=FLOAT"`
// // 	Ask       float32 `parquet:"name=ask, type=FLOAT"`
// // }

// // func main() {
// // 	var err error
// // 	fw, err := local.NewLocalFileWriter("flat.parquet")
// // 	if err != nil {
// // 		log.Println("Can't create local file", err)
// // 		return
// // 	}

// // 	//write
// // 	pw, err := writer.NewParquetWriter(fw, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet writer", err)
// // 		return
// // 	}

// // 	pw.RowGroupSize = 128 * 1024 * 1024 //128M
// // 	pw.CompressionType = parquet.CompressionCodec_GZIP
// // 	num := 100000
// // 	for i := 0; i < num; i++ {
// // 		stu := Student{
// // 			Name:   "StudentName",
// // 			Age:    int32(20 + i%5),
// // 			ID:     int64(i),
// // 			Weight: float32(50.0 + float32(i)*0.1),
// // 			Sex:    bool(i%2 == 0),
// // 			Day:    int32(time.Now().Unix() / 3600 / 24),
// // 		}
// // 		if err = pw.Write(stu); err != nil {
// // 			log.Println("Write error", err)
// // 		}
// // 	}

// // 	if err = pw.WriteStop(); err != nil {
// // 		log.Println("WriteStop error", err)
// // 		return
// // 	}
// // 	log.Println("Write Finished")
// // 	fw.Close()

// // 	///read
// // 	fr, err := local.NewLocalFileReader("flat.parquet")
// // 	if err != nil {
// // 		log.Println("Can't open file")
// // 		return
// // 	}

// // 	pr, err := reader.NewParquetReader(fr, new(Student), 4)
// // 	if err != nil {
// // 		log.Println("Can't create parquet reader", err)
// // 		return
// // 	}
// // 	num = int(pr.GetNumRows())
// // 	for i := 0; i < num/10; i++ {
// // 		if i%2 == 0 {
// // 			pr.SkipRows(10) //skip 10 rows
// // 			continue
// // 		}
// // 		stus := make([]Student, 10) //read 10 rows
// // 		if err = pr.Read(&stus); err != nil {
// // 			log.Println("Read error", err)
// // 		}
// // 		log.Println(stus)
// // 	}

// // 	pr.ReadStop()
// // 	fr.Close()

// // }
