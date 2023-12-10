package model

var urlType = "url"
var numberType = "number"

type PropertyData struct {
	BaseData
	ExternalData
	DatabaseId string      `json:"database_id,omitempty"`
	Date       *DateData   `json:"date,omitempty"`
	Number     float64     `json:"number,omitempty"`
	Title      []TitleData `json:"title,omitempty"`
	Files      []FileData  `json:"files,omitempty"`
}

func NewUrl(fileUrl string) *PropertyData {
	data := PropertyData{}
	data.Url = fileUrl
	data.Type = urlType
	return &data
}

func NewNumber(weight float64) *PropertyData {
	data := PropertyData{
		Number: weight,
	}
	data.Type = numberType
	return &data
}
