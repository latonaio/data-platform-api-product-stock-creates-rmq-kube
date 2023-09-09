package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Input_Reader"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToProductStockCreates(sdc *dpfm_api_input_reader.SDC) (*ProductStock, error) {
	data := sdc.ProductStock

	productStock, err := TypeConverter[*ProductStock](data)
	if err != nil {
		return nil, err
	}

	return productStock, nil
}

func ConvertToProductStockUpdates(productStockdata dpfm_api_input_reader.ProductStock) (*ProductStock, error) {
	data := productStockdata

	productStock, err := TypeConverter[*ProductStock](data)
	if err != nil {
		return nil, err
	}

	return productStock, nil
}

func ConvertToProductStockAvailabilityCreates(sdc *dpfm_api_input_reader.SDC) (*ProductStockAvailability, error) {
	data := sdc.ProductStockAvailability

	productStockAvailability, err := TypeConverter[*ProductStockAvailability](data)
	if err != nil {
		return nil, err
	}

	return productStockAvailability, nil
}

func ConvertToProductStockAvailabilityUpdates(productStockAvailabilitydata dpfm_api_input_reader.ProductStockAvailability) (*ProductStockAvailability, error) {
	data := productStockAvailabilitydata

	productStockAvailability, err := TypeConverter[*ProductStockAvailability](data)
	if err != nil {
		return nil, err
	}

	return productStockAvailability, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
