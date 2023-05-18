package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-product-stock-creates-rmq-kube/config"
	"data-platform-api-product-stock-creates-rmq-kube/existence_conf"
	"sync"
	"time"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type DPFMAPICaller struct {
	ctx  context.Context
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient

	configure *existence_conf.ExistenceConf
}

func NewDPFMAPICaller(
	conf *config.Conf, rmq *rabbitmq.RabbitmqClient,

	confirmor *existence_conf.ExistenceConf,
) *DPFMAPICaller {
	return &DPFMAPICaller{
		ctx:       context.Background(),
		conf:      conf,
		rmq:       rmq,
		configure: confirmor,
	}
}

func (c *DPFMAPICaller) AsyncProductStockCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	log *logger.Logger,
) (interface{}, []error) {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	exconfFin := make(chan error)

	// 他PODへ問い合わせ
	wg.Add(1)
	go c.exconfProcess(&mtx, &wg, exconfFin, input, output, &exconfAllExist, accepter, &errs, log)

	// 処理待ち
	ticker := time.NewTicker(10 * time.Second)
	if err := c.finWait(&mtx, exconfFin, ticker); err != nil || len(errs) != 0 {
		if err != nil {
			errs = append(errs, err)
		}
		return dpfm_api_output_formatter.Message{}, errs
	}
	if !exconfAllExist {
		mtx.Lock()
		return dpfm_api_output_formatter.Message{}, nil
	}
	wg.Wait()

	var response interface{}
	// SQL処理
	if input.APIType == "creates" {
		response = c.createSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	} else if input.APIType == "updates" {
		response = c.updateSqlProcess(nil, &mtx, input, output, accepter, &errs, log)
	}

	return response, nil
}

func (c *DPFMAPICaller) exconfProcess(
	mtx *sync.Mutex,
	wg *sync.WaitGroup,
	exconfFin chan error,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	exconfAllExist *bool,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) {
	defer wg.Done()
	var e []error
	*exconfAllExist, e = c.configure.Conf(input, output, accepter, log)
	if len(e) != 0 {
		mtx.Lock()
		*errs = append(*errs, e...)
		mtx.Unlock()
		exconfFin <- xerrors.New("exconf error")
		return
	}
	exconfFin <- nil
}

func (c *DPFMAPICaller) finWait(
	mtx *sync.Mutex,
	finChan chan error,
	ticker *time.Ticker,
) error {
	select {
	case e := <-finChan:
		if e != nil {
			mtx.Lock()
			return e
		}
	case <-ticker.C:
		return xerrors.New("time out")
	}
	return nil
}

func checkResult(msg rabbitmq.RabbitmqMessage) bool {
	data := msg.Data()
	d, ok := data["result"]
	if !ok {
		return false
	}
	result, ok := d.(string)
	if !ok {
		return false
	}
	return result == "success"
}

func getBoolPtr(b bool) *bool {
	return &b
}

func contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// 	for _, fn := range accepter {
// 		wg.Add(1)
// 		switch fn {
// 		case "ProductStock":
// 			go c.ProductStock(&wg, &mtx, sqlUpdateFin, log, &errs, input)
// 		case "ProductStockAvailability":
// 			go c.ProductStockAvailability(&wg, &mtx, sqlUpdateFin, log, &errs, input)
// 		default:
// 			wg.Done()
// 		}
// 	}

// 	// 後処理
// 	ticker := time.NewTicker(10 * time.Second)
// 	select {
// 	case e := <-sqlUpdateFin:
// 		if e != nil {
// 			mtx.Lock()
// 			errs = append(errs, e)
// 			return errs
// 		}
// 	case <-ticker.C:
// 		mtx.Lock()
// 		errs = append(errs, xerrors.New("time out"))
// 		return errs
// 	}

// 	return nil
// }

// func (c *DPFMAPICaller) ProductStock(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
// 	var err error = nil
// 	defer wg.Done()
// 	defer func() {
// 		errFin <- err
// 	}()
// 	sessionID := sdc.RuntimeSessionID
// 	ctx := context.Background()

// 	// data_platform_product_stock_product_stock_dataの更新
// 	productStockData := sdc.ProductStock
// 	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productStockData, "function": "ProductStockProductStock", "runtime_session_id": sessionID})
// 	if err != nil {
// 		err = xerrors.Errorf("rmq error: %w", err)
// 		return
// 	}
// 	res.Success()
// 	if !checkResult(res) {
// 		// err = xerrors.New("Product Stock Data cannot insert")
// 		sdc.SQLUpdateResult = getBoolPtr(false)
// 		sdc.SQLUpdateError = "Product Stock Data cannot insert"
// 		return
// 	}

// 	sdc.SQLUpdateResult = getBoolPtr(true)
// 	return
// }

// func (c *DPFMAPICaller) ProductStockAvailability(wg *sync.WaitGroup, mtx *sync.Mutex, errFin chan error, log *logger.Logger, errs *[]error, sdc *dpfm_api_input_reader.SDC) {
// 	var err error = nil
// 	defer wg.Done()
// 	defer func() {
// 		errFin <- err
// 	}()
// 	sessionID := sdc.RuntimeSessionID
// 	ctx := context.Background()

// 	// data_platform_product_stock_availability_dataの更新
// 	productStockAvailabilityData := sdc.ProductStock.ProductStockAvailability
// 	res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": productStockAvailabilityData, "function": "ProductStockAvailability", "runtime_session_id": sessionID})
// 	if err != nil {
// 		err = xerrors.Errorf("rmq error: %w", err)
// 		return
// 	}
// 	res.Success()
// 	if !checkResult(res) {
// 		// err = xerrors.New("Product Stock Availability Data cannot insert")
// 		sdc.SQLUpdateResult = getBoolPtr(false)
// 		sdc.SQLUpdateError = "Product Stock Availability Data cannot insert"
// 		return
// 	}

// 	sdc.SQLUpdateResult = getBoolPtr(true)
// 	return
// }
