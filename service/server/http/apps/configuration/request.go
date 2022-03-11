package configuration

type newConfigurationRequest struct {
	Uuid        string                 `json:"uuid"`
	Description string                 `json:"description"`
	Data        map[string]interface{} `json:"data"`
}

type updateConfigurationRequest struct {
	Description string                 `json:"description"`
	Data        map[string]interface{} `json:"data"`
}
