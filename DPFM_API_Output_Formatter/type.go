package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	ProductStock                                *ProductStock                                  `json:"ProductStock"`
	ProductStockByBatch                         *[]ProductStockByBatch                         `json:"ProductStockByBatch"`
	ProductStockByOrder                         *[]ProductStockByOrder                         `json:"ProductStockByOrder"`
	ProductStockByProject                       *[]ProductStockByProject                       `json:"ProductStockByProject"`
	ProductStockByStorageBin                    *[]ProductStockByStorageBin                    `json:"ProductStockByStorageBin"`
	ProductStockByStorageBinByBatch             *[]ProductStockByStorageBinByBatch             `json:"ProductStockByStorageBinByBatch"`
	ProductStockAvailabilityByOrder             *[]ProductStockAvailabilityByOrder             `json:"ProductStockAvailabilityByOrder"`
	ProductStockAvailabilityByProject           *[]ProductStockAvailabilityByProject           `json:"ProductStockAvailabilityByProject"`
	ProductStockAvailability                    *[]ProductStockAvailability                    `json:"ProductStockAvailability"`
	ProductStockAvailabilityByBatch             *[]ProductStockAvailabilityByBatch             `json:"ProductStockAvailabilityByBatch"`
	ProductStockAvailabilityByStorageBin        *[]ProductStockAvailabilityByStorageBin        `json:"ProductStockAvailabilityByStorageBin"`
	ProductStockAvailabilityByStorageBinByBatch *[]ProductStockAvailabilityByStorageBinByBatch `json:"ProductStockAvailabilityByStorageBinByBatch"`
}

type ProductStock struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	InventoryStockType                     string   `json:"InventoryStockType"`
	ProductStock                           *float32 `json:"ProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockAvailability struct {
	Product                                string   `json:"Product"`
	BusinessPartner                        int      `json:"BusinessPartner"`
	Plant                                  string   `json:"Plant"`
	SupplyChainRelationshipID              int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int      `json:"Buyer"`
	Seller                                 int      `json:"Seller"`
	DeliverToParty                         int      `json:"DeliverToParty"`
	DeliverFromParty                       int      `json:"DeliverFromParty"`
	DeliverToPlant                         string   `json:"DeliverToPlant"`
	DeliverFromPlant                       string   `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string   `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  *float32 `json:"AvailableProductStock"`
	CreationDate                           *string  `json:"CreationDate"`
	CreationTime                           *string  `json:"CreationTime"`
	LastChangeDate                         *string  `json:"LastChangeDate"`
	LastChangeTime                         *string  `json:"LastChangeTime"`
}

type ProductStockByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	InventoryStockType                     string  `json:"InventoryStockType"`
	ProductStock                           float32 `json:"ProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockByOrder struct {
	Product                                string  `json:Product`
	OrderID                                int     `json:OrderID`
	OrderItem                              int     `json:OrderItem`
	SupplyChainRelationshipID              int     `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int     `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int     `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int     `json:Buyer`
	Seller                                 int     `json:Seller`
	DeliverToParty                         int     `json:DeliverToParty`
	DeliverFromParty                       int     `json:DeliverFromParty`
	DeliverToPlant                         string  `json:DeliverToPlant`
	DeliverFromPlant                       string  `json:DeliverFromPlant`
	InventoryStockType                     string  `json:InventoryStockType`
	ProductStock                           float32 `json:ProductStock`
	CreationDate                           string  `json:CreationDate`
	CreationTime                           string  `json:CreationTime`
	LastChangeDate                         string  `json:LastChangeDate`
	LastChangeTime                         string  `json:LastChangeTime`
}

type ProductStockByProject struct {
	Product                                string  `json:Product`
	Project                                int     `json:Project`
	WBSElement                             int     `json:WBSElement`
	SupplyChainRelationshipID              int     `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int     `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int     `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int     `json:Buyer`
	Seller                                 int     `json:Seller`
	DeliverToParty                         int     `json:DeliverToParty`
	DeliverFromParty                       int     `json:DeliverFromParty`
	DeliverToPlant                         string  `json:DeliverToPlant`
	DeliverFromPlant                       string  `json:DeliverFromPlant`
	InventoryStockType                     string  `json:InventoryStockType`
	ProductStock                           float32 `json:ProductStock`
	CreationDate                           string  `json:CreationDate`
	CreationTime                           string  `json:CreationTime`
	LastChangeDate                         string  `json:LastChangeDate`
	LastChangeTime                         string  `json:LastChangeTime`
}

type ProductStockByStorageBin struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	InventoryStockType                     string  `json:"InventoryStockType"`
	ProductStock                           float32 `json:"ProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockByStorageBinByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	InventoryStockType                     string  `json:"InventoryStockType"`
	ProductStock                           float32 `json:"ProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  float32 `json:"AvailableProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityByOrder struct {
	Product                                string  `json:Product`
	OrderID                                int     `json:OrderID`
	OrderItem                              int     `json:OrderItem`
	SupplyChainRelationshipID              int     `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int     `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int     `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int     `json:Buyer`
	Seller                                 int     `json:Seller`
	DeliverToParty                         int     `json:DeliverToParty`
	DeliverFromParty                       int     `json:DeliverFromParty`
	DeliverToPlant                         string  `json:DeliverToPlant`
	DeliverFromPlant                       string  `json:DeliverFromPlant`
	ProductStockAvailabilityDate           string  `json:ProductStockAvailabilityDate`
	AvailableProductStock                  float32 `json:AvailableProductStock`
	CreationDate                           string  `json:CreationDate`
	CreationTime                           string  `json:CreationTime`
	LastChangeDate                         string  `json:LastChangeDate`
	LastChangeTime                         string  `json:LastChangeTime`
}

type ProductStockAvailabilityByProject struct {
	Product                                string  `json:Product`
	Project                                int     `json:Project`
	WBSElement                             int     `json:WBSElement`
	SupplyChainRelationshipID              int     `json:SupplyChainRelationshipID`
	SupplyChainRelationshipDeliveryID      int     `json:SupplyChainRelationshipDeliveryID`
	SupplyChainRelationshipDeliveryPlantID int     `json:SupplyChainRelationshipDeliveryPlantID`
	Buyer                                  int     `json:Buyer`
	Seller                                 int     `json:Seller`
	DeliverToParty                         int     `json:DeliverToParty`
	DeliverFromParty                       int     `json:DeliverFromParty`
	DeliverToPlant                         string  `json:DeliverToPlant`
	DeliverFromPlant                       string  `json:DeliverFromPlant`
	ProductStockAvailabilityDate           string  `json:ProductStockAvailabilityDate`
	AvailableProductStock                  float32 `json:AvailableProductStock`
	CreationDate                           string  `json:CreationDate`
	CreationTime                           string  `json:CreationTime`
	LastChangeDate                         string  `json:LastChangeDate`
	LastChangeTime                         string  `json:LastChangeTime`
}

type ProductStockAvailabilityByStorageBin struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  float32 `json:"AvailableProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}

type ProductStockAvailabilityByStorageBinByBatch struct {
	Product                                string  `json:"Product"`
	BusinessPartner                        int     `json:"BusinessPartner"`
	Plant                                  string  `json:"Plant"`
	StorageLocation                        string  `json:"StorageLocation"`
	StorageBin                             string  `json:"StorageBin"`
	Batch                                  string  `json:"Batch"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	ProductStockAvailabilityDate           string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock                  float32 `json:"AvailableProductStock"`
	CreationDate                           string  `json:"CreationDate"`
	CreationTime                           string  `json:"CreationTime"`
	LastChangeDate                         string  `json:"LastChangeDate"`
	LastChangeTime                         string  `json:"LastChangeTime"`
}
