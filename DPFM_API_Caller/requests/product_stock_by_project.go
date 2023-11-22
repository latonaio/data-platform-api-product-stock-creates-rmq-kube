package requests

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
