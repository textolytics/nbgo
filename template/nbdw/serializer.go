package nbdw

//DataWarehouseDataFrame interface {
type DataWarehouseDataFrameSerializer interface {
	Decode(input []byte) (*DataWarehouseDataFrame, error)
	Encode(input *DataWarehouseDataFrame) ([]byte, error)
}

// //RedirectSerializer interface {
// type RedirectSerializer interface {
// 	Decode(input []byte) (*Redirect, error)
// 	Encode(input *Redirect) ([]byte, error)
// }
