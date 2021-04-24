package dw

//RedirectService dwindows 10 built dools
type DwService interface {
	Find(code string) (*DataWarehouseDataFrame, error)
	Store(dataWarehouseDataFrame *DataWarehouseDataFrame) error
}
