package existence_conf

type Returns struct {
	ConnectionKey       string              `json:"connection_key"`
	Result              bool                `json:"result"`
	RedisKey            string              `json:"redis_key"`
	RuntimeSessionID    string              `json:"runtime_session_id"`
	BusinessPartnerID   *int                `json:"business_partner"`
	Filepath            string              `json:"filepath"`
	ServiceLabel        string              `json:"service_label"`
	ProductMasterReturn ProductMasterReturn `json:"ProductMaster"`
	APISchema           string              `json:"api_schema"`
	Accepter            []string            `json:"accepter"`
	Deleted             bool                `json:"deleted"`
}

type ProductMasterReturn struct {
	BPPlant struct {
		Product         string `json:"Product"`
		BusinessPartner int    `json:"BusinessPartner"`
		Plant           string `json:"Plant"`
		ExistenceConf   bool   `json:"ExistenceConf"`
	} `json:"BPPlant"`
}
