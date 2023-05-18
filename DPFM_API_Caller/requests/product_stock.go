package requests

type ProductStock struct {
	Product                   string   `json:"Product"`
	BusinessPartner           int      `json:"BusinessPartner"`
	Plant                     string   `json:"Plant"`
	InventoryStockType        *string  `json:"InventoryStockType"`
	InventorySpecialStockType *string  `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}
