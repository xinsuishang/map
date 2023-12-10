package model

type BaseData struct {
	Type string `json:"type"`
}

type CommonData struct {
	Parent     *PropertyData           `json:"parent,omitempty"`
	Properties map[string]PropertyData `json:"properties"`
}
