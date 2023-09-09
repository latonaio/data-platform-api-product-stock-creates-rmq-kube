package dpfm_api_processing_formatter

type ProductStockUpdates struct {
	ProductStock 			float32	`json:"ProductStock"`
}

type ProductStockAvailabilityUpdates struct {
	AvailableProductStock	float32 `json:"AvailableProductStock"`
}
