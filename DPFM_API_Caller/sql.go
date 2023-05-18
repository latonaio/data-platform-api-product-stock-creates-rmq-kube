package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var productStock *dpfm_api_output_formatter.ProductStock
	for _, fn := range accepter {
		switch fn {
		case "ProductStock":
			productStock = c.productStockCreateSql(nil, mtx, input, output, errs, log)

		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductStock: productStock,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var productStock *dpfm_api_output_formatter.ProductStock
	for _, fn := range accepter {
		switch fn {
		case "ProductStock":
			productStock = c.productStockUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		ProductStock: productStock,
	}

	return data
}

func (c *DPFMAPICaller) productStockCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductStock {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_orders_header_dataの更新
	productStockData := input.ProductStock
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productStockData, "function": "ProductStock", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "ProductStock Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductStockCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) productStockUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.ProductStock {
	productStock := input.ProductStock
	productStockData := dpfm_api_processing_formatter.ConvertToProductStockUpdates(productStock)

	sessionID := input.RuntimeSessionID
	if productStockIsUpdate(productStockData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productStockData, "function": "ProductStock", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ProductStock Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProductStockUpdates(productStock)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func productStockIsUpdate(productStockdata *dpfm_api_processing_formatter.ProductStockUpdates) bool {
	productStock := productStockdata.ProductStock

	return !(productStock == 0)
}
