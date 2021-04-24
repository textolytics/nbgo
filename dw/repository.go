package dw

// DataframeRepository interface {
type DataWarehouseDataFrameRepository interface {
	Find(code string) (*DataWarehouseDataFrame, error)
	Store(dataWarehouseDataFrame *DataWarehouseDataFrame) error
}

// // RedirectRepository interface {
// type RedirectRepository interface {
// 	Find(code string) (*Redirect, error)
// 	Store(redirect *Redirect) error
// }
