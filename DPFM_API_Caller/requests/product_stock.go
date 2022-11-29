package requests

type ProductStock struct {
	BusinessPartner           *int     `json:"BusinessPartner"`
	Product                   string   `json:"Product"`
	Plant                     string   `json:"Plant"`
	StorageLocation           string   `json:"StorageLocation"`
	Batch                     string   `json:"Batch"`
	OrderID                   *int     `json:"OrderID"`
	OrderItem                 *int     `json:"OrderItem"`
	Project                   string   `json:"Project"`
	InventoryStockType        string   `json:"InventoryStockType"`
	InventorySpecialStockType string   `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}
