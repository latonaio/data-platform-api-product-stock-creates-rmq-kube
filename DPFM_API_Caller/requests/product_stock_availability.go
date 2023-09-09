package requests

type ProductStockAvailability struct {
	Product									string	`json:"Product"`
	BusinessPartner							int		`json:"BusinessPartner"`
	Plant									string	`json:"Plant"`
	SupplyChainRelationshipID				int		`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID		int		`json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID	int		`json:"SupplyChainRelationshipDeliveryPlantID"`
	Buyer									int		`json:"Buyer"`
	Seller									int		`json:"Seller"`
	DeliverToParty							int		`json:"DeliverToParty"`
	DeliverFromParty						int		`json:"DeliverFromParty"`
	DeliverToPlant							string	`json:"DeliverToPlant"`
	DeliverFromPlant						string	`json:"DeliverFromPlant"`
	ProductStockAvailabilityDate			string	`json:"ProductStockAvailabilityDate"`
	AvailableProductStock					float32	`json:"AvailableProductStock"`
	CreationDate							string	`json:"CreationDate"`
	CreationTime							string	`json:"CreationTime"`
	LastChangeDate							string	`json:"LastChangeDate"`
	LastChangeTime							string	`json:"LastChangeTime"`
}
