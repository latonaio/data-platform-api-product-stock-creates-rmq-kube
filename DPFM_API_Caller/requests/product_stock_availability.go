package requests

type ProductStockAvailability struct {
	BusinessPartner              *int     `json:"BusinessPartner"`
	Product                      string   `json:"Product"`
	Plant                        string   `json:"Plant"`
	Batch                        string   `json:"Batch"`
	BatchValidityEndDate         *string  `json:"BatchValidityEndDate"`
	OrderID                      *int     `json:"OrderID"`
	OrderItem                    *int     `json:"OrderItem"`
	Project                      string   `json:"Project"`
	InventoryStockType           string   `json:"InventoryStockType"`
	InventorySpecialStockType    string   `json:"InventorySpecialStockType"`
	ProductStockAvailabilityDate *string  `json:"ProductStockAvailabilityDate"`
	AvailableProductStock        *float32 `json:"AvailableProductStock"`
}
