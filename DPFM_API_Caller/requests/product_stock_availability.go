package requests

type ProductStockAvailability struct {
	Product                      string   `json:"Product"`
	BusinessPartner              int      `json:"BusinessPartner"`
	Plant                        string   `json:"Plant"`
	ProductStockAvailabilityDate string   `json:"ProductStockAvailabilityDate"`
	InventoryStockType           *string  `json:"InventoryStockType"`
	InventorySpecialStockType    *string  `json:"InventorySpecialStockType"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
