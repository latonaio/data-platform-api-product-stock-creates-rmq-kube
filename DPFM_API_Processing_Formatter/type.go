package dpfm_api_processing_formatter

type ProductStockUpdates struct {
	ProductStock float32 `json:"ProductStock"`
}

type ProductStockByBatchUpdates struct {
	ProductStock *float32 `json:"ProductStock"`
}

type ProductStocByStorageBinUpdates struct {
	InventorySpecialStockType *string  `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}

type ProductStocByStorageBinByBatchBinUpdates struct {
	InventorySpecialStockType *string  `json:"InventorySpecialStockType"`
	ProductStock              *float32 `json:"ProductStock"`
}

type ProductStockAvailabilityUpdates struct {
	AvailabileProductStock *float32 `json:"AvailabileProductStock"`
}

type ProductStockAvailabilityByBatchUpdates struct {
	AvailabileProductStock *float32 `json:"AvailabileProductStock"`
}

type ProductStockAvailabilityByStorageBinUpdates struct {
	AvailabileProductStock *float32 `json:"AvailabileProductStock"`
}

type ProductStockAvailabilityByStorageBinByBatchUpdates struct {
	AvailabileProductStock *float32 `json:"AvailabileProductStock"`
}
