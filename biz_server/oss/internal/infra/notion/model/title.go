package model

var contentType = "text"
var titleType = "title"

type TitleData struct {
	BaseData
	Text *ContentData `json:"text,omitempty"`
}

func NewTitles(titles ...string) *PropertyData {
	titleData := make([]TitleData, 0, len(titles))
	for _, title := range titles {
		titleData = append(titleData, *newTitle(title))
	}
	data := PropertyData{
		Title: titleData,
	}
	data.Type = titleType
	return &data
}

func newTitle(title string) *TitleData {
	content := &ContentData{
		Content: title,
	}
	data := TitleData{
		//PlanText: title,
		Text: content,
	}

	data.Type = contentType
	return &data
}
