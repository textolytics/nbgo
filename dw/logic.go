package dw

import (
	"errors"

	errs "github.com/pkg/errors"
	validate "gopkg.in/go-playground/validator.v9"
)

var (
	//ErrRedirectNotFound sdsdsd
	ErrRedirectNotFound = errors.New("DataWarehouseDataFrame Not Found")
	//ErrRedirectInvalid sdsds
	ErrRedirectInvalid = errors.New("DataWarehouseDataFrame Invalid")
)

type dwService struct {
	dataWarehouseDataFrameRepository DataWarehouseDataFrameRepository
}

//NewRedirectService (redirectRepo RedirectRepository) RedirectService
func NewDwService(dataWarehouseDataFrameRepository DataWarehouseDataFrameRepository) DwService {
	return &dwService{
		dataWarehouseDataFrameRepository,
	}
}

func (d *dwService) Find(code string) (*DataWarehouseDataFrame, error) {
	return d.dataWarehouseDataFrameRepository.Find(code)
}

func (d *dwService) Store(DataWarehouseDataFrame *DataWarehouseDataFrame) error {
	if err := validate.New().Struct(DataWarehouseDataFrame); err != nil {
		return errs.Wrap(ErrRedirectInvalid, "service.DataWarehouseDataFrame.Store")
	}
	errs.Wrap(ErrRedirectInvalid, "service.DataWarehouseDataFrame.Store")
	// DataWarehouseDataFrame.Destination = shortid.MustGenerate()
	// DataWarehouseDataFrame.Data = time.Now().UTC().Unix()
	return d.dataWarehouseDataFrameRepository.Store(DataWarehouseDataFrame)
}
