package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey                   string                          `json:"connection_key"`
	Result                          bool                            `json:"result"`
	RedisKey                        string                          `json:"redis_key"`
	Filepath                        string                          `json:"filepath"`
	APIStatusCode                   int                             `json:"api_status_code"`
	RuntimeSessionID                string                          `json:"runtime_session_id"`
	BusinessPartner                 *int                            `json:"business_partner"`
	ServiceLabel                    string                          `json:"service_label"`
	APIType                         string                          `json:"api_type"`
	ProductStock                    ProductStock                    `json:"ProductStock"`
	ProductStockByBatch             ProductStockByBatch             `json:"ProductStockByBatch"`
	ProductStockByStorageBin        ProductStockByStorageBin        `json:"ProductStockByStorageBin"`
	ProductStockByStorageBinByBatch ProductStockByStorageBinByBatch `json:"ProductStockByStorageBinByBatch"`
	APISchema                       string                          `json:"api_schema"`
	Accepter                        []string                        `json:"accepter"`
	Deleted                         bool                            `json:"deleted"`
}

type ProductStock struct {
	Product                   string                     `json:"Product"`
	BusinessPartner           int                        `json:"BusinessPartner"`
	Plant                     string                     `json:"Plant"`
	InventoryStockType        *string                    `json:"InventoryStockType"`
	InventorySpecialStockType *string                    `json:"InventorySpecialStockType"`
	ProductStock              *float32                   `json:"ProductStock"`
	ProductStockAvailability  []ProductStockAvailability `json:"ProductStockAvailability"`
}

type ProductStockByBatch struct {
	Product                         string                            `json:"Product"`
	BusinessPartner                 int                               `json:"BusinessPartner"`
	Plant                           string                            `json:"Plant"`
	Batch                           string                            `json:"Batch"`
	InventoryStockType              *string                           `json:"InventoryStockType"`
	InventorySpecialStockType       *string                           `json:"InventorySpecialStockType"`
	ProductStock                    *float32                          `json:"ProductStock"`
	ProductStockAvailabilityByBatch []ProductStockAvailabilityByBatch `json:"ProductStockAvailabilityByBatch"`
}

type ProductStockByStorageBin struct {
	Product                              string                                 `json:"Product"`
	BusinessPartner                      int                                    `json:"BusinessPartner"`
	Plant                                string                                 `json:"Plant"`
	StorageLocation                      string                                 `json:"StorageLocation"`
	StorageBin                           string                                 `json:"StorageBin"`
	InventoryStockType                   *string                                `json:"InventoryStockType"`
	InventorySpecialStockType            *string                                `json:"InventorySpecialStockType"`
	ProductStock                         *float32                               `json:"ProductStock"`
	ProductStockAvailabilityByStorageBin []ProductStockAvailabilityByStorageBin `json:"ProductStockAvailabilityByStorageBin"`
}

type ProductStockByStorageBinByBatch struct {
	Product                                     string                                        `json:"Product"`
	BusinessPartner                             int                                           `json:"BusinessPartner"`
	Plant                                       string                                        `json:"Plant"`
	StorageLocation                             string                                        `json:"StorageLocation"`
	StorageBin                                  string                                        `json:"StorageBin"`
	Batch                                       string                                        `json:"Batch"`
	InventoryStockType                          *string                                       `json:"InventoryStockType"`
	InventorySpecialStockType                   *string                                       `json:"InventorySpecialStockType"`
	ProductStock                                *float32                                      `json:"ProductStock"`
	ProductStockAvailabilityByStorageBinByBatch []ProductStockAvailabilityByStorageBinByBatch `json:"ProductStockAvailabilityByStorageBinByBatch"`
}

type ProductStockAvailability struct {
	Product                      string   `json:"Product"`
	BusinessPartner              int      `json:"BusinessPartner"`
	Plant                        string   `json:"Plant"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}

type ProductStockAvailabilityByBatch struct {
	Product                      string   `json:"Product"`
	BusinessPartner              int      `json:"BusinessPartner"`
	Plant                        string   `json:"Plant"`
	Batch                        string   `json:"Batch"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}

type ProductStockAvailabilityByStorageBin struct {
	Product                      string   `json:"Product"`
	BusinessPartner              int      `json:"BusinessPartner"`
	Plant                        string   `json:"Plant"`
	StorageLocation              string   `json:"StorageLocation"`
	StorageBin                   string   `json:"StorageBin"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}

type ProductStockAvailabilityByStorageBinByBatch struct {
	Product                      string  `json:"Product"`
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageBin                   string  `json:"StorageBin"`
	Batch                        string  `json:"Batch"`
	ProductStockAvailabilityDate string  `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string `json:"InventoryStockType"`
	InventorySpecialStockType    *string `json:"InventorySpecialStockType"`
	AvailableProductStock        float32 `json:"AvailableProductStock"`
}
