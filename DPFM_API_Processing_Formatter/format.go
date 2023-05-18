package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToProductStockUpdates(productStock dpfm_api_input_reader.ProductStock) *ProductStockUpdates {
	data := productStock

	return &ProductStockUpdates{
		ProductStock: *data.ProductStock,
	}
}
