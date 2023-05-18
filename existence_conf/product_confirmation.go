package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-product-stock-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) productMasterExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	productStocks := make([]dpfm_api_input_reader.ProductStock, 0, 1)
	productStocks = append(productStocks, input.ProductStock)
	for _, productStock := range productStocks {
		product, businessPartner, plant := getProductMasterBpPlantExistenceConfKey(mapper, &productStock, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(product) {
				wg2.Done()
				return
			}
			res, err := c.productMasterBpPlantExistenceConfRequest(product, businessPartner, plant, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) productMasterBpPlantExistenceConfRequest(product string, businessPartner int, plant string, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"Product":         product,
		"BusinessPartner": businessPartner,
		"Plant":           plant,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.ProductMasterReturn.BPPlant.Product = product
	req.ProductMasterReturn.BPPlant.BusinessPartner = businessPartner
	req.ProductMasterReturn.BPPlant.Plant = plant
	req.Accepter = []string{"BPPlant"}

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}

	return "", nil
}

func getProductMasterBpPlantExistenceConfKey(mapper ExConfMapper, productStock *dpfm_api_input_reader.ProductStock, exconfErrMsg *string) (string, int, string) {
	var product string
	var businessPartner int
	var plant string

	switch mapper.Field {
	case "Product":
		product = productStock.Product
		businessPartner = productStock.BusinessPartner
		plant = productStock.Plant
	}
	return product, businessPartner, plant
}

func productMasterConfKeyExistence(res map[string]interface{}, tableTag string) bool {
	req, err := jsonTypeConversion[Returns](res)
	if err != nil {
		return false
	}

	if tableTag == "ProductMasterBPPlant" {
		return req.ProductMasterReturn.BPPlant.ExistenceConf
	}

	return false
}
