package model

var dateType = "date"

type DateData struct {
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
}

func NewDate(dataTime string) *PropertyData {
	date := &DateData{
		Start: dataTime,
	}
	data := PropertyData{
		Date: date,
	}
	data.Type = dateType
	return &data
}
