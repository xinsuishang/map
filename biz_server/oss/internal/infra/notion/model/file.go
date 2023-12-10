package model

var externalType = "external"
var filesType = "files"

type FileData struct {
	BaseData
	Name     string        `json:"name,omitempty"`
	External *ExternalData `json:"external,omitempty"`
}

func NewFiles(urls ...string) *PropertyData {
	files := make([]FileData, 0, len(urls))
	for _, url := range urls {
		files = append(files, *newFile(url))
	}
	data := PropertyData{
		Files: files,
	}
	data.Type = filesType
	return &data
}

func newFile(url string) *FileData {
	ext := &ExternalData{
		Url: url,
	}
	data := FileData{
		Name:     url,
		External: ext,
	}
	data.Type = externalType
	return &data
}
